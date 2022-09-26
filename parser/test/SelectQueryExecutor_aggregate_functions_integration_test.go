//go:build integration
// +build integration

package test

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"testing"
)

func TestResultsWithProjectionsIncludingCountFunction1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), count(lower(name)), count() from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.Uint32Value(4), context.Uint32Value(4)},
		{context.StringValue("testresultswithprojections_b.log"), context.Uint32Value(4), context.Uint32Value(4)},
		{context.StringValue("testresultswithprojections_c.txt"), context.Uint32Value(4), context.Uint32Value(4)},
		{context.StringValue("testresultswithprojections_d.txt"), context.Uint32Value(4), context.Uint32Value(4)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAverageFunctionWithLimitReturningTheAverageForAllTheValues(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select avg(len(name)) from ../resources/test/TestResultsWithProjections/multi limit 3", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(32)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAverage(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select fmtsize(avg(size)) from ../resources/test/TestResultsWithProjections/multi", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("61 B")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAggregateFunctionInsideAScalar(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(count()) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("4")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedCountFunction1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count(count()) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedCountFunction2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count(count(lower(name))) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedCountFunction3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(count(count())) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("1")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedAverageFunction1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select avg(avg(avg(len(name)))) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(32)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedAverageFunction2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count(avg(avg(len(name)))) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedAverageFunction3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select add(count(), 7) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(11)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedAverageFunctionWithoutAnyOrderBy(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select add(count(), 7) from ../resources/test/TestResultsWithProjections/multi", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(11)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingCountDistinct1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select countd(lower(ext)) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(2)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingCountDistinct2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count(countd(lower(ext))) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingSum1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select sum(len(ext)) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(16)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingSum2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select sum(countd(lower(ext))) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(2)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMin1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select min(name) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("TestResultsWithProjections_A.log")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMin2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select min(min(len(name))) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.IntValue(32)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMax1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select max(name) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("TestResultsWithProjections_D.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMax2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select max(max(len(name))) from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.IntValue(32)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeTypeAndCountDistinct(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select countd(mime) from ../resources/test/TestResultsWithProjections/", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(3)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithMimeTypeAndMin(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select min(mime) from ../resources/test/TestResultsWithProjections/", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("NA")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAggregateFunctionAndWhereClause1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count() from ../resources/test/TestResultsWithProjections/ where endsWith(name, log) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(3)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAggregateFunctionAndWhereClause2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count() from ../resources/test/TestResultsWithProjections/ where isText(mime) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(7)},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingAggregateFunctionAndWhereClause3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select min(name) from ../resources/test/TestResultsWithProjections/ where endsWith(name, log) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("Empty.log")},
	}
	executor.AssertMatch(t, expected, queryResults)
}
