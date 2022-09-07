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

	state, _ := allFunctions.ExecuteAggregate("countd", initialState, trueBooleanValue)
	state, _ = allFunctions.ExecuteAggregate("countd", state, trueBooleanValue)
	state, _ = allFunctions.ExecuteAggregate("countd", state, trueBooleanValue)

	finalValue, _ := allFunctions.FinalValue("countd", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "1" {
		t.Fatalf("Expected count distinct to be %v, received %v", "1", actualValue)
	}
}

func TestCountDistinct3(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("countd")

	state, _ := allFunctions.ExecuteAggregate("countd", initialState, trueBooleanValue)
	state, _ = allFunctions.ExecuteAggregate("countd", state, trueBooleanValue)
	state, _ = allFunctions.ExecuteAggregate("countd", state, falseBooleanValue)

	finalValue, _ := allFunctions.FinalValue("countd", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "2" {
		t.Fatalf("Expected count distinct to be %v, received %v", "2", actualValue)
	}
}

func TestSum(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("sum")

	state, _ := allFunctions.ExecuteAggregate("sum", initialState, IntValue(10))
	state, _ = allFunctions.ExecuteAggregate("sum", state, IntValue(11))
	state, _ = allFunctions.ExecuteAggregate("sum", state, IntValue(11))

	finalValue, _ := allFunctions.FinalValue("sum", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "32.00" {
		t.Fatalf("Expected sum to be %v, received %v", "32.00", actualValue)
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
		t.Fatalf("Expected average to be %v, received %v", "10.67", actualValue)
	}
}

func TestMin1(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("min")

	state, _ := allFunctions.ExecuteAggregate("min", initialState, IntValue(10))
	state, _ = allFunctions.ExecuteAggregate("min", state, IntValue(15))
	state, _ = allFunctions.ExecuteAggregate("min", state, IntValue(8))

	finalValue, _ := allFunctions.FinalValue("min", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "8" {
		t.Fatalf("Expected min to be %v, received %v", "8", actualValue)
	}
}

func TestMin2(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("min")

	state, _ := allFunctions.ExecuteAggregate("min", initialState, StringValue("name"))
	state, _ = allFunctions.ExecuteAggregate("min", state, StringValue("abc"))
	state, _ = allFunctions.ExecuteAggregate("min", state, StringValue("pqr"))

	finalValue, _ := allFunctions.FinalValue("min", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "abc" {
		t.Fatalf("Expected min to be %v, received %v", "abc", actualValue)
	}
}

func TestSumGivenANonNumericParameter(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("sum")

	_, err := allFunctions.ExecuteAggregate("sum", initialState, StringValue("a"))
	if err == nil {
		t.Fatalf("Expected an error on running sum with a non numeric parameter")
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
