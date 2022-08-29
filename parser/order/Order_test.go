package order

import (
	"goselect/parser/tokenizer"
	"reflect"
	"testing"
)

func TestOrderByAAttributeInAscending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		Attributes: []AttributeRef{{ProjectionPosition: 1}},
		directions: []bool{true},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderByAnAttributeInAscendingWithExplicitAsc(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.AscendingOrder, "asc"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		Attributes: []AttributeRef{{ProjectionPosition: 1}},
		directions: []bool{true},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2AttributesInAscending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "2"))

	order, _ := NewOrder(tokens.Iterator(), 2)
	expectedOrder := Order{
		Attributes: []AttributeRef{{ProjectionPosition: 1}, {ProjectionPosition: 2}},
		directions: []bool{true, true},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderByAAttributeInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		Attributes: []AttributeRef{{ProjectionPosition: 1}},
		directions: []bool{false},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2AttributesInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "2"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator(), 2)
	expectedOrder := Order{
		Attributes: []AttributeRef{{ProjectionPosition: 1}, {ProjectionPosition: 2}},
		directions: []bool{false, false},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2AttributesOneInAscendingOtherInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "2"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator(), 2)
	expectedOrder := Order{
		Attributes: []AttributeRef{{ProjectionPosition: 1}, {ProjectionPosition: 2}},
		directions: []bool{true, false},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestThrowsAErrorGivenNoAttributeAfterOrderBy(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))

	_, err := NewOrder(tokens.Iterator(), 1)

	if err == nil {
		t.Fatalf("Expected an error when no attributes are given after order by but received none")
	}
}

func TestThrowsAErrorOrderByAttributePositionZero(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "0"))

	_, err := NewOrder(tokens.Iterator(), 1)

	if err == nil {
		t.Fatalf("Expected an error when 0 is given as the order by position")
	}
}

func TestOrderBy2AttributesWithOneAsTheProjectionPosition(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "2"))

	order, _ := NewOrder(tokens.Iterator(), 2)
	expectedOrder := Order{
		Attributes: []AttributeRef{{ProjectionPosition: 1}, {ProjectionPosition: 2}},
		directions: []bool{false, true},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}
