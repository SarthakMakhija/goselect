//go:build integration
// +build integration

package test

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"math"
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
	_, err = executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
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
	_, err = executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	if err == nil {
		t.Fatalf("Expected an error while executing a query with lower function without any parameter values but did not receive any error")
	}
}

func TestResultsWithProjections1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name, now() from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("TestResultsWithProjections_A.txt"), context.EmptyValue},
	}
	executor.AssertMatch(t, expected, queryResults, 1)
}

func TestResultsWithProjections2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), base64(name) from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ=")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjections4(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), replace(name, txt, log), replaceall(name, i, u) from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{
			context.StringValue("testresultswithprojections_a.txt"),
			context.StringValue("TestResultsWithProjections_A.log"),
			context.StringValue("TestResultsWuthProjectuons_A.txt"),
		},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsInNestedDirectoriesWithOptionToTraverseNestedDirectoriesAsFalse(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name) from ./resources/TestResultsWithProjections/ order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions().DisableNestedTraversal()).Execute()
	expected := [][]context.Value{
		{context.BooleanValue(true), context.StringValue("empty")},
		{context.BooleanValue(true), context.StringValue("hidden")},
		{context.BooleanValue(true), context.StringValue("multi")},
		{context.BooleanValue(true), context.StringValue("single")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsInNestedDirectoriesWithOptionToIgnoreTraversalOfDirectories(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name) from ./resources/TestResultsWithProjections/ order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}

	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions().EnableNestedTraversal().DirectoriesToIgnoreTraversal([]string{"multi", "empty"})).Execute()
	expected := [][]context.Value{
		{context.BooleanValue(true), context.StringValue("empty")},
		{context.BooleanValue(true), context.StringValue("hidden")},
		{context.BooleanValue(true), context.StringValue("multi")},
		{context.BooleanValue(true), context.StringValue("single")},
		{context.BooleanValue(false), context.StringValue(".make")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsInCaseInsensitiveManner(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("SELECT LOWER(NAME), BASE64(NAME) FROM ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("VGVzdFJlc3VsdHNXaXRoUHJvamVjdGlvbnNfQS50eHQ=")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjections3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue(".txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithSubstringFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), substr(lower(name), 15) from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()

	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("projections_a.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithDayDifference(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select daydiff(now(), now()) from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()

	result := queryResults.AtIndex(0).AllAttributes()[0]
	asFloat64, _ := result.GetNumericAsFloat64()

	if math.Round(asFloat64) != float64(0) {
		t.Fatalf("Expected day difference of 2 current times to be equal to zero but received %v and round resulted in %v", asFloat64, math.Round(asFloat64))
	}
}

func TestResultsWithProjectionsWithHourDifference(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select hourdiff(now(), now()) from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()

	result := queryResults.AtIndex(0).AllAttributes()[0]
	asFloat64, _ := result.GetNumericAsFloat64()

	if math.Round(asFloat64) != float64(0) {
		t.Fatalf("Expected hour difference of 2 current times to be equal to zero but received %v and round resulted in %v", asFloat64, math.Round(asFloat64))
	}
}

func TestResultsWithProjectionsAndLimit1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ./resources/TestResultsWithProjections/multi limit 3", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	if queryResults.Count() != 3 {
		t.Fatalf("Expected result Count to be %v, received %v", 3, queryResults.Count())
	}
}

func TestResultsWithProjectionsAndLimit2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ./resources/TestResultsWithProjections/multi limit 0", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	if queryResults.Count() != 0 {
		t.Fatalf("Expected result Count to be %v, received %v", 3, queryResults.Count())
	}
}

func TestResultsWithProjectionsAndLimit3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi order by 1 limit 2", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsOrderBy1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select upper(lower(name)), ext from ./resources/TestResultsWithProjections/multi order by 1 desc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("TESTRESULTSWITHPROJECTIONS_D.TXT"), context.StringValue(".txt")},
		{context.StringValue("TESTRESULTSWITHPROJECTIONS_C.TXT"), context.StringValue(".txt")},
		{context.StringValue("TESTRESULTSWITHPROJECTIONS_B.LOG"), context.StringValue(".log")},
		{context.StringValue("TESTRESULTSWITHPROJECTIONS_A.LOG"), context.StringValue(".log")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsOrderBy2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), ext from ./resources/TestResultsWithProjections/multi order by 2, 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue(".log")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue(".log")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue(".txt")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue(".txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingConcatFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), concat(lower(name), '-FILE') from ./resources/TestResultsWithProjections/multi order by 2, 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue("testresultswithprojections_a.log-FILE")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue("testresultswithprojections_b.log-FILE")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue("testresultswithprojections_c.txt-FILE")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue("testresultswithprojections_d.txt-FILE")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingConcatWsFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), concatws(lower(name), 'FILE', '@') from ./resources/TestResultsWithProjections/multi order by 2, 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue("testresultswithprojections_a.log@FILE")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue("testresultswithprojections_b.log@FILE")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue("testresultswithprojections_c.txt@FILE")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue("testresultswithprojections_d.txt@FILE")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingContainsFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), contains(lower(name), 'log') from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_b.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_c.txt"), context.BooleanValue(false)},
		{context.StringValue("testresultswithprojections_d.txt"), context.BooleanValue(false)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAddFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), add(len(name), 4) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(36)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(36)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(36)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(36)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingSubtractFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), sub(len(name), 2) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(30)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(30)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(30)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(30)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMultiplyFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), mul(len(name), 2) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(64)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(64)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(64)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(64)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingDivideFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), div(len(name), 2) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(16)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(16)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(16)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(16)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNegativeValueInAddSubMulDivFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select add(len(name), -2), sub(len(name), -2), mul(len(name), -2), div(len(name), -2) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(30), context.Float64Value(34), context.Float64Value(-64), context.Float64Value(-16)},
		{context.Float64Value(30), context.Float64Value(34), context.Float64Value(-64), context.Float64Value(-16)},
		{context.Float64Value(30), context.Float64Value(34), context.Float64Value(-64), context.Float64Value(-16)},
		{context.Float64Value(30), context.Float64Value(34), context.Float64Value(-64), context.Float64Value(-16)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithIdentity(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), identity(add(1,2)) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Float64Value(3.0)},
		{context.StringValue("testresultswithprojections_b.log"), context.Float64Value(3.0)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Float64Value(3.0)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Float64Value(3.0)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithBaseName(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), lower(basename) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue("testresultswithprojections_a")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue("testresultswithprojections_b")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue("testresultswithprojections_c")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue("testresultswithprojections_d")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeType1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), mimetype from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue("text/plain; charset=utf-8")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeType2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), mimetype from ./resources/images where eq(mimetype, image/png)", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("where.png"), context.StringValue("image/png")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeType3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name), mime from ./resources/TestResultsWithProjections/ order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
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
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeType4(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name), mime from ./resources/TestResultsWithProjections/ where eq(istext(mime), true) order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.BooleanValue(false), context.StringValue(".make"), context.StringValue("text/plain")},
		{context.BooleanValue(false), context.StringValue("empty.log"), context.StringValue("text/plain")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.log"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.txt"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_b.log"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_c.txt"), context.StringValue("text/plain; charset=utf-8")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_d.txt"), context.StringValue("text/plain; charset=utf-8")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsUsingStartsWith(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), startsWith(name, Test) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_b.log"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_c.txt"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_d.txt"), context.BooleanValue(true)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsUsingEndsWith(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), endsWith(name, txt) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.BooleanValue(false)},
		{context.StringValue("testresultswithprojections_b.log"), context.BooleanValue(false)},
		{context.StringValue("testresultswithprojections_c.txt"), context.BooleanValue(true)},
		{context.StringValue("testresultswithprojections_d.txt"), context.BooleanValue(true)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithoutProperParametersToAFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name, lower() from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	_, err = executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	if err == nil {
		t.Fatalf("Expected an error on running a query with lower() without any parameter")
	}
}
