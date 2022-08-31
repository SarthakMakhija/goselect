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

	actualValue := state.Initial.GetAsString()
	if actualValue != "3" {
		t.Fatalf("Expected count to be %v, received %v", "3", actualValue)
	}
}
