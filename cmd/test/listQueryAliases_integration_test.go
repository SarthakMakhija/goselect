//go:build integration
// +build integration

package test

import (
	"bytes"
	"goselect/cmd"
	"goselect/parser/alias"
	"os"
	"strings"
	"testing"
)

func TestQueryAliases(t *testing.T) {
	queryAlias := alias.NewQueryAlias()
	_ = queryAlias.Add(alias.Alias{Query: "select * from .", Alias: "ls"})
	defer os.Remove(queryAlias.FilePath)

	cmd.GetRootCommand().SetArgs([]string{"listQueryAliases"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()
	contents := buffer.String()

	aliases, _ := queryAlias.All()
	for alias, query := range aliases {
		if !strings.Contains(contents, alias) {
			t.Fatalf("Expected alias %v to be contained in the query alises but was not, received %v", alias, contents)
		}
		if !strings.Contains(contents, query) {
			t.Fatalf("Expected query %v to be contained in the query alises but was not, received %v", alias, contents)
		}
	}
}

func TestQueryAliasesInNonExistingEmptyFile(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"listQueryAliases"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()
	contents := buffer.String()

	if !strings.Contains(contents, cmd.NoAliases) {
		t.Fatalf("Expected a message %v, received %v", cmd.NoAliases, contents)
	}
}
