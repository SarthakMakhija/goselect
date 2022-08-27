package order

import (
	"goselect/parser/tokenizer"
	"reflect"
	"testing"
)

func TestOrderByAColumnInAscending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		AscendingColumns: []ColumnRef{{Name: "Name", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderByAColumnInAscendingWithExplicitAsc(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.AscendingOrder, "asc"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		AscendingColumns: []ColumnRef{{Name: "Name", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2ColumnsInAscending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		AscendingColumns: []ColumnRef{{Name: "Name", ProjectionPosition: -1}, {Name: "size", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderByAColumnInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		DescendingColumns: []ColumnRef{{Name: "Name", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2ColumnsInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		DescendingColumns: []ColumnRef{{Name: "Name", ProjectionPosition: -1}, {Name: "size", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2ColumnsOneInAscendingOtherInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		AscendingColumns:  []ColumnRef{{Name: "Name", ProjectionPosition: -1}},
		DescendingColumns: []ColumnRef{{Name: "size", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestThrowsAErrorGivenNoColumnAfterOrderBy(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))

	_, err := NewOrder(tokens.Iterator(), 1)

	if err == nil {
		t.Fatalf("Expected an error when no columns are given after order by but received none")
	}
}

func TestOrderBy2ColumnsWithOneAsTheProjectionPosition(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Name"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))

	order, _ := NewOrder(tokens.Iterator(), 1)
	expectedOrder := Order{
		AscendingColumns:  []ColumnRef{{Name: "", ProjectionPosition: 1}},
		DescendingColumns: []ColumnRef{{Name: "Name", ProjectionPosition: -1}},
	}

	if !reflect.DeepEqual(expectedOrder, *order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}
