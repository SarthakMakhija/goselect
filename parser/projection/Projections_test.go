package projection

import (
	"goselect/parser/tokenizer"
	"reflect"
	"testing"
)

func TestAllColumns1(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	projections, _ := newProjections(tokens.Iterator())
	expressions := projections.expressions
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, expressions.displayableColumns()) {
		t.Fatalf("Expected columns to be %v, received %v", expected, expressions.displayableColumns())
	}
}

func TestAllColumns2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	projections, _ := newProjections(tokens.Iterator())
	expressions := projections.expressions
	expected := []string{"fName", "size"}

	if !reflect.DeepEqual(expected, expressions.displayableColumns()) {
		t.Fatalf("Expected columns to be %v, received %v", expected, expressions.displayableColumns())
	}
}

func TestAllColumns3(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "*"))

	projections, _ := newProjections(tokens.Iterator())
	expressions := projections.expressions
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, expressions.displayableColumns()) {
		t.Fatalf("Expected fields to be %v, received %v", expected, expressions.displayableColumns())
	}
}

func TestAllColumns4(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "*"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))

	projections, _ := newProjections(tokens.Iterator())
	expressions := projections.expressions
	expected := []string{"name", "size", "name"}

	if !reflect.DeepEqual(expected, expressions.displayableColumns()) {
		t.Fatalf("Expected fields to be %v, received %v", expected, expressions.displayableColumns())
	}
}

func TestAllColumnsWithAnErrorMissingComma(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	_, err := newProjections(tokens.Iterator())

	if err == nil {
		t.Fatalf("Expected an error on missing comma in projection but did not receive one")
	}
}

func TestAllColumnsWithAFunction(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "lower"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "upper"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	projections, _ := newProjections(tokens.Iterator())
	expressions := projections.expressions

	functionAsString := "lower(upper(fName))"
	if expressions.displayableColumns()[0] != functionAsString {
		t.Fatalf("Expected function representation as %v, received %v", functionAsString, expressions.displayableColumns()[0])
	}
}

func TestAllColumnsWithAFunctionWithSpaceAsATokenAfterFunction(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "lower"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "upper"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, " "))

	projections, _ := newProjections(tokens.Iterator())
	expressions := projections.expressions

	oneFunctionAsString := "lower(upper(fName))"
	if expressions.displayableColumns()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, expressions.displayableColumns()[0])
	}
}

func TestAllColumnsWith2Functions(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "lower"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "upper"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "lower"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	projections, _ := newProjections(tokens.Iterator())
	expressions := projections.expressions

	oneFunctionAsString := "lower(upper(fName))"
	if expressions.displayableColumns()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, expressions.displayableColumns()[0])
	}
	otherFunctionAsString := "lower(fName)"
	if expressions.displayableColumns()[1] != otherFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", otherFunctionAsString, expressions.displayableColumns()[0])
	}
}

func TestAllColumnsWithFunctionsAndColumns(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "lower"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "upper"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.Comma, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	projections, _ := newProjections(tokens.Iterator())
	expressions := projections.expressions

	oneFunctionAsString := "lower(upper(fName))"
	if expressions.displayableColumns()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, expressions.displayableColumns()[0])
	}
	column := "size"
	if expressions.displayableColumns()[1] != column {
		t.Fatalf("Expected column to be %v, received %v", column, expressions.displayableColumns()[1])
	}
}
