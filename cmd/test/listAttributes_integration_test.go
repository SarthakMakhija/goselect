//go:build integration
// +build integration

package test

import (
	"bytes"
	"goselect/cmd"
	"goselect/parser/context"
	"strings"
	"testing"
)

func TestAttributes(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"listAttributes"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()
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
