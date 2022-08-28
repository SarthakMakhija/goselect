package executor

import (
	"goselect/parser"
	"testing"
)

func TestResultsWithProjections1(t *testing.T) {
	aParser, err := parser.NewParser("select name, now() from ../resources/TestResultsWithProjections")
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery).Execute()
	expected := [][]string{
		{"TestResultsWithProjections_A.txt", ""},
	}
	assertMatch(t, expected, queryResults, 1)
}

func TestResultsWithProjections2(t *testing.T) {
	aParser, err := parser.NewParser("select lower(name), base64(name) from ../resources/TestResultsWithProjections")
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery).Execute()
	expected := [][]string{
		{"testresultswithprojections_a.txt", "VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ="},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjections3(t *testing.T) {
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/TestResultsWithProjections")
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery).Execute()
	expected := [][]string{
		{"testresultswithprojections_a.txt", ".txt"},
	}
	assertMatch(t, expected, queryResults)
}

func assertMatch(t *testing.T, expected [][]string, queryResults [][]interface{}, skipColumnIndices ...int) {
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
			t.Fatalf("Expected length of the columns in row index %v to be %v, received %v", rowIndex, len(row), len(queryResults[rowIndex]))
		}
		for colIndex, col := range row {
			if !contains(skipColumnIndices, colIndex) && queryResults[rowIndex][colIndex] != col {
				t.Fatalf("Expected %v to match %v at row index %v, col index %v",
					col,
					queryResults[rowIndex][colIndex],
					rowIndex,
					colIndex,
				)
			}
		}
	}
}
