package executor

import (
	"goselect/parser/context"
	"goselect/parser/order"
	"goselect/parser/tokenizer"
	"testing"
)

func TestAscendingOrderWithASingleColumn(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))

	anOrder, _ := order.NewOrder(tokens.Iterator(), 1)
	rows := [][]context.Value{
		{context.StringValue("fileB")},
		{context.StringValue("fileA")},
	}
	expected := [][]context.Value{
		{context.StringValue("fileA")},
		{context.StringValue("fileB")},
	}

	ordering := newOrdering(anOrder)
	ordering.doOrder(rows)

	assertMatch(t, expected, rows)
}

func TestAscendingOrderWith2Columns(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "2"))

	anOrder, _ := order.NewOrder(tokens.Iterator(), 2)
	rows := [][]context.Value{
		{context.StringValue("fileB"), context.IntValue(10)},
		{context.StringValue("fileA"), context.IntValue(20)},
		{context.StringValue("fileA"), context.IntValue(30)},
	}
	expected := [][]context.Value{
		{context.StringValue("fileA"), context.IntValue(20)},
		{context.StringValue("fileA"), context.IntValue(30)},
		{context.StringValue("fileB"), context.IntValue(10)},
	}

	ordering := newOrdering(anOrder)
	ordering.doOrder(rows)

	assertMatch(t, expected, rows)
}

func TestDescendingOrderWithASingleColumn(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "desc"))

	anOrder, _ := order.NewOrder(tokens.Iterator(), 1)
	rows := [][]context.Value{
		{context.StringValue("fileA")},
		{context.StringValue("fileB")},
	}
	expected := [][]context.Value{
		{context.StringValue("fileB")},
		{context.StringValue("fileA")},
	}

	ordering := newOrdering(anOrder)
	ordering.doOrder(rows)

	assertMatch(t, expected, rows)
}

func TestDescendingOrderWith2Columns(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "desc"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "2"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "desc"))

	anOrder, _ := order.NewOrder(tokens.Iterator(), 2)
	rows := [][]context.Value{
		{context.StringValue("fileB"), context.IntValue(10)},
		{context.StringValue("fileA"), context.IntValue(20)},
		{context.StringValue("fileA"), context.IntValue(30)},
	}
	expected := [][]context.Value{
		{context.StringValue("fileB"), context.IntValue(10)},
		{context.StringValue("fileA"), context.IntValue(30)},
		{context.StringValue("fileA"), context.IntValue(20)},
	}

	ordering := newOrdering(anOrder)
	ordering.doOrder(rows)

	assertMatch(t, expected, rows)
}

func TestMixOrderWith2Columns(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "asc"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "2"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "desc"))

	anOrder, _ := order.NewOrder(tokens.Iterator(), 2)
	rows := [][]context.Value{
		{context.StringValue("fileB"), context.IntValue(10)},
		{context.StringValue("fileA"), context.IntValue(20)},
		{context.StringValue("fileA"), context.IntValue(30)},
	}
	expected := [][]context.Value{
		{context.StringValue("fileA"), context.IntValue(30)},
		{context.StringValue("fileA"), context.IntValue(20)},
		{context.StringValue("fileB"), context.IntValue(10)},
	}

	ordering := newOrdering(anOrder)
	ordering.doOrder(rows)

	assertMatch(t, expected, rows)
}
