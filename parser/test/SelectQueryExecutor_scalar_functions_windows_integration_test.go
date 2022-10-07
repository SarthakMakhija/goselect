//go:build integration && windows
// +build integration,windows

package test

import (
	"fmt"
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"os"
	"testing"
)

func TestResultsWithProjectionsUsingIfBlank1(t *testing.T) {
	directory, err := os.MkdirTemp(".", "blank")
	_, _ = os.Create(directory + string(os.PathSeparator) + "hello")
	defer os.RemoveAll(directory)

	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser(fmt.Sprintf("select ifBlank(lower(name), NA), ifBlank(ext, NA) from %v", directory), newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("hello"), context.StringValue("NA")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithFormatSize(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), fmtsize(size) from ./resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.log"), context.StringValue("72 B")},
		{context.StringValue("testresultswithprojections_b.log"), context.StringValue("58 B")},
		{context.StringValue("testresultswithprojections_c.txt"), context.StringValue("58 B")},
		{context.StringValue("testresultswithprojections_d.txt"), context.StringValue("58 B")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsWithConcatWsFunction(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), concatWs(lower(name), uid, gid, '#') from ./resources/TestResultsWithProjections/single", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	uid, gid := "", ""

	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_a.txt"), context.StringValue(fmt.Sprintf("testresultswithprojections_a.txt#%v#%v", uid, gid))},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithProjectionsInNestedDirectories(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select isdir, lower(name), path from .\\resources\\TestResultsWithProjections\\ order by 1 desc, 2 asc", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.BooleanValue(true), context.StringValue("empty"), context.StringValue(".\\resources\\TestResultsWithProjections\\empty")},
		{context.BooleanValue(true), context.StringValue("hidden"), context.StringValue(".\\resources\\TestResultsWithProjections\\hidden")},
		{context.BooleanValue(true), context.StringValue("multi"), context.StringValue(".\\resources\\TestResultsWithProjections\\multi")},
		{context.BooleanValue(true), context.StringValue("single"), context.StringValue(".\\resources\\TestResultsWithProjections\\single")},
		{context.BooleanValue(false), context.StringValue(".make"), context.StringValue(".\\resources\\TestResultsWithProjections\\hidden\\.Make")},
		{context.BooleanValue(false), context.StringValue("empty.log"), context.StringValue(".\\resources\\TestResultsWithProjections\\empty\\Empty.log")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.log"), context.StringValue(".\\resources\\TestResultsWithProjections\\multi\\TestResultsWithProjections_A.log")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_a.txt"), context.StringValue(".\\resources\\TestResultsWithProjections\\single\\TestResultsWithProjections_A.txt")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_b.log"), context.StringValue(".\\resources\\TestResultsWithProjections\\multi\\TestResultsWithProjections_B.log")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_c.txt"), context.StringValue(".\\resources\\TestResultsWithProjections\\multi\\TestResultsWithProjections_C.txt")},
		{context.BooleanValue(false), context.StringValue("testresultswithprojections_d.txt"), context.StringValue(".\\resources\\TestResultsWithProjections\\multi\\TestResultsWithProjections_D.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}
