//go:build integration
// +build integration

package executor

import (
	"fmt"
	"goselect/parser"
	"goselect/parser/context"
	"math"
	"os"
	"testing"
)

func TestWithAnErrorWhileRunningAProjection(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower() from .", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	_, err = NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	if err == nil {
		t.Fatalf("Expected an error while executing a query with lower function without any parameter values but did not receive any error")
	}
}

func TestWithAnErrorWhileRunningWhere(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from . where eq(lower(), test)", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	_, err = NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	if err == nil {
		t.Fatalf("Expected an error while executing a query with lower function without any parameter values but did not receive any error")
	}
}

func TestResultsWithProjections1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name, now() from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("TestResultsWithProjections_A.txt"), context.EmptyValue},
	}
	assertMatch(t, expected, queryResults, 1)
}

func TestResultsWithProjections2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), base64(name) from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ=")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjections4(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), replace(name, txt, log), replaceall(name, i, u) from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{
			context.StringValue("testresultswithprojections_a.txt"),
			context.StringValue("TestResultsWithProjections_A.log"),
			context.StringValue("TestResultsWuthProjectuons_A.txt"),
		},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsInNestedDirectories(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name), path from ../resources/test/TestResultsWithProjections/ order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.BooleanValue(true), context.StringValue("empty"), context.StringValue("../resources/test/TestResultsWithProjections/empty")},
		{context.BooleanValue(true), context.StringValue("hidden"), context.StringValue("../resources/test/TestResultsWithProjections/hidden")},
		{context.BooleanValue(true), context.StringValue("multi"), context.StringValue("../resources/test/TestResultsWithProjections/multi")},
		{context.BooleanValue(true), context.StringValue("single"), context.StringValue("../resources/test/TestResultsWithProjections/single")},
		{context.BooleanValue(false), context.StringValue(".make"), context.StringValue("../resources/test/TestResultsWithProjections/hidden/.Make")},
		{context.BooleanValue(false), context.StringValue("empty.log"), context.StringValue("../resources/test/TestResultsWithProjections/empty/Empty.log")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.log"), context.StringValue("../resources/test/TestResultsWithProjections/multi/TestResultsWithProjections_A.log")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.txt"), context.StringValue("../resources/test/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_b.log"), context.StringValue("../resources/test/TestResultsWithProjections/multi/TestResultsWithProjections_B.log")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_c.txt"), context.StringValue("../resources/test/TestResultsWithProjections/multi/TestResultsWithProjections_C.txt")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_d.txt"), context.StringValue("../resources/test/TestResultsWithProjections/multi/TestResultsWithProjections_D.txt")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsInNestedDirectoriesWithOptionToTraverseNestedDirectoriesAsFalse(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name), path from ../resources/test/TestResultsWithProjections/ order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions().DisableNestedTraversal()).Execute()
	expected := [][]context.Value{
		{context.BooleanValue(true), context.StringValue("empty"), context.StringValue("../resources/test/TestResultsWithProjections/empty")},
		{context.BooleanValue(true), context.StringValue("hidden"), context.StringValue("../resources/test/TestResultsWithProjections/hidden")},
		{context.BooleanValue(true), context.StringValue("multi"), context.StringValue("../resources/test/TestResultsWithProjections/multi")},
		{context.BooleanValue(true), context.StringValue("single"), context.StringValue("../resources/test/TestResultsWithProjections/single")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsInNestedDirectoriesWithOptionToIgnoreTraversalOfDirectories(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name), path from ../resources/test/TestResultsWithProjections/ order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}

	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions().EnableNestedTraversal().DirectoriesToIgnoreTraversal([]string{"multi", "empty"})).Execute()
	expected := [][]context.Value{
		{context.BooleanValue(true), context.StringValue("empty"), context.StringValue("../resources/test/TestResultsWithProjections/empty")},
		{context.BooleanValue(true), context.StringValue("hidden"), context.StringValue("../resources/test/TestResultsWithProjections/hidden")},
		{context.BooleanValue(true), context.StringValue("multi"), context.StringValue("../resources/test/TestResultsWithProjections/multi")},
		{context.BooleanValue(true), context.StringValue("single"), context.StringValue("../resources/test/TestResultsWithProjections/single")},
		{context.BooleanValue(false), context.StringValue(".make"), context.StringValue("../resources/test/TestResultsWithProjections/hidden/.Make")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.txt"), context.StringValue("../resources/test/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsInCaseInsensitiveManner(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("SELECT LOWER(NAME), BASE64(NAME) FROM ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ=")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjections3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue(".txt")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithConcatWsFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), concatWs(lower(name), uid, gid, '#') from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	uid, gid := os.Getuid(), os.Getgid()

	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue(fmt.Sprintf("testresultswithprojections_a.txt#%v#%v", uid, gid))},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithSubstringFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), substr(lower(name), 15) from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()

	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("projections_a.txt")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithDayDifference(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select daydiff(now(), now()) from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()

	result := queryResults.atIndex(0).AllAttributes()[0]
	asFloat64, _ := result.GetNumericAsFloat64()

	if math.Round(asFloat64) != float64(0) {
		t.Fatalf("Expected day difference of 2 current times to be equal to zero but received %v and round resulted in %v", asFloat64, math.Round(asFloat64))
	}
}

func TestResultsWithProjectionsWithHourDifference(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select hourdiff(now(), now()) from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()

	result := queryResults.atIndex(0).AllAttributes()[0]
	asFloat64, _ := result.GetNumericAsFloat64()

	if math.Round(asFloat64) != float64(0) {
		t.Fatalf("Expected hour difference of 2 current times to be equal to zero but received %v and round resulted in %v", asFloat64, math.Round(asFloat64))
	}
}

func TestResultsWithProjectionsAndLimit1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/test/TestResultsWithProjections/multi limit 3", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	if queryResults.Count() != 3 {
		t.Fatalf("Expected result Count to be %v, received %v", 3, queryResults.Count())
	}
}

func TestResultsWithProjectionsAndLimit2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/test/TestResultsWithProjections/multi limit 0", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	if queryResults.Count() != 0 {
		t.Fatalf("Expected result Count to be %v, received %v", 3, queryResults.Count())
	}
}

func TestResultsWithProjectionsAndLimit3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/test/TestResultsWithProjections/multi order by 1 limit 2", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsOrderBy1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select upper(lower(name)), ext from ../resources/test/TestResultsWithProjections/multi order by 1 desc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
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
	aParser, err := parser.NewParser("select lower(name), ext from ../resources/test/TestResultsWithProjections/multi order by 2, 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
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
	aParser, err := parser.NewParser("select lower(name), concat(lower(name), '-FILE') from ../resources/test/TestResultsWithProjections/multi order by 2, 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
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
	aParser, err := parser.NewParser("select lower(name), concatws(lower(name), 'FILE', '@') from ../resources/test/TestResultsWithProjections/multi order by 2, 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
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
	aParser, err := parser.NewParser("select lower(name), contains(lower(name), 'log') from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_b.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_c.txt"), context.BooleanValue(false)},
		{context.StringValue("testresultswithprojections_d.txt"), context.BooleanValue(false)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAddFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), add(len(name), 4) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(36)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(36)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(36)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(36)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingSubtractFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), sub(len(name), 2) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(30)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(30)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(30)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(30)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMultiplyFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), mul(len(name), 2) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(64)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(64)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(64)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(64)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingDivideFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), div(len(name), 2) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(16)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(16)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(16)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(16)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNegativeValueInAddSubMulDivFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select add(len(name), -2), sub(len(name), -2), mul(len(name), -2), div(len(name), -2) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(30), context.Float64Value(34), context.Float64Value(-64), context.Float64Value(-16)},
		{context.Float64Value(30), context.Float64Value(34), context.Float64Value(-64), context.Float64Value(-16)},
		{context.Float64Value(30), context.Float64Value(34), context.Float64Value(-64), context.Float64Value(-16)},
		{context.Float64Value(30), context.Float64Value(34), context.Float64Value(-64), context.Float64Value(-16)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithIdentity(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), identity(add(1,2)) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(3.0)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(3.0)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(3.0)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(3.0)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithBaseName(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), lower(basename) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue("testresultswithprojections_a")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue("testresultswithprojections_b")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue("testresultswithprojections_c")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue("testresultswithprojections_d")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithFormatSize(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), fmtsize(size) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue("71 B")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue("58 B")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue("58 B")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue("58 B")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeType1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), mimetype from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("text/plain; charset=utf-8")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeType2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), mimetype from ../resources/test/images where eq(mimetype, image/png)", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("where.png"), context.StringValue("image/png")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeType3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name), mime from ../resources/test/TestResultsWithProjections/ order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.BooleanValue(true), context.StringValue("empty"), context.StringValue("NA")},
		{context.BooleanValue(true), context.StringValue("hidden"), context.StringValue("NA")},
		{context.BooleanValue(true), context.StringValue("multi"), context.StringValue("NA")},
		{context.BooleanValue(true), context.StringValue("single"), context.StringValue("NA")},
		{context.BooleanValue(false), context.StringValue(".make"), context.StringValue("text/plain")},
		{context.BooleanValue(false), context.StringValue("empty.log"), context.StringValue("text/plain")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.log"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.txt"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_b.log"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_c.txt"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_d.txt"), context.StringValue("text/plain; charset=utf-8")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeType4(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name), mime from ../resources/test/TestResultsWithProjections/ where eq(istext(mime), true) order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.BooleanValue(false), context.StringValue(".make"), context.StringValue("text/plain")},
		{context.BooleanValue(false), context.StringValue("empty.log"), context.StringValue("text/plain")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.log"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.txt"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_b.log"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_c.txt"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_d.txt"), context.StringValue("text/plain; charset=utf-8")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsUsingIfBlank1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ifBlank(ext, NA) from ../resources/test/TestResultsWithProjections/hidden", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue(".make"), context.StringValue("NA")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsUsingIfBlank2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select ifBlank(lower(name), Unknown), ifBlank(ext, NA) from ../resources/test/TestResultsWithProjections/hidden", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue(".make"), context.StringValue("NA")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsUsingStartsWith(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), startsWith(name, Test) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_b.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_c.txt"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_d.txt"), context.BooleanValue(true)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsUsingEndsWith(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), endsWith(name, txt) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.BooleanValue(false)},
		{context.StringValue("testresultswithprojections_b.log"), context.BooleanValue(false)},
		{context.StringValue("testresultswithprojections_c.txt"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_d.txt"), context.BooleanValue(true)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithoutProperParametersToAFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name, lower() from ../resources/test/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	_, err = NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	if err == nil {
		t.Fatalf("Expected an error on running a query with lower() without any parameter")
	}
}
