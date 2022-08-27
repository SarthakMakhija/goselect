package parser

import (
	"goselect/parser/order"
	"reflect"
	"testing"
)

func TestParsesAQueryIntoAnASTWithASingleProjection(t *testing.T) {
	parser := NewParser("SELECT name from ~")
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.projections.Count()
	expected := 1
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}
}

func TestParsesAQueryIntoAnASTWithMultipleProjections(t *testing.T) {
	parser := NewParser("SELECT name, lower(name) from ~")
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}
}

func TestParsesAQueryIntoAnASTWithAnOrderBy(t *testing.T) {
	parser := NewParser("SELECT name, upper(lower(name)) from ~ order by name")
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name", "upper(lower(name))"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}

	ascendingColumns := selectStatement.order.AscendingColumns
	expectedAscending := []order.ColumnRef{{
		Name:               "name",
		ProjectionPosition: -1,
	}}

	if !reflect.DeepEqual(ascendingColumns, expectedAscending) {
		t.Fatalf("Expected ordering columns to be %v, received %v", expectedAscending, ascendingColumns)
	}
}

func TestParsesAQueryIntoAnASTWithLimit(t *testing.T) {
	parser := NewParser("SELECT name, lower(name) from ~ order by name limit 10")
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}

	ascendingColumns := selectStatement.order.AscendingColumns
	expectedAscending := []order.ColumnRef{{
		Name:               "name",
		ProjectionPosition: -1,
	}}

	if !reflect.DeepEqual(ascendingColumns, expectedAscending) {
		t.Fatalf("Expected ordering columns to be %v, received %v", expectedAscending, ascendingColumns)
	}

	limit := selectStatement.limit.Limit
	var expectedLimit uint32 = 10

	if expectedLimit != limit {
		t.Fatalf("Expected limit to be %v, received %v", expectedLimit, limit)
	}
}

func TestParsesAQueryIntoAnASTWithLimitWithoutAnyOrdering(t *testing.T) {
	parser := NewParser("SELECT name, lower(name) from ~/home limit 10")
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	expressions := selectStatement.projections.AllExpressions()
	columns := expressions.DisplayableColumns()
	expectedColumns := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Fatalf("Expected columns to be %v, received %v", columns, expectedColumns)
	}

	limit := selectStatement.limit.Limit
	var expectedLimit uint32 = 10

	if expectedLimit != limit {
		t.Fatalf("Expected limit to be %v, received %v", expectedLimit, limit)
	}
}
