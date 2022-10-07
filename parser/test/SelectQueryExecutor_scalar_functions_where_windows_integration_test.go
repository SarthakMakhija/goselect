//go:build integration && windows
// +build integration,windows

package test

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"testing"
)

func TestResultsWithAWhereClause28(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources/ where eq(size, parsesize(0 Mb)) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue(".Make")},
		{context.StringValue("Empty.log")},
		{context.StringValue("TestResultsWithProjections")},
		{context.StringValue("archive")},
		{context.StringValue("empty")},
		{context.StringValue("hidden")},
		{context.StringValue("images")},
		{context.StringValue("multi")},
		{context.StringValue("single")},
		{context.StringValue("special")},
	}
	executor.AssertMatch(t, expected, queryResults)
}
