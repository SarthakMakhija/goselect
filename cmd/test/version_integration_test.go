//go:build integration
// +build integration

package test

import (
	"bytes"
	"goselect/cmd"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"version"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expectedVersion := "v0.0.5"

	if !strings.Contains(contents, expectedVersion) {
		t.Fatalf("Expected version to be %v, received %v", expectedVersion, contents)
	}
}
