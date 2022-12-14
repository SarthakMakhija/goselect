//go:build unit
// +build unit

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

func TestWhereWithIllegalKeywordAfterWhere1(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "ok"))

	_, err := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error for invalid where clause but received none")
	}
}

func TestWhereWithIllegalKeywordAfterWhere2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "by"))

	_, err := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error for invalid where clause but received none")
	}
}

func TestWhereWithOrderKeywordAfterWhere(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "order"))

	where, _ := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	expected := "eq(1,1)"

	if expected != where.Display() {
		t.Fatalf("Expected where clause to be %v, received %v", expected, where.Display())
	}
}

func TestWhereWithLimitKeywordAfterWhere(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "limit"))

	where, _ := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	expected := "eq(1,1)"

	if expected != where.Display() {
		t.Fatalf("Expected where clause to be %v, received %v", expected, where.Display())
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

func TestWhereWithIncorrectTokenInterpretation(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.FloatingPoint, "."))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "value"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	where, _ := NewWhere(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	display := where.Display()
	expected := "eq(.,value)"

	if display != expected {
		t.Fatalf("Expected where to be %v, received %v", expected, display)
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
		t.Fatalf("Expected an error clause given an unsupported function in where clause")
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
		t.Fatalf("Expected an error clause given an unsupported function in where clause")
	}
}

func TestWhereWithoutAnyExpressionAfterWhere(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))

	functions := context.NewFunctions()
	_, err := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error clause given where clause without any expression")
	}
}

func TestWhereWithMultipleExpressionsAfterWhere(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "contains"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "\"Dummylog\""))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "log"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "eq"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "1"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	functions := context.NewFunctions()
	_, err := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error clause given where clause with multiple expressions")
	}
}

func TestWhereWithAFunctionWithInvalidWhereClause(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "contains"))

	functions := context.NewFunctions()
	_, err := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error clause given invalid where clause")
	}
}

func TestWhereWithAFunctionWithoutOpeningParentheses(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "contains"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))

	functions := context.NewFunctions()
	_, err := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error clause given where clause with improperly closed function")
	}
}

func TestEvaluatesWhereWithAFunctionContainingInsufficientParameters(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "where"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "contains"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	functions := context.NewFunctions()
	where, _ := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))
	_, err := where.EvaluateWith(nil, functions)

	if err == nil {
		t.Fatalf("Expected an error clause given where clause with insufficient parameter values")
	}
}

func TestEvaluatesWhereWithoutAnyCondition(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()

	functions := context.NewFunctions()
	where, _ := NewWhere(tokens.Iterator(), context.NewContext(functions, context.NewAttributes()))
	value, _ := where.EvaluateWith(nil, functions)

	if value != true {
		t.Fatalf("Expected where clause to evaluate to true but it did not")
	}
}
