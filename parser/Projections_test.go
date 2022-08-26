package parser

import (
	"reflect"
	"testing"
)

func TestAllColumns1(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "name"))
	tokens.add(newToken(RawString, ","))
	tokens.add(newToken(RawString, "size"))

	projections := newProjections(tokens.iterator())
	columns, _ := projections.all()
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, columns) {
		t.Fatalf("Expected columns to be %v, received %v", expected, columns)
	}
}

func TestAllColumns2(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "fName"))
	tokens.add(newToken(RawString, ","))
	tokens.add(newToken(RawString, "size"))

	projections := newProjections(tokens.iterator())
	columns, _ := projections.all()
	expected := []string{"fName", "size"}

	if !reflect.DeepEqual(expected, columns) {
		t.Fatalf("Expected columns to be %v, received %v", expected, columns)
	}
}

func TestAllColumns3(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "*"))

	projections := newProjections(tokens.iterator())
	columns, _ := projections.all()
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, columns) {
		t.Fatalf("Expected fields to be %v, received %v", expected, columns)
	}
}

func TestAllColumns4(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "*"))
	tokens.add(newToken(RawString, ","))
	tokens.add(newToken(RawString, "name"))

	projections := newProjections(tokens.iterator())
	columns, _ := projections.all()
	expected := []string{"name", "size", "name"}

	if !reflect.DeepEqual(expected, columns) {
		t.Fatalf("Expected fields to be %v, received %v", expected, columns)
	}
}

func TestAllColumnsWithAnErrorMissingComma(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "name"))
	tokens.add(newToken(RawString, "size"))

	projections := newProjections(tokens.iterator())
	_, err := projections.all()

	if err == nil {
		t.Fatalf("Expected an error on missing comma in projections but did not receive one")
	}
}