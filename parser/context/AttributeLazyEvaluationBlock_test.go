//go:build unit
// +build unit

package context

import (
	"github.com/dustin/go-humanize"
	"os"
	"testing"
)

func TestMimeTypeAsText(t *testing.T) {
	block := MimeTypeAttributeEvaluationBlock{}
	value := block.evaluate("./Value.go")
	expected := "text/plain; charset=utf-8"

	if value.GetAsString() != expected {
		t.Fatalf("Expected mime type to be %v, received %v", expected, value.GetAsString())
	}
}

func TestMimeTypeAsImage(t *testing.T) {
	block := MimeTypeAttributeEvaluationBlock{}
	value := block.evaluate("../test/resources/images/where.png")
	expected := "image/png"

	if value.GetAsString() != expected {
		t.Fatalf("Expected mime type to be %v, received %v", expected, value.GetAsString())
	}
}

func TestMimeTypeForANonExistingFile(t *testing.T) {
	block := MimeTypeAttributeEvaluationBlock{}
	value := block.evaluate("non-existent")

	if value.GetAsString() != "NA" {
		t.Fatalf("Expected %v while determining the mime type of a non-existent file, received %v", "NA", value.GetAsString())
	}
}

func TestContentsForATextFile(t *testing.T) {
	block := ContentsAttributeLazyEvaluationBlock{}
	value := block.evaluate("../test/resources/TestResultsWithProjections/single/TestResultsWithProjections_A.txt")
	expected := "Sample content for a test named TestResultsWithProjections"

	if value.GetAsString() != expected {
		t.Fatalf("Expected contents to be %v, received %v", expected, value.GetAsString())
	}
}

func TestContentsForAImage(t *testing.T) {
	block := ContentsAttributeLazyEvaluationBlock{}
	value := block.evaluate("../test/resources/images/where.png")

	if value.GetAsString() != "" {
		t.Fatalf("Expected contents to be %v, received %v", "", value.GetAsString())
	}
}

func TestContentsForANonExistentFile(t *testing.T) {
	block := ContentsAttributeLazyEvaluationBlock{}
	value := block.evaluate("../non-existent")

	if value.GetAsString() != "" {
		t.Fatalf("Expected contents to be %v, received %v", "", value.GetAsString())
	}
}

func TestContentsForAFileExceedingSizeLimits(t *testing.T) {
	file := createFileOfSize(contentsEvaluationSize)
	block := ContentsAttributeLazyEvaluationBlock{}
	value := block.evaluate(file.Name())

	if value.GetAsString() != "" {
		t.Fatalf("Expected contents to be %v, received %v", "", value.GetAsString())
	}
}

func createFileOfSize(size string) *os.File {
	directoryName, _ := os.MkdirTemp(".", "temp")
	file, _ := os.CreateTemp(directoryName, "temp-file")
	defer os.RemoveAll(directoryName)

	sizeInBytes, _ := humanize.ParseBytes(size)
	content := make([]byte, sizeInBytes)
	for index := 0; index < int(sizeInBytes); index++ {
		content[index] = byte(100)
	}
	_, _ = file.Write(content)
	return file
}
