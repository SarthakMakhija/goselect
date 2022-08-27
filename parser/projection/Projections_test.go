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

	projections, _ := NewProjections(tokens.Iterator())
	expressions := projections.expressions
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, expressions.DisplayableColumns()) {
		t.Fatalf("Expected columns to be %v, received %v", expected, expressions.DisplayableColumns())
	}
}

func TestAllColumns2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	projections, _ := NewProjections(tokens.Iterator())
	expressions := projections.expressions
	expected := []string{"fName", "size"}

	if !reflect.DeepEqual(expected, expressions.DisplayableColumns()) {
		t.Fatalf("Expected columns to be %v, received %v", expected, expressions.DisplayableColumns())
	}
}

func TestAllColumns3(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "*"))

	projections, _ := NewProjections(tokens.Iterator())
	expressions := projections.expressions
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, expressions.DisplayableColumns()) {
		t.Fatalf("Expected fields to be %v, received %v", expected, expressions.DisplayableColumns())
	}
}

func TestAllColumns4(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "*"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))

	projections, _ := NewProjections(tokens.Iterator())
	expressions := projections.expressions
	expected := []string{"name", "size", "name"}

	if !reflect.DeepEqual(expected, expressions.DisplayableColumns()) {
		t.Fatalf("Expected fields to be %v, received %v", expected, expressions.DisplayableColumns())
	}
}

func TestAllColumnsWithAnErrorMissingComma(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	_, err := NewProjections(tokens.Iterator())

	if err == nil {
		t.Fatalf("Expected an errors on missing comma in projection but did not receive one")
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

	projections, _ := NewProjections(tokens.Iterator())
	expressions := projections.expressions

	functionAsString := "lower(upper(fName))"
	if expressions.DisplayableColumns()[0] != functionAsString {
		t.Fatalf("Expected function representation as %v, received %v", functionAsString, expressions.DisplayableColumns()[0])
	}
}

func TestAllColumnsWithAFunctionWithFromAsATokenAfterFunction(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "lower"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "upper"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "from"))

	projections, _ := NewProjections(tokens.Iterator())
	expressions := projections.expressions

	oneFunctionAsString := "lower(upper(fName))"
	if expressions.DisplayableColumns()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, expressions.DisplayableColumns()[0])
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

	projections, _ := NewProjections(tokens.Iterator())
	expressions := projections.expressions

	oneFunctionAsString := "lower(upper(fName))"
	if expressions.DisplayableColumns()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, expressions.DisplayableColumns()[0])
	}
	otherFunctionAsString := "lower(fName)"
	if expressions.DisplayableColumns()[1] != otherFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", otherFunctionAsString, expressions.DisplayableColumns()[0])
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

	projections, _ := NewProjections(tokens.Iterator())
	expressions := projections.expressions

	oneFunctionAsString := "lower(upper(fName))"
	if expressions.DisplayableColumns()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, expressions.DisplayableColumns()[0])
	}
	column := "size"
	if expressions.DisplayableColumns()[1] != column {
		t.Fatalf("Expected column to be %v, received %v", column, expressions.DisplayableColumns()[1])
	}
}

func TestProjectionCount(t *testing.T) {
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

	projections, _ := NewProjections(tokens.Iterator())
	columnCount := projections.Count()
	expectedCount := 2

	if expectedCount != columnCount {
		t.Fatalf("Expected column Count %v, received %v", expectedCount, columnCount)
	}
}
