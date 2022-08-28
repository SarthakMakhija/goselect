package executor

import (
	"fmt"
	"goselect/parser"
	"testing"
)

func TestResultsWithProjections1(t *testing.T) {
	aParser, err := parser.NewParser("select lower(name), base64(name) from ../resources/TestResultsWithProjections")
	if err != nil {
		fmt.Println(fmt.Errorf("error is %v", err))
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		fmt.Println(fmt.Errorf("error is %v", err))
	}
	queryResults := ExecuteSelect(selectQuery)
	expected := [][]string{
		{"testresultswithprojections_a.txt", "VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ="},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjections2(t *testing.T) {
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/TestResultsWithProjections")
	if err != nil {
		fmt.Println(fmt.Errorf("error is %v", err))
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		fmt.Println(fmt.Errorf("error is %v", err))
	}
	queryResults := ExecuteSelect(selectQuery)
	expected := [][]string{
		{"testresultswithprojections_a.txt", ".txt"},
	}
	assertMatch(t, expected, queryResults)
}

func assertMatch(t *testing.T, expected [][]string, queryResults [][]interface{}) {
	if len(expected) != len(queryResults) {
		t.Fatalf("Expected length of the query results to be %v, received %v", len(expected), len(queryResults))
	}
	for rowIndex, row := range expected {
		for colIndex, col := range row {
			if queryResults[rowIndex][colIndex] != col {
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
