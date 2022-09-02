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

	attributes := selectStatement.Projections.DisplayableAttributes()
	expectedAttributes := []string{"name"}

	if !reflect.DeepEqual(attributes, expectedAttributes) {
		t.Fatalf("Expected attributes to be %v, received %v", attributes, expectedAttributes)
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

	attributes := selectStatement.Projections.DisplayableAttributes()
	expectedAttributes := []string{"name", "lower(name)"}

	if !reflect.DeepEqual(attributes, expectedAttributes) {
		t.Fatalf("Expected attributes to be %v, received %v", attributes, expectedAttributes)
	}
}

func TestParsesAQueryIntoAnASTWithWhereClause(t *testing.T) {
	parser, _ := NewParser("SELECT name, lower(name) from ~ where contains(lower(name), log)", context.NewContext(context.NewFunctions(), context.NewAttributes()))
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
	parser, _ := NewParser("SELECT name, upper(lower(name)) from ~ Order by 1", context.NewContext(context.NewFunctions(), context.NewAttributes()))
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
	parser, _ := NewParser("SELECT name, lower(name) from ~ Order by 2 Limit 10", context.NewContext(context.NewFunctions(), context.NewAttributes()))
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
	parser, _ := NewParser("SELECT name, lower(name) from ~ Limit 10", context.NewContext(context.NewFunctions(), context.NewAttributes()))
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
	parser, _ := NewParser("SELECT NAME, LOWER(NAME) FROM ./resources", context.NewContext(context.NewFunctions(), context.NewAttributes()))
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
