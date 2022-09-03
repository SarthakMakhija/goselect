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

	finalValue, _ := allFunctions.FinalValue("count", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "3" {
		t.Fatalf("Expected count to be %v, received %v", "3", actualValue)
	}
}

func TestCountDistinct1(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("countd")

	state, _ := allFunctions.ExecuteAggregate("countd", initialState, IntValue(10))
	state, _ = allFunctions.ExecuteAggregate("countd", state, IntValue(2))
	state, _ = allFunctions.ExecuteAggregate("countd", state, IntValue(10))

	finalValue, _ := allFunctions.FinalValue("countd", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "2" {
		t.Fatalf("Expected count distinct to be %v, received %v", "2", actualValue)
	}
}

func TestCountDistinct2(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("countd")

	state, _ := allFunctions.ExecuteAggregate("countd", initialState, BooleanValue(true))
	state, _ = allFunctions.ExecuteAggregate("countd", state, BooleanValue(true))
	state, _ = allFunctions.ExecuteAggregate("countd", state, BooleanValue(true))

	finalValue, _ := allFunctions.FinalValue("countd", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "1" {
		t.Fatalf("Expected count distinct to be %v, received %v", "1", actualValue)
	}
}

func TestCountDistinct3(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("countd")

	state, _ := allFunctions.ExecuteAggregate("countd", initialState, BooleanValue(true))
	state, _ = allFunctions.ExecuteAggregate("countd", state, BooleanValue(true))
	state, _ = allFunctions.ExecuteAggregate("countd", state, BooleanValue(false))

	finalValue, _ := allFunctions.FinalValue("countd", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "2" {
		t.Fatalf("Expected count distinct to be %v, received %v", "2", actualValue)
	}
}

func TestAverage(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("average")

	state, _ := allFunctions.ExecuteAggregate("average", initialState, IntValue(10))
	state, _ = allFunctions.ExecuteAggregate("average", state, IntValue(11))
	state, _ = allFunctions.ExecuteAggregate("average", state, IntValue(11))

	finalValue, _ := allFunctions.FinalValue("average", state, nil)
	actualValue := finalValue.GetAsString()
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
