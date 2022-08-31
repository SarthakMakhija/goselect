package executor

import (
	"fmt"
	"goselect/parser"
	"goselect/parser/context"
	"os"
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

func TestResultsWithProjectionsWithConcatWsFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), concatWs(lower(name), uid, gid, '#') from ../resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	uid, gid := os.Getuid(), os.Getgid()

	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue(fmt.Sprintf("testresultswithprojections_a.txt#%v#%v", uid, gid))},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithSubstringFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), substr(lower(name), 15) from ../resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()

	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("projections_a.txt")},
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
	if queryResults.Count() != 3 {
		t.Fatalf("Expected result Count to be %v, received %v", 3, queryResults.Count())
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
	if queryResults.Count() != 0 {
		t.Fatalf("Expected result Count to be %v, received %v", 3, queryResults.Count())
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

func TestResultsWithProjectionsIncludingConcatWsFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), concatws(lower(name), 'FILE', '@') from ../resources/TestResultsWithProjections/multi order by 2, 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue("testresultswithprojections_a.log@FILE")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue("testresultswithprojections_b.log@FILE")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue("testresultswithprojections_c.txt@FILE")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue("testresultswithprojections_d.txt@FILE")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingContainsFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), contains(lower(name), 'log') from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_b.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_c.txt"), context.BooleanValue(false)},
		{context.StringValue("testresultswithprojections_d.txt"), context.BooleanValue(false)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingCountFunction1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), count(lower(name)), count() from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Uint32Value(4), context.Uint32Value(4)},
		{context.StringValue("testresultswithprojections_b.log"), context.Uint32Value(4), context.Uint32Value(4)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Uint32Value(4), context.Uint32Value(4)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Uint32Value(4), context.Uint32Value(4)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAverageFunctionWithLimitReturningTheAverageForAllTheValues(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select avg(len(name)) from ../resources/TestResultsWithProjections/multi limit 3", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.Float64Value(32)},
		{context.Float64Value(32)},
		{context.Float64Value(32)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAggregateFunctionInsideAScalar(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(count()) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("4")},
		{context.StringValue("4")},
		{context.StringValue("4")},
		{context.StringValue("4")},
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

func assertMatch(t *testing.T, expected [][]context.Value, queryResults *EvaluatingRows, skipAttributeIndices ...int) {
	contains := func(slice []int, value int) bool {
		for _, v := range slice {
			if value == v {
				return true
			}
		}
		return false
	}
	if len(expected) != queryResults.Count() {
		t.Fatalf("Expected length of the query results to be %v, received %v", len(expected), queryResults.Count())
	}
	for rowIndex, row := range expected {
		if len(row) != queryResults.atIndex(rowIndex).TotalAttributes() {
			t.Fatalf("Expected length of the rowAttributes in row index %v to be %v, received %v", rowIndex, len(row), queryResults.atIndex(rowIndex).TotalAttributes())
		}
		rowAttributes := queryResults.atIndex(rowIndex).AllAttributes()
		for attributeIndex, attributeValue := range row {
			if !contains(skipAttributeIndices, attributeIndex) && rowAttributes[attributeIndex].CompareTo(attributeValue) != 0 {
				t.Fatalf("Expected %v to match %v at row index %v, attribute index %v",
					attributeValue,
					rowAttributes[attributeIndex],
					rowIndex,
					attributeIndex,
				)
			}
		}
	}
}
