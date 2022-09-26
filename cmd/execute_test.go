package cmd

import (
	"bytes"
	"fmt"
	"goselect/parser/error/messages"
	"os"
	"strings"
	"testing"
)

func TestExecutesAQuery(t *testing.T) {
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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

func TestAttemptToExecuteAnEmptyQuery(t *testing.T) {
	rootCmd.SetArgs([]string{"execute", "--query", ""})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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
	rootCmd.SetArgs([]string{"execute", "--query", "select from ."})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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
	rootCmd.SetArgs([]string{"execute", "--query", "select lower() from ."})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "--minWidth", "10"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "--maxWidth", "100", "--minWidth", "0"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "--minWidth", "6", "--maxWidth", "10"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "-f", "unknown"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

	contents := buffer.String()
	expected := fmt.Sprintf(ErrorMessageInvalidExportFormat, supportedFormats())

	if !strings.Contains(contents, expected) {
		t.Fatalf(
			"Expected an error %v while trying to export the result in an unknown export format but received %v",
			expected,
			contents,
		)
	}
}

func TestExecutesWithJsonExport(t *testing.T) {
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "-f", "json"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "-f", "html"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "-f", "table", "-p", "."})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

	contents := buffer.String()
	expected := ErrorMessageAttemptedToExportTableToFile

	if !strings.Contains(contents, expected) {
		t.Fatalf(
			"Expected an error %v while trying to export the result in table format to file %v",
			expected,
			contents,
		)
	}
}

func TestAttemptsToExecuteWithExportInANonExistingDirectory(t *testing.T) {
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "-f", "json", "-p", "/123"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

	contents := buffer.String()
	expected := "stat /123: no such file or directory"

	if !strings.Contains(contents, expected) {
		t.Fatalf(
			"Expected an error %v while trying to export the result in a non-existing directory %v",
			expected,
			contents,
		)
	}
}

func TestAttemptsToExecuteWithExportInAFileInsteadOfADirectory(t *testing.T) {
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "-f", "json", "-p", "./resources/test/TestResultsWithProjections_A.log"})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

	contents := buffer.String()
	expected := ErrorMessageExpectedFilePathToBeADirectory

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

	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "-f", "json", "-p", directoryName})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

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
	rootCmd.SetArgs([]string{"execute", "--query", "select name from ./resources/test/ order by 1", "-f", "json", "-p", withSeparator})
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)

	_ = rootCmd.Execute()

	fileName := fmt.Sprintf("%v/results.json", directoryName)
	_, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("Expected file %v to exist but received an err %v", fileName, err)
	}
}
