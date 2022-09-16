package context

import "testing"

func TestIsASupportedAttribute(t *testing.T) {
	context := NewContext(nil, NewAttributes())
	isASupportedAttribute := context.IsASupportedAttribute(AttributeSize)

	if isASupportedAttribute != true {
		t.Fatalf("Expected AttributeSize to be a supported attribute but was not")
	}
}

func TestIsNotASupportedAttribute(t *testing.T) {
	context := NewContext(nil, NewAttributes())
	isASupportedAttribute := context.IsASupportedAttribute("unknown")

	if isASupportedAttribute != false {
		t.Fatalf("Expected unknown to be an un-supported attribute but was")
	}
}

func TestIsASupportedFunction(t *testing.T) {
	context := NewContext(NewFunctions(), nil)
	isASupportedFunction := context.IsASupportedFunction(FunctionNameLower)

	if isASupportedFunction != true {
		t.Fatalf("Expected FunctionNameLower to be a supported function but was not")
	}
}

func TestIsNotASupportedFunction(t *testing.T) {
	context := NewContext(NewFunctions(), nil)
	isASupportedFunction := context.IsASupportedFunction("unknown")

	if isASupportedFunction != false {
		t.Fatalf("Expected unknown to be an un-supported function but was")
	}
}

func TestFunctionContainsATag(t *testing.T) {
	context := NewContext(NewFunctions(), nil)
	containsATag := context.FunctionContainsATag("eq", "where")

	if containsATag != true {
		t.Fatalf("Expected containsATag to be true but did not contain the tag")
	}
}

func TestFunctionDoesNotContainATag(t *testing.T) {
	context := NewContext(NewFunctions(), nil)
	containsATag := context.FunctionContainsATag("lower", "where")

	if containsATag != false {
		t.Fatalf("Expected containsATag to be false but did it contained the tag")
	}
}

func TestIsAnAggregateFunction(t *testing.T) {
	context := NewContext(NewFunctions(), nil)
	isAnAggregateFunction := context.IsAnAggregateFunction(FunctionNameMin)

	if isAnAggregateFunction != true {
		t.Fatalf("Expected FunctionNameMin to be an aggregate function but was not")
	}
}

func TestIsNotAnAggregateFunction(t *testing.T) {
	context := NewContext(NewFunctions(), nil)
	isAnAggregateFunction := context.IsAnAggregateFunction(FunctionNameLower)

	if isAnAggregateFunction != false {
		t.Fatalf("Expected FunctionNameLower to be not be an aggregate function but was")
	}
}

func TestInitialStateOfAFunction(t *testing.T) {
	context := NewContext(NewFunctions(), nil)
	initialState := context.InitialState(FunctionNameMin)

	if initialState == nil {
		t.Fatalf("Expected FunctionNameMin to have an initial state but it did not")
	}
}

func TestAllFunctions(t *testing.T) {
	context := NewContext(NewFunctions(), nil)
	allFunctions := context.AllFunctions()

	if allFunctions == nil {
		t.Fatalf("Expected allFunctions to be non-nil but was nil")
	}
}
