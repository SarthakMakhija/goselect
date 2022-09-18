package writer

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"testing"
)

func TestHtmlFormatter(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), contains(lower(name), 'log') from ../resources/test/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext, executor.NewDefaultOptions()).Execute()

	html := NewHtmlFormatter().Format(selectQuery.Projections, queryResults)
	expected := "<html><body><table style=\"width:100%; border: 1px solid black\"><tr><th style=\"border: 1px solid black\">lower(name)</th><th style=\"border: 1px solid black\">contains(lower(name),log)</th></tr><tr><td style=\"border: 1px solid black\">testresultswithprojections_a.log</td><td style=\"border: 1px solid black\">Y</td></tr><tr><td style=\"border: 1px solid black\">testresultswithprojections_b.log</td><td style=\"border: 1px solid black\">Y</td></tr><tr><td style=\"border: 1px solid black\">testresultswithprojections_c.txt</td><td style=\"border: 1px solid black\">N</td></tr><tr><td style=\"border: 1px solid black\">testresultswithprojections_d.txt</td><td style=\"border: 1px solid black\">N</td></tr><tr><td colspan=\"2\" style=\"border: 1px solid black\">Total Rows: 4</td></tr></table></body></html>"

	if expected != html {
		t.Fatalf("Expected html formatter to format %v, received %v", expected, html)
	}
}
