//go:build integration
// +build integration

package test

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"goselect/parser/writer"
	"testing"
)

func TestJsonFormatter(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), contains(lower(name), 'log') from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()

	json := writer.NewJsonFormatter().Format(selectQuery.Projections, queryResults)
	expected := "[{\"lower(name)\" : \"testresultswithprojections_a.log\", \"contains(lower(name),log)\" : \"Y\"}, {\"lower(name)\" : \"testresultswithprojections_b.log\", \"contains(lower(name),log)\" : \"Y\"}, {\"lower(name)\" : \"testresultswithprojections_c.txt\", \"contains(lower(name),log)\" : \"N\"}, {\"lower(name)\" : \"testresultswithprojections_d.txt\", \"contains(lower(name),log)\" : \"N\"}]"

	if expected != json {
		t.Fatalf("Expected json formatter to format %v, received %v", expected, json)
	}
}
