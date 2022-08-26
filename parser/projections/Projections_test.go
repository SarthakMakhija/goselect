package projections

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

	projections := newProjections(tokens.Iterator())
	expressions, _ := projections.all()
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, expressions.allColumns()) {
		t.Fatalf("Expected columns to be %v, received %v", expected, expressions.allColumns())
	}
}

func TestAllColumns2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	projections := newProjections(tokens.Iterator())
	expressions, _ := projections.all()
	expected := []string{"fName", "size"}

	if !reflect.DeepEqual(expected, expressions.allColumns()) {
		t.Fatalf("Expected columns to be %v, received %v", expected, expressions.allColumns())
	}
}

func TestAllColumns3(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "*"))

	projections := newProjections(tokens.Iterator())
	expressions, _ := projections.all()
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, expressions.allColumns()) {
		t.Fatalf("Expected fields to be %v, received %v", expected, expressions.allColumns())
	}
}

func TestAllColumns4(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "*"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))

	projections := newProjections(tokens.Iterator())
	expressions, _ := projections.all()
	expected := []string{"name", "size", "name"}

	if !reflect.DeepEqual(expected, expressions.allColumns()) {
		t.Fatalf("Expected fields to be %v, received %v", expected, expressions.allColumns())
	}
}

func TestAllColumnsWithAnErrorMissingComma(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	projections := newProjections(tokens.Iterator())
	_, err := projections.all()

	if err == nil {
		t.Fatalf("Expected an error on missing comma in projections but did not receive one")
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

	projections := newProjections(tokens.Iterator())
	expressions, _ := projections.all()

	functionAsString := "lower(upper(fName))"
	if expressions.functions()[0] != functionAsString {
		t.Fatalf("Expected function representation as %v, received %v", functionAsString, expressions.functions()[0])
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

	projections := newProjections(tokens.Iterator())
	expressions, _ := projections.all()

	oneFunctionAsString := "lower(upper(fName))"
	if expressions.functions()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, expressions.functions()[0])
	}
	otherFunctionAsString := "lower(fName)"
	if expressions.functions()[1] != otherFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", otherFunctionAsString, expressions.functions()[0])
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

	projections := newProjections(tokens.Iterator())
	expressions, _ := projections.all()

	oneFunctionAsString := "lower(upper(fName))"
	if expressions.functions()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, expressions.functions()[0])
	}
	column := "size"
	if expressions.allColumns()[1] != column {
		t.Fatalf("Expected column to be %v, received %v", column, expressions.allColumns()[1])
	}
}
