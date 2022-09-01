package where

import (
	"goselect/parser/context"
	"goselect/parser/tokenizer"
	"testing"
)

func TestWhereWithoutAnyWhereClause1(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()

	where, _ := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	expected := ""

	if expected != where.Display() {
		t.Fatalf("Expected where clause to be blank %v, received %v", "", where.Display())
	}
}

func TestWhereWithoutAnyWhereClause2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "by"))

	where, _ := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	expected := ""

	if expected != where.Display() {
		t.Fatalf("Expected where clause to be blank %v, received %v", "", where.Display())
	}
}

func TestWhere(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "contains"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, "log"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	where, _ := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	expected := "contains(name,log)"

	if expected != where.Display() {
		t.Fatalf("Expected where clause to be %v, received %v", expected, where.Display())
	}
}

func TestThrowsAnErrorWithInvalidParentheses(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "contains"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, "log"))

	_, err := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	if err == nil {
		t.Fatalf("Expected an error for invalid where clause but received none")
	}
}
