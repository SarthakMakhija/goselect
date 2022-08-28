package order

import (
	"goselect/parser/context"
	"goselect/parser/tokenizer"
	"reflect"
	"testing"
)

func TestOrderByAAttributeInAscending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))

	order, _ := NewOrder(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()), 1)
	expectedOrder := Order{
		AscendingAttributes: []AttributeRef{{Name: "Name", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderByAAttributeInAscendingWithExplicitAsc(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.AscendingOrder, "asc"))

	order, _ := NewOrder(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()), 1)
	expectedOrder := Order{
		AscendingAttributes: []AttributeRef{{Name: "Name", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2AttributesInAscending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	order, _ := NewOrder(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()), 1)
	expectedOrder := Order{
		AscendingAttributes: []AttributeRef{{Name: "Name", ProjectionPosition: -1}, {Name: "size", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderByAAttributeInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()), 1)
	expectedOrder := Order{
		DescendingAttributes: []AttributeRef{{Name: "Name", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2AttributesInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()), 1)
	expectedOrder := Order{
		DescendingAttributes: []AttributeRef{{Name: "Name", ProjectionPosition: -1}, {Name: "size", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2AttributesOneInAscendingOtherInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()), 1)
	expectedOrder := Order{
		AscendingAttributes:  []AttributeRef{{Name: "Name", ProjectionPosition: -1}},
		DescendingAttributes: []AttributeRef{{Name: "size", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestThrowsAErrorGivenNoAttributeAfterOrderBy(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))

	_, err := NewOrder(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()), 1)

	if err == nil {
		t.Fatalf("Expected an error when no attributes are given after order by but received none")
	}
}

func TestOrderBy2AttributesWithOneAsTheProjectionPosition(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))

	order, _ := NewOrder(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()), 1)
	expectedOrder := Order{
		AscendingAttributes:  []AttributeRef{{Name: "", ProjectionPosition: 1}},
		DescendingAttributes: []AttributeRef{{Name: "Name", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}
