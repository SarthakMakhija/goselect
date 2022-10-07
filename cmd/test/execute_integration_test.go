//go:build integration
// +build integration

package test

import (
	"bytes"
	"fmt"
	"goselect/cmd"
	"goselect/parser/error/messages"
	"os"
	"strings"
	"testing"
)

func TestExecutesAQuery(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/ order by 1"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := []string{"TestResultsWithProjections_A.log", "TestResultsWithProjections_B.log", "TestResultsWithProjections_C.txt", "log"}

	for _, name := range expected {
		if !strings.Contains(contents, name) {
			t.Fatalf(
				"Expected file name %v to be contained in the result but was not, received %v",
				name,
				contents,
			)
		}
	}
}

func TestExecutesAQueryWithNestedTraversalOff(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/ order by 1", "--nestedTraversal=false"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := []string{"log"}

	for _, name := range expected {
		if !strings.Contains(contents, name) {
			t.Fatalf(
				"Expected file name %v to be contained in the result but was not, received %v",
				name,
				contents,
			)
		}
	}
	unexpected := []string{"TestResultsWithProjections_A.log", "TestResultsWithProjections_B.log", "TestResultsWithProjections_C.txt"}

	for _, name := range unexpected {
		if strings.Contains(contents, name) {
			t.Fatalf(
				"Expected file name %v to not be contained in the result but was, received %v",
				name,
				contents,
			)
		}
	}
}

func TestAttemptToExecuteAnEmptyQuery(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", ""})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := messages.ErrorMessageEmptyQuery

	if !strings.Contains(contents, expected) {
		t.Fatalf(
			"Expected an error %v while trying to run execute command with an empty query, received %v",
			messages.ErrorMessageEmptyQuery,
			contents,
		)
	}
}

func TestAttemptToExecuteAnInvalidQuery1(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select from ."})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := messages.ErrorMessageExpectedExpressionInProjection

	if !strings.Contains(contents, expected) {
		t.Fatalf(
			"Expected an error %v while trying to run execute command with an invalid query, received %v",
			messages.ErrorMessageExpectedExpressionInProjection,
			contents,
		)
	}
}

func TestAttemptToExecuteAnInvalidQuery2(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select lower() from ."})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()

	if !strings.Contains(contents, "lower") {
		t.Fatalf(
			"Expected an error %v while trying to run execute command with an invalid query, received %v",
			"error must contain the term lower",
			contents,
		)
	}
}

func TestExecutesWithMinWidthQuery(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "--minWidth", "10"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := []string{"TestResultsWithProjections_A.l", "og", "TestResultsWithProjections_B.l", "TestResultsWithProjections_C.t", "xt"}

	for _, name := range expected {
		if !strings.Contains(contents, name) {
			t.Fatalf(
				"Expected file name %v to be contained in the result but was not, received %v",
				name,
				contents,
			)
		}
	}
}

func TestExecutesWithMaxWidthQuery(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "--maxWidth", "100", "--minWidth", "0"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := []string{"TestResultsWithProjections_A.log", "TestResultsWithProjections_B.log", "TestResultsWithProjections_C.txt"}

	for _, name := range expected {
		if !strings.Contains(contents, name) {
			t.Fatalf(
				"Expected file name %v to be contained in the result but was not, received %v",
				name,
				contents,
			)
		}
	}
}

func TestExecutesWithMinMaxWidthQuery(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "--minWidth", "6", "--maxWidth", "10"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := []string{"TestResult", "sWithProje", "ctions_A.l", "og", "ctions_B.l", "ctions_C.t", "xt"}

	for _, name := range expected {
		if !strings.Contains(contents, name) {
			t.Fatalf(
				"Expected file name %v to be contained in the result but was not, received %v",
				name,
				contents,
			)
		}
	}
}

func TestExecutesWithInvalidExportFormat(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "-f", "unknown"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := fmt.Sprintf(cmd.ErrorMessageInvalidExportFormat, cmd.SupportedExportFormats())

	if !strings.Contains(contents, expected) {
		t.Fatalf(
			"Expected an error %v while trying to export the result in an unknown export format but received %v",
			expected,
			contents,
		)
	}
}

func TestExecutesWithJsonExport(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "-f", "json"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := []string{"TestResultsWithProjections_A.log", "TestResultsWithProjections_B.log", "TestResultsWithProjections_C.txt"}

	for _, name := range expected {
		if !strings.Contains(contents, name) {
			t.Fatalf(
				"Expected file name %v to be contained in the json result but was not, received %v",
				name,
				contents,
			)
		}
	}
}

func TestExecutesWithHtmlExport(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "-f", "html"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := []string{"TestResultsWithProjections_A.log", "TestResultsWithProjections_B.log", "TestResultsWithProjections_C.txt"}

	for _, name := range expected {
		if !strings.Contains(contents, name) {
			t.Fatalf(
				"Expected file name %v to be contained in the json result but was not, received %v",
				name,
				contents,
			)
		}
	}
}

func TestAttemptsToExecuteWithTableFormatExportToAFile(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "-f", "table", "-p", "."})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := cmd.ErrorMessageAttemptedToExportTableToFile

	if !strings.Contains(contents, expected) {
		t.Fatalf(
			"Expected an error %v while trying to export the result in table format to file %v",
			expected,
			contents,
		)
	}
}

func TestAttemptsToExecuteWithExportInANonExistingDirectory(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "-f", "json", "-p", "/123"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected1 := "stat /123: no such file or directory"
	expected2 := "CreateFile /123: The system cannot find the file specified"

	if !strings.Contains(contents, expected1) && !strings.Contains(contents, expected2) {
		t.Fatalf(
			"Expected an error %v or %v while trying to export the result in a non-existing directory %v",
			expected1,
			expected2,
			contents,
		)
	}
}

func TestAttemptsToExecuteWithExportInAFileInsteadOfADirectory(t *testing.T) {
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "-f", "json", "-p", "./resources/log/TestResultsWithProjections_A.log"})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	contents := buffer.String()
	expected := cmd.ErrorMessageExpectedFilePathToBeADirectory

	if !strings.Contains(contents, expected) {
		t.Fatalf(
			"Expected an error %v while trying to export the result in a file instead of a directory %v",
			expected,
			contents,
		)
	}
}

func TestExecuteWithExportToAFileInADirectory(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "export-result")
	defer os.RemoveAll(directoryName)

	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "-f", "json", "-p", directoryName})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	fileName := fmt.Sprintf("%v/results.json", directoryName)
	_, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("Expected file %v to exist but received an err %v", fileName, err)
	}
}

func TestExecuteWithExportToAFileInADirectoryWithPathSeparator(t *testing.T) {
	directoryName, _ := os.MkdirTemp(".", "export-result")
	defer os.RemoveAll(directoryName)

	withSeparator := directoryName + string(os.PathSeparator)
	cmd.GetRootCommand().SetArgs([]string{"execute", "--query", "select name from ./resources/log/ order by 1", "-f", "json", "-p", withSeparator})
	buffer := new(bytes.Buffer)
	cmd.GetRootCommand().SetOut(buffer)

	_ = cmd.GetRootCommand().Execute()

	fileName := fmt.Sprintf("%v/results.json", directoryName)
	_, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("Expected file %v to exist but received an err %v", fileName, err)
	}
}
