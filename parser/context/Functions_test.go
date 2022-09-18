//go:build unit
// +build unit

package context

import "testing"

func TestAllFunctionsWithAliases(t *testing.T) {
	functions := NewFunctions()
	functionsWithAliases := functions.AllFunctionsWithAliases()

	if len(functionsWithAliases) == 0 {
		t.Fatalf("Expected AllFunctionsWithAliases to be greater than zero")
	}
}

func TestAllFunctionsWithAliasesHavingTag(t *testing.T) {
	functions := NewFunctions()
	functionsWithAliases := functions.AllFunctionsWithAliasesHavingTag("where")

	if len(functionsWithAliases) == 0 {
		t.Fatalf("Expected AllFunctionsWithAliasesHavingTag to be greater than zero")
	}
}

func TestDescriptionOfASupportedFunction(t *testing.T) {
	functions := NewFunctions()
	description := functions.DescriptionOf("lower")

	if description == "" {
		t.Fatalf("Expected description of lower to be non-blank but was blank")
	}
}

func TestDescriptionOfAnUnSupportedFunction(t *testing.T) {
	functions := NewFunctions()
	description := functions.DescriptionOf("unknown")

	if description != "" {
		t.Fatalf("Expected description of unknown to be blank but was %v", description)
	}
}

func TestFinalValueOfAggregateFunction(t *testing.T) {
	functions := NewFunctions()
	value, _ := functions.FinalValue("count", &FunctionState{Initial: Uint32Value(15), isUpdated: true}, []Value{})

	if value.uint32Value != 15 {
		t.Fatalf("Expected final value to be %v, received %v", 15, value.uint32Value)
	}
}

func TestFinalValueOfNonAggregateFunction(t *testing.T) {
	functions := NewFunctions()
	value, _ := functions.FinalValue("lower", nil, []Value{})

	if value != EmptyValue {
		t.Fatalf("Expected final value of a non-aggregate function to be an empty value but was %v", value)
	}
}
