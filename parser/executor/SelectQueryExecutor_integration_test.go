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
	expected := [][]context.Value{
		{context.StringValue("TestResultsWithProjections_A.txt"), context.EmptyValue()},
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
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ=")},
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
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ=")},
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
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue(".txt")},
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

func TestResultsWithProjectionsOrderBy1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select upper(lower(name)), ext from ../resources/TestResultsWithProjections/multi order by 1 desc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("TESTRESULTSWITHPROJECTIONS_D.TXT"), context.StringValue(".txt")},
		{context.StringValue("TESTRESULTSWITHPROJECTIONS_C.TXT"), context.StringValue(".txt")},
		{context.StringValue("TESTRESULTSWITHPROJECTIONS_B.LOG"), context.StringValue(".log")},
		{context.StringValue("TESTRESULTSWITHPROJECTIONS_A.LOG"), context.StringValue(".log")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsOrderBy2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/TestResultsWithProjections/multi order by 2, 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue(".log")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue(".log")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue(".txt")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue(".txt")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingConcatFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), concat(lower(name), '-FILE') from ../resources/TestResultsWithProjections/multi order by 2, 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue("testresultswithprojections_a.log-FILE")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue("testresultswithprojections_b.log-FILE")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue("testresultswithprojections_c.txt-FILE")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue("testresultswithprojections_d.txt-FILE")},
	}
	assertMatch(t, expected, queryResults)
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

func assertMatch(t *testing.T, expected [][]context.Value, queryResults [][]context.Value, skipAttributeIndices ...int) {
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
			if !contains(skipAttributeIndices, attributeIndex) && queryResults[rowIndex][attributeIndex].CompareTo(col) != 0 {
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
