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

func TestWhereClauseFunctions(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"listWhereClauseFunctions"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()
	contents := buffer.String()

	aliasesByFunction := context.NewFunctions().AllFunctionsWithAliasesHavingTag("where")
	for function, aliases := range aliasesByFunction {
		if !strings.Contains(contents, function) {
			t.Fatalf("Expected function %v to be contained in the where clause supported functions but was not, received %v", function, contents)
		}
		for _, alias := range aliases {
			if !strings.Contains(contents, alias) {
				t.Fatalf("Expected alias %v to be contained in the where clause supported functions but was not, received %v", alias, contents)
			}
		}
	}
}
