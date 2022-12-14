//go:build unit
// +build unit

package writer

import (
	"os"
	"testing"
)

func TestFileWriter(t *testing.T) {
	directory, _ := os.MkdirTemp(".", "file-writer-dir")
	filePath := directory + string(os.PathSeparator) + "results"
	defer os.RemoveAll(directory)

	writer, _ := NewFileWriter(filePath)
	_ = writer.Write("some content")

	file, _ := os.Open(filePath)
	readBytes := make([]byte, len("some content"))
	_, _ = file.Read(readBytes)

	expected := "some content"

	if string(readBytes) != expected {
		t.Fatalf("Expected file content to be %v, received %v", expected, string(readBytes))
	}
}

func TestFileWriterWithAnError(t *testing.T) {
	_, err := NewFileWriter("/non-existing-dir/results.log")

	if err == nil {
		t.Fatalf("Expected an error while creating a file writer but received none")
	}
}
