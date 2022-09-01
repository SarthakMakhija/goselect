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
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
		{context.Uint32Value(1)},
		{context.Uint32Value(1)},
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
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
		{context.Uint32Value(1)},
		{context.Uint32Value(1)},
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
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("1")},
		{context.StringValue("1")},
		{context.StringValue("1")},
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
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.Float64Value(32)},
		{context.Float64Value(32)},
		{context.Float64Value(32)},
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
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.Uint32Value(1)},
		{context.Uint32Value(1)},
		{context.Uint32Value(1)},
		{context.Uint32Value(1)},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsIncludingNestedAverageFunction200(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select add(count(), 7) from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.Float64Value(11)},
		{context.Float64Value(11)},
		{context.Float64Value(11)},
		{context.Float64Value(11)},
	}
	assertMatch(t, expected, queryResults)
}