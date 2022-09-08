package executor

import (
	"goselect/parser"
	"goselect/parser/context"
	"testing"
)

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
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
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
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
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
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("4")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedCountFunction1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count(count()) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedCountFunction2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count(count(lower(name))) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedCountFunction3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(count(count())) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("1")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedAverageFunction1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select avg(avg(avg(len(name)))) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(32)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedAverageFunction2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count(avg(avg(len(name)))) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedAverageFunction3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select add(count(), 7) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(11)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingCountDistinct1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select countd(lower(ext)) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(2)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingCountDistinct2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select count(countd(lower(ext))) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingSum1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select sum(len(ext)) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(16)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingSum2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select sum(countd(lower(ext))) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.Float64Value(2)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMin1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select min(name) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("TestResultsWithProjections_A.log")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMin2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select min(min(len(name))) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.IntValue(32)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMax1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select max(name) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("TestResultsWithProjections_D.txt")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingMax2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select max(max(len(name))) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext, NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.IntValue(32)},
	}
	assertMatch(t, expected, queryResults)
}
