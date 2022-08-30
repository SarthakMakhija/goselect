package writer

import (
	"bytes"
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"testing"
)

func TestJsonWriter(t *testing.T) {
	newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
	aParser, err := parser.NewParser("select lower(name), contains(lower(name), 'log') from ../resources/TestResultsWithProjections/multi order by 1", newContext)
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	selectQuery, err := aParser.Parse()
	if err != nil {
		t.Fatalf("error is %v", err)
	}
	queryResults, _ := executor.NewSelectQueryExecutor(selectQuery, newContext).Execute()

	backingWriter := new(bytes.Buffer)
	_ = NewJsonWriter(backingWriter).Write(selectQuery.Projections, queryResults)

	op := backingWriter.String()
	expected := "[{\"lower(name)\" : \"testresultswithprojections_a.log\", \"contains(lower(name),log)\" : \"Y\"}, {\"lower(name)\" : \"testresultswithprojections_b.log\", \"contains(lower(name),log)\" : \"Y\"}, {\"lower(name)\" : \"testresultswithprojections_c.txt\", \"contains(lower(name),log)\" : \"N\"}, {\"lower(name)\" : \"testresultswithprojections_d.txt\", \"contains(lower(name),log)\" : \"N\"}]"

	if expected != op {
		t.Fatalf("Expected json writer to write %v, received %v", expected, op)
	}
}
