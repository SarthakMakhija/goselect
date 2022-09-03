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
	tokens.Add(tokenizer.NewToken(tokenizer.Order, "order"))
	tokens.Add(tokenizer.NewToken(tokenizer.By, "by"))

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
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "log"))
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
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "log"))

	_, err := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	if err == nil {
		t.Fatalf("Expected an error for invalid where clause but received none")
	}
}

func TestEvaluatesWhereClauseWithContains(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "contains"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "\"Dummylog\""))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "log"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	functions := context.NewFunctions()
	where, _ := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))
	value, _ := where.EvaluateWith(nil, functions)

	if value != true {
		t.Fatalf("Expected where clause to evaluate to true but it did not")
	}
}

func TestEvaluatesWhereClauseEq1(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "substr"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "Dummylog"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "3"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "umm"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	functions := context.NewFunctions()
	where, _ := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))
	value, _ := where.EvaluateWith(nil, functions)

	if value != true {
		t.Fatalf("Expected where clause to evaluate to true but it did not")
	}
}

func TestEvaluatesWhereClauseEq2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "add"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "2"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "3"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "5"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	functions := context.NewFunctions()
	where, _ := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))
	value, _ := where.EvaluateWith(nil, functions)

	if value != true {
		t.Fatalf("Expected where clause to evaluate to true but it did not")
	}
}

func TestEvaluatesWhereClauseEq3(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "div"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "5"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ".2"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	functions := context.NewFunctions()
	where, _ := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))
	value, _ := where.EvaluateWith(nil, functions)

	if value != true {
		t.Fatalf("Expected where clause to evaluate to true but it did not")
	}
}

func TestEvaluatesWhereClauseWithAnErrorGivenWhereClauseContainsAggregateFunction(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "count"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	functions := context.NewFunctions()
	_, err := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error clause given an aggregate function inside where clause, but received none")
	}
}

func TestEvaluatesWhereClauseWithAnErrorGivenAFunctionOtherThanWhereClauseSupportedFunctionIsUsed1(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "add"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "2"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "3"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	functions := context.NewFunctions()
	_, err := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error clause given an supported function in where clause")
	}
}

func TestEvaluatesWhereClauseWithAnErrorGivenAFunctionOtherThanWhereClauseSupportedFunctionIsUsed2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "substr"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "\"Dummylog\""))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "3"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	functions := context.NewFunctions()
	_, err := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error clause given an supported function in where clause")
	}
}
