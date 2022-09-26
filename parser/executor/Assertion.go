package executor

import (
	"goselect/parser/context"
	"testing"
)

func AssertMatch(t *testing.T, expected [][]context.Value, queryResults *EvaluatingRows, skipAttributeIndices ...int) {
	contains := func(slice []int, value int) bool {
		for _, v := range slice {
			if value == v {
				return true
			}
		}
		return false
	}
	if uint32(len(expected)) != queryResults.Count() {
		t.Fatalf("Expected length of the query results to be %v, received %v", len(expected), queryResults.Count())
	}
	for rowIndex, row := range expected {
		if len(row) != queryResults.AtIndex(rowIndex).TotalAttributes() {
			t.Fatalf("Expected length of the rowAttributes in row index %v to be %v, received %v", rowIndex, len(row), queryResults.AtIndex(rowIndex).TotalAttributes())
		}
		rowAttributes := queryResults.AtIndex(rowIndex).AllAttributes()
		for attributeIndex, attributeValue := range row {
			if !contains(skipAttributeIndices, attributeIndex) && rowAttributes[attributeIndex].CompareTo(attributeValue) != 0 {
				t.Fatalf("Expected %v to match %v at row index %v, attribute index %v",
					attributeValue,
					rowAttributes[attributeIndex],
					rowIndex,
					attributeIndex,
				)
			}
		}
	}
}
