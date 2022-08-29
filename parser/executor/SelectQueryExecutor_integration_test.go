package executor

import (
	"goselect/parser"
	"goselect/parser/context"
	"testing"
)

func TestResultsWithProjections1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name, now() from ../resources/TestResultsWithProjections/single", newContext)
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
	aParser, err := parser.NewParser("select lower(name), base64(name) from ../resources/TestResultsWithProjections/single", newContext)
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
	aParser, err := parser.NewParser("SELECT LOWER(NAME), BASE64(NAME) FROM ../resources/TestResultsWithProjections/single", newContext)
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
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/TestResultsWithProjections/single", newContext)
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

func TestResultsWithProjectionsAndLimit1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/TestResultsWithProjections/multi limit 3", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	if len(queryResults) != 3 {
		t.Fatalf("Expected result count to be %v, received %v", 3, len(queryResults))
	}
}

func TestResultsWithProjectionsAndLimit2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/TestResultsWithProjections/multi limit 0", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	if len(queryResults) != 0 {
		t.Fatalf("Expected result count to be %v, received %v", 3, len(queryResults))
	}
}

func TestResultsWithProjectionsWithoutProperParametersToAFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name, lower() from ../resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	_, err = NewSelectQueryExecutor(selectQuery, newContext).Execute()
	if err == nil {
		t.Fatalf("Expected an error on running a query with lower() without any parameter")
	}
}

func assertMatch(t *testing.T, expected [][]string, queryResults [][]context.Value, skipAttributeIndices ...int) {
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
			if !contains(skipAttributeIndices, attributeIndex) && queryResults[rowIndex][attributeIndex].Get() != col {
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