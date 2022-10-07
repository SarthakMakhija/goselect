//go:build integration && !windows
// +build integration,!windows

package test

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"os"
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
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClauseWithDoubleQuotedLiteral2(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "quoted")
	_, _ = os.CreateTemp(directoryName, "\"File_(45)\".log")

	defer os.RemoveAll(directoryName)

	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select basename from . where eq(basename, \\\"File_(45)\\\") order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("\"File_(45)\"")},
	}
	executor.AssertMatch(t, expected, queryResults)
}
