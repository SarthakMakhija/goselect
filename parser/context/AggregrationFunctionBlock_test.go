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

func TestCountWithoutUpdatingTheInitialState(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("count")

	finalValue, _ := allFunctions.FinalValue("count", initialState, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "1" {
		t.Fatalf("Expected count to be %v, received %v", "1", actualValue)
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

func TestCountDistinctWithMissingParameter(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("countd")

	_, err := allFunctions.ExecuteAggregate("countd", initialState)
	if err == nil {
		t.Fatalf("Expected an error while running countd given insufficient parameters")
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

func TestCountDistinctWithoutUpdatingTheInitialState(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("countd")

	finalValue, _ := allFunctions.FinalValue("countd", initialState, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "1" {
		t.Fatalf("Expected countd to be %v, received %v", "1", actualValue)
	}
}

func TestSumWithMissingParameter(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("sum")

	_, err := allFunctions.ExecuteAggregate("sum", initialState)
	if err == nil {
		t.Fatalf("Expected an error while running sum given insufficient parameters")
	}
}

func TestSumFinalValueWithMissingParameterValue(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("sum")

	_, err := allFunctions.FinalValue("sum", initialState, []Value{})
	if err == nil {
		t.Fatalf("Expected an error while running sum final value given insufficient parameters")
	}
}

func TestSumWithIllegalParameterValue(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("sum")

	_, err := allFunctions.ExecuteAggregate("sum", initialState, StringValue("a"))
	if err == nil {
		t.Fatalf("Expected an error while running sum given illegal parameter value")
	}
}

func TestSumFinalValueWithIllegalParameterValue(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("sum")

	_, err := allFunctions.FinalValue("sum", initialState, []Value{StringValue("a")})
	if err == nil {
		t.Fatalf("Expected an error while running sum final value given illegal parameter value")
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

func TestSumGivenTheInitialStateIsNotUpdated(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("sum")

	finalValue, _ := allFunctions.FinalValue("sum", initialState, []Value{IntValue(32)})
	actualValue := finalValue.GetAsString()
	if actualValue != "32.00" {
		t.Fatalf("Expected sum to be %v, received %v", "32.00", actualValue)
	}
}

func TestAverageWithMissingParameterValue(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("avg")

	_, err := allFunctions.ExecuteAggregate("avg", initialState)
	if err == nil {
		t.Fatalf("Expected an error while running avg given insufficient parameters")
	}
}

func TestAverageFinalValueWithMissingParameterValue(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("avg")

	_, err := allFunctions.FinalValue("avg", initialState, []Value{})
	if err == nil {
		t.Fatalf("Expected an error while running avg final value given insufficient parameters")
	}
}

func TestAverageWithIllegalParameterValue(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("avg")

	_, err := allFunctions.ExecuteAggregate("avg", initialState, StringValue("a"))
	if err == nil {
		t.Fatalf("Expected an error while running avg given illegal parameter value")
	}
}

func TestAverageFinalValueWithIllegalParameterValue(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("avg")

	_, err := allFunctions.FinalValue("avg", initialState, []Value{StringValue("a")})
	if err == nil {
		t.Fatalf("Expected an error while running avg final value given illegal parameter value")
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

func TestAverageGivenTheInitialStateIsNotUpdated(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("average")

	finalValue, _ := allFunctions.FinalValue("average", initialState, []Value{Float64Value(10.67)})
	actualValue := finalValue.GetAsString()
	if actualValue != "10.67" {
		t.Fatalf("Expected average to be %v, received %v", "10.67", actualValue)
	}
}

func TestMinFinalStateWithMissingParameter(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("min")
	_, err := allFunctions.ExecuteAggregate("min", initialState)

	if err == nil {
		t.Fatalf("Expected an error while running min given insufficient parameters")
	}
}

func TestMinFinalValueWithMissingParameterValue(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("min")

	_, err := allFunctions.FinalValue("min", initialState, []Value{})
	if err == nil {
		t.Fatalf("Expected an error while running min final value given insufficient parameters")
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

func TestMinGivenTheInitialStateIsNotUpdated(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("min")

	finalValue, _ := allFunctions.FinalValue("min", initialState, []Value{StringValue("abc")})
	actualValue := finalValue.GetAsString()
	if actualValue != "abc" {
		t.Fatalf("Expected min to be %v, received %v", "abc", actualValue)
	}
}

func TestMaxFinalStateWithMissingParameter(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("max")
	_, err := allFunctions.ExecuteAggregate("max", initialState)

	if err == nil {
		t.Fatalf("Expected an error while running max given insufficient parameters")
	}
}

func TestMaxFinalValueWithMissingParameterValue(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("max")

	_, err := allFunctions.FinalValue("max", initialState, []Value{})
	if err == nil {
		t.Fatalf("Expected an error while running max final value given insufficient parameters")
	}
}

func TestMax1(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("max")

	state, _ := allFunctions.ExecuteAggregate("max", initialState, IntValue(10))
	state, _ = allFunctions.ExecuteAggregate("max", state, IntValue(15))
	state, _ = allFunctions.ExecuteAggregate("max", state, IntValue(8))

	finalValue, _ := allFunctions.FinalValue("max", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "15" {
		t.Fatalf("Expected max to be %v, received %v", "15", actualValue)
	}
}

func TestMax2(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("max")

	state, _ := allFunctions.ExecuteAggregate("max", initialState, StringValue("name"))
	state, _ = allFunctions.ExecuteAggregate("max", state, StringValue("abc"))
	state, _ = allFunctions.ExecuteAggregate("max", state, StringValue("pqr"))

	finalValue, _ := allFunctions.FinalValue("max", state, nil)
	actualValue := finalValue.GetAsString()
	if actualValue != "pqr" {
		t.Fatalf("Expected max to be %v, received %v", "pqr", actualValue)
	}
}

func TestMaxGivenTheInitialStateIsNotUpdated(t *testing.T) {
	allFunctions := NewFunctions()
	initialState := allFunctions.InitialState("max")

	finalValue, _ := allFunctions.FinalValue("max", initialState, []Value{StringValue("pqr")})
	actualValue := finalValue.GetAsString()
	if actualValue != "pqr" {
		t.Fatalf("Expected max to be %v, received %v", "pqr", actualValue)
	}
}
