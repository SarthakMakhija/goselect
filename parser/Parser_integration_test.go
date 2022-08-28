package parser

import (
	"goselect/parser/context"
	"goselect/parser/order"
	"reflect"
	"testing"
)

func TestParsesAnEmptyQueryWithAnError(t *testing.T) {
	_, err := NewParser("", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	if err == nil {
		t.Fatalf("Expected an error while parsing an empty query")
	}
}

func TestParsesANonSelectQueryWithAnError(t *testing.T) {
	parser, _ := NewParser("delete from ~", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	_, err := parser.Parse()

	if err == nil {
		t.Fatalf("Expected an error while parsing a non-select query")
	}
}

func TestParsesAQueryIntoAnASTWithASingleProjection(t *testing.T) {
	parser, _ := NewParser("SELECT name from ~", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 1
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.Projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}
}

func TestParsesAQueryIntoAnASTWithMultipleProjections(t *testing.T) {
	parser, _ := NewParser("SELECT name, lower(name) from ~", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.Projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}
}

func TestParsesAQueryIntoAnASTWithAnOrderBy(t *testing.T) {
	parser, _ := NewParser("SELECT name, upper(lower(name)) from ~ Order by name", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.Projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name", "upper(lower(name))"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}

	ascendingColumns := selectStatement.Order.AscendingColumns
	expectedAscending := []order.ColumnRef{{
		Name:               "name",
		ProjectionPosition: -1,
	}}

	if !reflect.DeepEqual(ascendingColumns, expectedAscending) {
		t.Fatalf("Expected ordering columns to be %v, received %v", expectedAscending, ascendingColumns)
	}
}

func TestParsesAQueryIntoAnASTWithLimit(t *testing.T) {
	parser, _ := NewParser("SELECT name, lower(name) from ~ Order by name Limit 10", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.Projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}

	ascendingColumns := selectStatement.Order.AscendingColumns
	expectedAscending := []order.ColumnRef{{
		Name:               "name",
		ProjectionPosition: -1,
	}}

	if !reflect.DeepEqual(ascendingColumns, expectedAscending) {
		t.Fatalf("Expected ordering columns to be %v, received %v", expectedAscending, ascendingColumns)
	}

	limit := selectStatement.Limit.Limit
	var expectedLimit uint32 = 10

	if expectedLimit != limit {
		t.Fatalf("Expected Limit to be %v, received %v", expectedLimit, limit)
	}
}

func TestParsesAQueryIntoAnASTWithLimitWithoutAnyOrdering(t *testing.T) {
	parser, _ := NewParser("SELECT name, lower(name) from ~/home Limit 10", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.Projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}

	limit := selectStatement.Limit.Limit
	var expectedLimit uint32 = 10

	if expectedLimit != limit {
		t.Fatalf("Expected Limit to be %v, received %v", expectedLimit, limit)
	}
}
