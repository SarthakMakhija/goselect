//go:build integration
// +build integration

package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	rootCmd.SetArgs([]string{"version"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

	contents := buffer.String()
	expectedVersion := "v0.0.4"

	if !strings.Contains(contents, expectedVersion) {
		t.Fatalf("Expected version to be %v, received %v", expectedVersion, contents)
	}
}
