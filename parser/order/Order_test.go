package order

import (
	"fmt"
	"goselect/parser/tokenizer"
	"reflect"
	"testing"
)

func TestOrderByAColumnInAscending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))

	order, _ := NewOrder(tokens.Iterator())
	expectedOrder := Order{
		ascendingColumns: []string{"name"},
	}

	if !reflect.DeepEqual(expectedOrder, order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2ColumnsInAscending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	order, _ := NewOrder(tokens.Iterator())
	expectedOrder := Order{
		ascendingColumns: []string{"name", "size"},
	}

	if !reflect.DeepEqual(expectedOrder, order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderByAColumnInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, _ := NewOrder(tokens.Iterator())
	expectedOrder := Order{
		descendingColumns: []string{"name"},
	}

	if !reflect.DeepEqual(expectedOrder, order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2ColumnsInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, err := NewOrder(tokens.Iterator())
	fmt.Println(err)
	expectedOrder := Order{
		descendingColumns: []string{"name", "size"},
	}

	if !reflect.DeepEqual(expectedOrder, order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}

func TestOrderBy2ColumnsOneInAscendingOtherInDescending(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))
	tokens.Add(tokenizer.NewToken(tokenizer.DescendingOrder, "desc"))

	order, err := NewOrder(tokens.Iterator())
	fmt.Println(err)
	expectedOrder := Order{
		ascendingColumns:  []string{"name"},
		descendingColumns: []string{"size"},
	}

	if !reflect.DeepEqual(expectedOrder, order) {
		t.Fatalf("Expected Order to be %v, received %v", expectedOrder, order)
	}
}
