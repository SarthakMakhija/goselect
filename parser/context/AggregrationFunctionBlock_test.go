package context

import (
	"testing"
)

func TestCount(t *testing.T) {
	allFunctions := NewFunctions()
	_, _ = allFunctions.Execute("count", StringValue("anyValue1"))
	_, _ = allFunctions.Execute("count", StringValue("anyValue2"))
	_, _ = allFunctions.Execute("count", StringValue("anyValue2"))

	finalState, _ := allFunctions.FinalState("count")
	actualValue := finalState.GetAsString()
	if actualValue != "3" {
		t.Fatalf("Expected count to be %v, received %v", "3", actualValue)
	}
}

func TestCountDistinct(t *testing.T) {
	allFunctions := NewFunctions()
	_, _ = allFunctions.Execute("countDistinct", StringValue("anyValue1"))
	_, _ = allFunctions.Execute("countDistinct", StringValue("anyValue2"))
	_, _ = allFunctions.Execute("countDistinct", StringValue("anyValue2"))
	_, _ = allFunctions.Execute("countDistinct", StringValue("anyValue2"))
	_, _ = allFunctions.Execute("countDistinct", StringValue("anyValue2"))
	_, _ = allFunctions.Execute("countDistinct", StringValue("anyValue2"))

	finalState, _ := allFunctions.FinalState("countDistinct")
	actualValue := finalState.GetAsString()
	if actualValue != "2" {
		t.Fatalf("Expected count distinct to be %v, received %v", "2", actualValue)
	}
}
