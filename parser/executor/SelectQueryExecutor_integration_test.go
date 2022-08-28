package executor

import (
	"goselect/parser"
	"goselect/parser/context"
	"testing"
)

func TestResultsWithProjections1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name, now() from ../resources/TestResultsWithProjections", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]string{
		{"TestResultsWithProjections_A.txt", ""},
	}
	assertMatch(t, expected, queryResults, 1)
}

func TestResultsWithProjections2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), base64(name) from ../resources/TestResultsWithProjections", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]string{
		{"testresultswithprojections_a.txt", "VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ="},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsInCaseInsensitiveManner(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("SELECT LOWER(NAME), BASE64(NAME) FROM ../resources/TestResultsWithProjections", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]string{
		{"testresultswithprojections_a.txt", "VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ="},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjections3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/TestResultsWithProjections", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]string{
		{"testresultswithprojections_a.txt", ".txt"},
	}
	assertMatch(t, expected, queryResults)
}

func assertMatch(t *testing.T, expected [][]string, queryResults [][]interface{}, skipAttributeIndices ...int) {
	contains := func(slice []int, value int) bool {
		for _, v := range slice {
			if value == v {
				return true
			}
		}
		return false
	}
	if len(expected) != len(queryResults) {
		t.Fatalf("Expected length of the query results to be %v, received %v", len(expected), len(queryResults))
	}
	for rowIndex, row := range expected {
		if len(row) != len(queryResults[rowIndex]) {
			t.Fatalf("Expected length of the attributes in row index %v to be %v, received %v", rowIndex, len(row), len(queryResults[rowIndex]))
		}
		for attributeIndex, col := range row {
			if !contains(skipAttributeIndices, attributeIndex) && queryResults[rowIndex][attributeIndex] != col {
				t.Fatalf("Expected %v to match %v at row index %v, attribute index %v",
					col,
					queryResults[rowIndex][attributeIndex],
					rowIndex,
					attributeIndex,
				)
			}
		}
	}
}
