package cmd

import (
	"bytes"
	"goselect/parser/context"
	"strings"
	"testing"
)

func TestAttributesUnsorted(t *testing.T) {
	rootCmd.SetArgs([]string{"listAttributes"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()
	contents := buffer.String()

	aliasesByAttribute := context.NewAttributes().AllAttributeWithAliases()
	for attribute, aliases := range aliasesByAttribute {
		if !strings.Contains(contents, attribute) {
			t.Fatalf("Expected attribute %v to be contained in the supported attributes but was not, received %v", attribute, contents)
		}
		for _, alias := range aliases {
			if !strings.Contains(contents, alias) {
				t.Fatalf("Expected alias %v to be contained in the supported attributes but was not, received %v", alias, contents)
			}
		}
	}
}

func TestAttributesSorted(t *testing.T) {
	rootCmd.SetArgs([]string{"listAttributes", "--sorted", "true"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()
	contents := buffer.String()

	aliasesByAttribute := context.NewAttributes().AllAttributeWithAliases()
	for attribute, aliases := range aliasesByAttribute {
		if !strings.Contains(contents, attribute) {
			t.Fatalf("Expected attribute %v to be contained in the supported attributes but was not, received %v", attribute, contents)
		}
		for _, alias := range aliases {
			if !strings.Contains(contents, alias) {
				t.Fatalf("Expected alias %v to be contained in the supported attributes but was not, received %v", alias, contents)
			}
		}
	}
}

func TestAttributesSortedAsFalse(t *testing.T) {
	rootCmd.SetArgs([]string{"listAttributes", "--sorted", "false"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()
	contents := buffer.String()

	aliasesByAttribute := context.NewAttributes().AllAttributeWithAliases()
	for attribute, aliases := range aliasesByAttribute {
		if !strings.Contains(contents, attribute) {
			t.Fatalf("Expected attribute %v to be contained in the supported attributes but was not, received %v", attribute, contents)
		}
		for _, alias := range aliases {
			if !strings.Contains(contents, alias) {
				t.Fatalf("Expected alias %v to be contained in the supported attributes but was not, received %v", alias, contents)
			}
		}
	}
}

func TestAttributesUnSortedWithInvalidValue(t *testing.T) {
	rootCmd.SetArgs([]string{"listAttributes", "--sorted", "unknown"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()
	contents := buffer.String()

	aliasesByAttribute := context.NewAttributes().AllAttributeWithAliases()
	for attribute, aliases := range aliasesByAttribute {
		if !strings.Contains(contents, attribute) {
			t.Fatalf("Expected attribute %v to be contained in the supported attributes but was not, received %v", attribute, contents)
		}
		for _, alias := range aliases {
			if !strings.Contains(contents, alias) {
				t.Fatalf("Expected alias %v to be contained in the supported attributes but was not, received %v", alias, contents)
			}
		}
	}
}
