//go:build integration
// +build integration

package test

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"os"
	"path/filepath"
	"testing"
)

func TestResultsWithAWhereClause1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where eq(ext, .log) order by 1", newContext)
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

func TestResultsWithAWhereClause2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where eq(add(2,3), 5) order by 1", newContext)
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
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where ne(add(2,3), 6) order by 1", newContext)
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
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause4(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where ne(add(2,3), 5) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause5(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where contains(lower(name), a.log) order by 1", newContext)
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
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause6(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where eq(lower(substr(name, 0, 3)), test) order by 1", newContext)
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
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause7(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where lt(lower(substr(ext, 1)), txt) order by 1", newContext)
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

func TestResultsWithAWhereClause8(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where lt(add(2,1), 4) order by 1", newContext)
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
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause9(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where gt(lower(substr(ext, 1)), log) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause10(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where le(lower(substr(ext, 1)), txt) order by 1", newContext)
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
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause11(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where ge(lower(substr(ext, 1)), log) order by 1", newContext)
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
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause12(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where or(eq(lower(substr(ext, 0, 3)), .log), eq(add(2,3), 6)) order by 1", newContext)
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

func TestResultsWithAWhereClause13(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where and(eq(lower(substr(ext, 0, 3)), .log), eq(lower(substr(name, 0, 3)), test)) order by 1", newContext)
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

func TestResultsWithAWhereClause14(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where not(eq(lower(ext), .log)) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause15(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where like(name, .*.log) order by 1", newContext)
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

func TestResultsWithAWhereClause16(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where or(like(name, .*.log), eq(add(2,3), 5)) order by 1", newContext)
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
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause17(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "ctime-test-dir")
	file, _ := os.CreateTemp(directoryName, "ctime-test-file")

	defer os.RemoveAll(directoryName)

	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from "+directoryName+" where gte(ctime, parsedttime(2022-09-09, dt)) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue(filepath.Base(file.Name()))},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause18(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/images/ where eq(isimg(mimetype), true) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("where.png")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause19(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/images/ where eq(isimg(mimetype), Y) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("where.png")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause20(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/images/ where eq(istext(mimetype), true) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	var expected [][]context.Value
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause21(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where eq(istext(mimetype), true) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}

	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause22(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where istext(mimetype) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}

	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause23(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/TestResultsWithProjections/multi where lt(-0.12, +0.15) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{{context.StringValue("testresultswithprojections_a.log")},
		{context.StringValue("testresultswithprojections_b.log")},
		{context.StringValue("testresultswithprojections_c.txt")},
		{context.StringValue("testresultswithprojections_d.txt")},
	}

	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause24(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources/TestResultsWithProjections/ where startsWith(name, Test) order by 1", newContext)
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
		{context.StringValue("TestResultsWithProjections_A.txt")},
		{context.StringValue("TestResultsWithProjections_B.log")},
		{context.StringValue("TestResultsWithProjections_C.txt")},
		{context.StringValue("TestResultsWithProjections_D.txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause25(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources/TestResultsWithProjections/ where endsWith(name, log) order by 1", newContext)
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
		{context.StringValue("TestResultsWithProjections_A.log")},
		{context.StringValue("TestResultsWithProjections_B.log")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause26(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources/ where gt(size, parsesize(90 Kib))", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("where.png")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause27(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources/ where gt(size, parsesize(100 Mb))", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClauseWithSingleQuotedLiteral1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources/TestResultsWithProjections/ where endsWith(name, 'log') order by 1", newContext)
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
		{context.StringValue("TestResultsWithProjections_A.log")},
		{context.StringValue("TestResultsWithProjections_B.log")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClauseWithSingleQuotedLiteral2(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources/ where eq(name, 'File_(1).log') order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("File_(1).log")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClauseWithSingleQuotedLiteral3(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources/ where eq(name, 'File (1).txt') order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("File (1).txt")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClauseWithSingleQuotedLiteral4(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "quoted")
	_, _ = os.CreateTemp(directoryName, "'File (60)'.txt")

	defer os.RemoveAll(directoryName)

	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select basename from . where eq(basename, \"'File (60)'\".txt) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("'File (60)'")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClauseWithDoubleQuotedLiteral1(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources/ where eq(name, \"File_(1).log\") order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("File_(1).log")},
	}
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause29(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ./resources/images where isArchive(mimetype) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	var expected [][]context.Value
	executor.AssertMatch(t, expected, queryResults)
}

func TestResultsWithAWhereClause30(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select name from ./resources where isArchive(mimetype) order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()
	expected := [][]context.Value{
		{context.StringValue("README.md.zip")},
	}
	executor.AssertMatch(t, expected, queryResults)
}
