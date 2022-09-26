package cmd

import (
	"bytes"
	"goselect/parser/context"
	"strings"
	"testing"
)

func TestFunctionsUnsorted(t *testing.T) {
	rootCmd.SetArgs([]string{"listFunctions"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()
	contents := buffer.String()

	aliasesByFunction := context.NewFunctions().AllFunctionsWithAliases()
	for function, aliases := range aliasesByFunction {
		if !strings.Contains(contents, function) {
			t.Fatalf("Expected function %v to be contained in the supported functions but was not, received %v", function, contents)
		}
		for _, alias := range aliases {
			if !strings.Contains(contents, alias) {
				t.Fatalf("Expected alias %v to be contained in the supported functions but was not, received %v", alias, contents)
			}
		}
	}
}

func TestFunctionsSorted(t *testing.T) {
	rootCmd.SetArgs([]string{"listFunctions", "--sorted", "true"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()
	contents := buffer.String()

	aliasesByFunction := context.NewFunctions().AllFunctionsWithAliases()
	for function, aliases := range aliasesByFunction {
		if !strings.Contains(contents, function) {
			t.Fatalf("Expected function %v to be contained in the supported functions but was not, received %v", function, contents)
		}
		for _, alias := range aliases {
			if !strings.Contains(contents, alias) {
				t.Fatalf("Expected alias %v to be contained in the supported functions but was not, received %v", alias, contents)
			}
		}
	}
}

func TestFunctionsSortedFalse(t *testing.T) {
	rootCmd.SetArgs([]string{"listFunctions", "--sorted", "false"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()
	contents := buffer.String()

	aliasesByFunction := context.NewFunctions().AllFunctionsWithAliases()
	for function, aliases := range aliasesByFunction {
		if !strings.Contains(contents, function) {
			t.Fatalf("Expected function %v to be contained in the supported functions but was not, received %v", function, contents)
		}
		for _, alias := range aliases {
			if !strings.Contains(contents, alias) {
				t.Fatalf("Expected alias %v to be contained in the supported functions but was not, received %v", alias, contents)
			}
		}
	}
}

func TestFunctionsUnSortedWithInvalidValue(t *testing.T) {
	rootCmd.SetArgs([]string{"listFunctions", "--sorted", "unknown"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()
	contents := buffer.String()

	aliasesByFunction := context.NewFunctions().AllFunctionsWithAliases()
	for function, aliases := range aliasesByFunction {
		if !strings.Contains(contents, function) {
			t.Fatalf("Expected function %v to be contained in the supported functions but was not, received %v", function, contents)
		}
		for _, alias := range aliases {
			if !strings.Contains(contents, alias) {
				t.Fatalf("Expected alias %v to be contained in the supported functions but was not, received %v", alias, contents)
			}
		}
	}
}
