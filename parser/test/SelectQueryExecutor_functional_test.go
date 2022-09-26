//go:build functional
// +build functional

package test

import (
	"encoding/json"
	"fmt"
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"os"
	"testing"
)

func TestSelectQueries(t *testing.T) {
	queries := readAllQueries()
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())

	for _, query := range queries.Queries {
		t.Run(query.Name, func(t *testing.T) {
			aParser, err := parser.NewParser(query.Query, newContext)
			if err != nil && !query.IsErrorExpected {
				t.Fatalf("Did not expect an error for the query named: %v, but received: %v", query.Name, err)
			}
			if err != nil {
				t.Logf("Received an error: %v for the query named: %v", err, query.Name)
				return
			}
			selectQuery, err := aParser.Parse()
			if err != nil && !query.IsErrorExpected {
				t.Fatalf("Did not expect an error for the query named: %v, but received: %v", query.Name, err)
			}
			if err != nil {
				t.Logf("Received an error: %v for the query named: %v", err, query.Name)
				return
			}
			queryResults, err := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
			if err != nil && !query.IsErrorExpected {
				t.Fatalf("Did not expect an error for the query named: %v, but received: %v", query.Name, err)
			}
			if err != nil {
				t.Logf("Received an error: %v for the query named: %v", err, query.Name)
				return
			}
			if queryResults.Count() != query.ResultCount {
				fmt.Println("==============Begin Query Results===============")
				iterator := queryResults.RowIterator()
				for iterator.HasNext() {
					for _, attribute := range iterator.Next().AllAttributes() {
						fmt.Print(attribute.GetAsString(), "\t")
					}
					fmt.Println()
				}
				fmt.Println("==============End Query Results===============")
				t.Fatalf("Expected %v results for the query named: %v, but received: %v", query.ResultCount, query.Name, queryResults.Count())
			}
		})
	}
}

func readAllQueries() Queries {
	var queries Queries
	fileContent, _ := os.ReadFile("./FunctionalQueries.json")
	_ = json.Unmarshal(fileContent, &queries)

	return queries
}

type Queries struct {
	Queries []Query `json:"queries"`
}

type Query struct {
	Name            string `json:"name"`
	Query           string `json:"query"`
	IsErrorExpected bool   `json:"isErrorExpected"`
	ResultCount     uint32 `json:"resultCount"`
}
