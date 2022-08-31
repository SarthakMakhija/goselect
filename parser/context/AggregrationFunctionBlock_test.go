package context

import (
	"testing"
)

func TestCount(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("count")

	state, _ := allFunctions.ExecuteAggregate("count", initialState, StringValue("arg1"))
	state, _ = allFunctions.ExecuteAggregate("count", state, StringValue("arg1"))
	state, _ = allFunctions.ExecuteAggregate("count", state, StringValue("arg1"))

	actualValue := allFunctions.FinalValue("count", state).GetAsString()
	if actualValue != "3" {
		t.Fatalf("Expected count to be %v, received %v", "3", actualValue)
	}
}

func TestAverage(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("average")

	state, _ := allFunctions.ExecuteAggregate("average", initialState, IntValue(10))
	state, _ = allFunctions.ExecuteAggregate("average", state, IntValue(11))
	state, _ = allFunctions.ExecuteAggregate("average", state, IntValue(11))

	actualValue := allFunctions.FinalValue("average", state).GetAsString()
	if actualValue != "10.67" {
		t.Fatalf("Expected count to be %v, received %v", "10.67", actualValue)
	}
}

func TestAverageGivenANonNumericParameter(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("average")

	_, err := allFunctions.ExecuteAggregate("average", initialState, StringValue("a"))
	if err == nil {
		t.Fatalf("Expected an error on running average with a non numeric parameter")
	}
}
