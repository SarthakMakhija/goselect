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

func TestDescribeAFunction(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"describe", "--term", "lower"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()
	contents := buffer.String()
	expectedDescription := context.NewFunctions().DescriptionOf("lower")

	if !strings.Contains(contents, expectedDescription) {
		t.Fatalf("Expected %v to be contained in the description of lower but was not, received %v", expectedDescription, contents)
	}
}

func TestDescribeAnAttribute(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"describe", "--term", "userId"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()
	contents := buffer.String()
	expectedDescription := context.NewAttributes().DescriptionOf("userId")

	if !strings.Contains(contents, expectedDescription) {
		t.Fatalf("Expected %v to be contained in the description of userId but was not, received %v", expectedDescription, contents)
	}
}

func TestInvalidTerm(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"describe", "--term", "unknown"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()
	contents := buffer.String()
	expected := cmd.ErrorMessageInvalidTerm

	if !strings.Contains(contents, expected) {
		t.Fatalf("Expected an error %v while trying to get the description of %v, received %v", cmd.ErrorMessageInvalidTerm, "unknown", contents)
	}
}

func TestBlankTerm(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"describe", "--term", ""})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()
	contents := buffer.String()
	expected := cmd.ErrorMessageEmptyTerm

	if !strings.Contains(contents, expected) {
		t.Fatalf("Expected an error %v while trying to get the description without a term value, received %v", cmd.ErrorMessageEmptyTerm, contents)
	}
}
