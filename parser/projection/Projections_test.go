package projection

import (
	"goselect/parser/context"
	"goselect/parser/tokenizer"
	"reflect"
	"testing"
)

func TestThrowsAnErrorWithoutAnyAttributes(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))

	_, err := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	if err == nil {
		t.Fatalf("Expected an error without any projection list but received none")
	}
}

func TestAllAttributes1(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, projections.DisplayableAttributes()) {
		t.Fatalf("Expected columns to be %v, received %v", expected, projections.DisplayableAttributes())
	}
}

func TestAllAttributes2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	expected := []string{"fName", "size"}

	if !reflect.DeepEqual(expected, projections.DisplayableAttributes()) {
		t.Fatalf("Expected columns to be %v, received %v", expected, projections.DisplayableAttributes())
	}
}

func TestAllAttributes3(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "*"))

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, projections.DisplayableAttributes()) {
		t.Fatalf("Expected fields to be %v, received %v", expected, projections.DisplayableAttributes())
	}
}

func TestAllAttributes4(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "*"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ","))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	expected := []string{"name", "size", "name"}

	if !reflect.DeepEqual(expected, projections.DisplayableAttributes()) {
		t.Fatalf("Expected fields to be %v, received %v", expected, projections.DisplayableAttributes())
	}
}

func TestAllAttributesWithAnErrorMissingComma(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "name"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "size"))

	_, err := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))

	if err == nil {
		t.Fatalf("Expected an error on missing comma in projection but did not receive one")
	}
}

func TestAllAttributesWithAFunction(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "lower"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "upper"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))

	functionAsString := "lower(upper(fName))"
	if projections.DisplayableAttributes()[0] != functionAsString {
		t.Fatalf("Expected function representation as %v, received %v", functionAsString, projections.DisplayableAttributes()[0])
	}
}

func TestAllAttributesWithAFunctionWithoutAnyParameters(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "now"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, ")"))

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))

	functionAsString := "now()"
	if projections.DisplayableAttributes()[0] != functionAsString {
		t.Fatalf("Expected function representation as %v, received %v", functionAsString, projections.DisplayableAttributes()[0])
	}
}

func TestAllAttributesWithAFunctionWithFromAsATokenAfterFunction(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "lower"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "upper"))
	tokens.Add(tokenizer.NewToken(tokenizer.OpeningParentheses, "("))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "fName"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.ClosingParentheses, ")"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "from"))

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))

	oneFunctionAsString := "lower(upper(fName))"
	if projections.DisplayableAttributes()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, projections.DisplayableAttributes()[0])
	}
}

func TestAllAttributesWith2Functions(t *testing.T) {
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

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))

	oneFunctionAsString := "lower(upper(fName))"
	if projections.DisplayableAttributes()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, projections.DisplayableAttributes()[0])
	}
	otherFunctionAsString := "lower(fName)"
	if projections.DisplayableAttributes()[1] != otherFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", otherFunctionAsString, projections.DisplayableAttributes()[0])
	}
}

func TestAllAttributesWithFunctionsAndAttributes(t *testing.T) {
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

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))

	oneFunctionAsString := "lower(upper(fName))"
	if projections.DisplayableAttributes()[0] != oneFunctionAsString {
		t.Fatalf("Expected function representation as %v, received %v", oneFunctionAsString, projections.DisplayableAttributes()[0])
	}
	column := "size"
	if projections.DisplayableAttributes()[1] != column {
		t.Fatalf("Expected attribute to be %v, received %v", column, projections.DisplayableAttributes()[1])
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

	projections, _ := NewProjections(tokens.Iterator(), context.NewContext(context.NewFunctions(), context.NewAttributes()))
	columnCount := projections.Count()
	expectedCount := 2

	if expectedCount != columnCount {
		t.Fatalf("Expected attribute Count %v, received %v", expectedCount, columnCount)
	}
}
