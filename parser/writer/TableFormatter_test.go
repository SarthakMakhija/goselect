package writer

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"strings"
	"testing"
)

func TestTableFormatter(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/TestResultsWithProjections/multi order by 1 limit 2", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()

	str := NewTableFormatter().Format(selectQuery.Projections, queryResults)

	if !strings.Contains(str, "lower(name)") {
		t.Fatalf("Expected lower(name) to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(str, "testresultswithprojections_a.log") {
		t.Fatalf("Expected testresultswithprojections_a.log to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(str, "testresultswithprojections_b.log") {
		t.Fatalf("Expected testresultswithprojections_b.log to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(str, "Total Rows: 2") {
		t.Fatalf("Expected Total Rows: 2 to be contained in the string but was not, received string is %v", str)
	}
}
