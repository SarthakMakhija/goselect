package executor

import (
	"goselect/parser"
	"goselect/parser/context"
	"testing"
)

func TestResultsWithAWhereClause1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/TestResultsWithProjections/multi where eq(ext, .log) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/TestResultsWithProjections/multi where eq(add(2,3), 5) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/TestResultsWithProjections/multi where contains(lower(name), a.log) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause4(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/TestResultsWithProjections/multi where eq(lower(substr(name, 0, 3)), test) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause5(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/TestResultsWithProjections/multi where lt(lower(substr(ext, 1)), txt) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
	}
	assertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause6(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/TestResultsWithProjections/multi where lt(add(2,1), 4) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := NewSelectQueryExecutor(selectQuery, newContext).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	assertMatch(t, expected, queryResults)
}
