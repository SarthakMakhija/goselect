//go:build integration
// +build integration

package test

import (
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/order"
	"reflect"
	"testing"
)

func TestParsesAnEmptyQueryWithAnError(t *testing.T) {
	_, err := parser.NewParser("", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	if err == nil {
		t.Fatalf("Expected an error while parsing an empty query")
	}
}

func TestParsesANonSelectQueryWithAnError(t *testing.T) {
	parser, _ := parser.NewParser("delete from ~", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	_, err := parser.Parse()

	if err == nil {
		t.Fatalf("Expected an error while parsing a non-select query")
	}
}

func TestParsesASelectQueryWithAnErrorInProjection(t *testing.T) {
	parser, _ := parser.NewParser("select from .", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	_, err := parser.Parse()

	if err == nil {
		t.Fatalf("Expected an error while parsing a select with an error in projection")
	}
}

func TestParsesASelectQueryWithAnErrorInSource(t *testing.T) {
	parser, _ := parser.NewParser("select name from ", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	_, err := parser.Parse()

	if err == nil {
		t.Fatalf("Expected an error while parsing a select with an error in source")
	}
}

func TestParsesASelectQueryWithAnErrorBecauseOfTypo(t *testing.T) {
	parser, _ := parser.NewParser("select name from . were eq(1,1)", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	_, err := parser.Parse()

	if err == nil {
		t.Fatalf("Expected an error while parsing a select with an error in where keyword")
	}
}

func TestParsesASelectQueryWithAnErrorInWhere(t *testing.T) {
	parser, _ := parser.NewParser("select name from . where", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	_, err := parser.Parse()

	if err == nil {
		t.Fatalf("Expected an error while parsing a select with an error in where")
	}
}

func TestParsesASelectQueryWithAnErrorInOrderBy(t *testing.T) {
	parser, _ := parser.NewParser("select name from . where eq(1,1) order", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	_, err := parser.Parse()

	if err == nil {
		t.Fatalf("Expected an error while parsing a select with an error in order")
	}
}

func TestParsesASelectQueryWithAnErrorInLimit(t *testing.T) {
	parser, _ := parser.NewParser("select name from . where eq(1,1) order by 1 limit", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	_, err := parser.Parse()

	if err == nil {
		t.Fatalf("Expected an error while parsing a select with an error in limit")
	}
}

func TestParsesAQueryIntoAnASTWithASingleProjection(t *testing.T) {
	parser, _ := parser.NewParser("SELECT name from ~", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 1
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	attributes := selectStatement.Projections.DisplayableAttributes()
	expectedAttributes := []string{"name"}

	if !reflect.DeepEqual(attributes, expectedAttributes) {
		t.Fatalf("Expected attributes to be %v, received %v", attributes, expectedAttributes)
	}
}

func TestParsesAQueryIntoAnASTWithMultipleProjections(t *testing.T) {
	parser, _ := parser.NewParser("SELECT name, lower(name) from ~", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	attributes := selectStatement.Projections.DisplayableAttributes()
	expectedAttributes := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(attributes, expectedAttributes) {
		t.Fatalf("Expected attributes to be %v, received %v", attributes, expectedAttributes)
	}
}

func TestParsesAQueryIntoAnASTWithoutAWhereClause(t *testing.T) {
	parser, _ := parser.NewParser("SELECT name, lower(name) from ~", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	where := selectStatement.Where
	if where.Display() != "" {
		t.Fatalf("Expected where clause to be blank, received %v", where.Display())
	}
}

func TestParsesAQueryIntoAnASTWithWhereClause(t *testing.T) {
	parser, _ := parser.NewParser("SELECT name, lower(name) from ~ where contains(lower(name), log)", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	attributes := selectStatement.Projections.DisplayableAttributes()
	expectedAttributes := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(attributes, expectedAttributes) {
		t.Fatalf("Expected attributes to be %v, received %v", attributes, expectedAttributes)
	}

	where := selectStatement.Where
	if where.Display() != "contains(lower(name),log)" {
		t.Fatalf("Expected where clause to be %v, received %v", "contains(lower(name),log)", where.Display())
	}
}

func TestParsesAQueryIntoAnASTWithAnOrderBy(t *testing.T) {
	parser, _ := parser.NewParser("SELECT name, upper(lower(name)) from ~ Order by 1", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	attributes := selectStatement.Projections.DisplayableAttributes()
	expectedAttributes := []string{"name", "upper(lower(name))"}

	if !reflect.DeepEqual(attributes, expectedAttributes) {
		t.Fatalf("Expected attributes to be %v, received %v", attributes, expectedAttributes)
	}

	attributeRefs := selectStatement.Order.Attributes
	expectedAscending := []order.AttributeRef{{
		ProjectionPosition: 1,
	}}

	if !reflect.DeepEqual(attributeRefs, expectedAscending) {
		t.Fatalf("Expected ordering attributes to be %v, received %v", expectedAscending, attributeRefs)
	}
}

func TestParsesAQueryIntoAnASTWithLimit(t *testing.T) {
	parser, _ := parser.NewParser("SELECT name, lower(name) from ~ Order by 2 Limit 10", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	attributes := selectStatement.Projections.DisplayableAttributes()
	expectedAttributes := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(attributes, expectedAttributes) {
		t.Fatalf("Expected attributes to be %v, received %v", attributes, expectedAttributes)
	}

	attributeRefs := selectStatement.Order.Attributes
	expectedAscending := []order.AttributeRef{{
		ProjectionPosition: 2,
	}}

	if !reflect.DeepEqual(attributeRefs, expectedAscending) {
		t.Fatalf("Expected ordering attributes to be %v, received %v", expectedAscending, attributeRefs)
	}

	limit := selectStatement.Limit.Limit
	var expectedLimit uint32 = 10

	if expectedLimit != limit {
		t.Fatalf("Expected Limit to be %v, received %v", expectedLimit, limit)
	}
}

func TestParsesAQueryIntoAnASTWithLimitWithoutAnyOrdering(t *testing.T) {
	parser, _ := parser.NewParser("SELECT name, lower(name) from ~ Limit 10", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	attributes := selectStatement.Projections.DisplayableAttributes()
	expectedAttributes := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(attributes, expectedAttributes) {
		t.Fatalf("Expected attributes to be %v, received %v", attributes, expectedAttributes)
	}

	limit := selectStatement.Limit.Limit
	var expectedLimit uint32 = 10

	if expectedLimit != limit {
		t.Fatalf("Expected Limit to be %v, received %v", expectedLimit, limit)
	}
}

func TestParsesAQueryIntoAnASTWithMultipleProjectionsInCaseInsensitiveManner(t *testing.T) {
	parser, _ := parser.NewParser("SELECT NAME, LOWER(NAME) FROM ./resources", context.NewContext(context.NewFunctions(), context.NewAttributes()))
	selectStatement, _ := parser.Parse()

	totalProjections := selectStatement.Projections.Count()
	expected := 2
	if totalProjections != expected {
		t.Fatalf("Expected projection count %v, received %v", expected, totalProjections)
	}

	attributes := selectStatement.Projections.DisplayableAttributes()
	expectedAttributes := []string{"NAME", "LOWER(NAME)"}

	if !reflect.DeepEqual(attributes, expectedAttributes) {
		t.Fatalf("Expected attributes to be %v, received %v", attributes, expectedAttributes)
	}
}
