//go:build unit
// +build unit

package context

import (
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
