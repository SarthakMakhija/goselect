//go:build integration
// +build integration

package writer

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"strings"
	"testing"
)

func TestTableFormatterWithoutWidthOptions(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/test/TestResultsWithProjections/multi order by 1 limit 2", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()

	str := NewTableFormatter().Format(selectQuery.Projections, queryResults)
	lower := strings.ToLower(str)
	if !strings.Contains(lower, "lower(name)") {
		t.Fatalf("Expected lower(name) to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(lower, "testresultswithprojections_a.log") {
		t.Fatalf("Expected testresultswithprojections_a.log to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(lower, "testresultswithprojections_b.log") {
		t.Fatalf("Expected testresultswithprojections_b.log to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(lower, "rows: 2") {
		t.Fatalf("Expected Total Rows: 2 to be contained in the string but was not, received string is %v", str)
	}
}

func TestTableFormatterWithWidthOptions(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name) from ../resources/test/TestResultsWithProjections/multi order by 1 limit 2", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()

	str := NewTableFormatterWithWidthOptions(NewAttributeWidthOptions(6, 15)).Format(selectQuery.Projections, queryResults)
	lower := strings.ToLower(str)
	if !strings.Contains(lower, "lower(name)") {
		t.Fatalf("Expected lower(name) to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(lower, "testresultswith") {
		t.Fatalf("Expected testresultswith to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(lower, "projections_a.l") {
		t.Fatalf("Expected projections_a.l to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(lower, "og") {
		t.Fatalf("Expected og to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(lower, "projections_b.l") {
		t.Fatalf("Expected projections_b.l to be contained in the string but was not, received string is %v", str)
	}
	if !strings.Contains(lower, "rows: 2") {
		t.Fatalf("Expected Total Rows: 2 to be contained in the string but was not, received string is %v", str)
	}
}
