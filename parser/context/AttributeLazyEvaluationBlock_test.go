//go:build unit
// +build unit

package context

import (
	"reflect"
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

func TestContentsForATextFileHavingLessThanTwentyMbSize(t *testing.T) {
	block := ContentsAttributeLazyEvaluationBlock{maxFileSizeInBytesSupported: twentyMb}
	value := block.evaluate("../test/resources/textfiles/sample.txt")

	if reflect.DeepEqual(value, StringValue("")) {
		t.Fatalf("Expected to get contents for a plain text file having size less than 20mb but got no content")
	}
}

func TestContentsForAHiddenTextFileHavingLessThanTwentyMbSize(t *testing.T) {
	block := ContentsAttributeLazyEvaluationBlock{maxFileSizeInBytesSupported: twentyMb}
	value := block.evaluate("../test/resources/textfiles/.sample.txt")

	if reflect.DeepEqual(value, StringValue("")) {
		t.Fatalf("Expected to get contents for a plain text file having size less than 20mb but got no content")
	}
}

func TestContentsForAImageFile(t *testing.T) {
	block := ContentsAttributeLazyEvaluationBlock{maxFileSizeInBytesSupported: twentyMb}
	value := block.evaluate("../test/resources/images/where.png")

	if !reflect.DeepEqual(value, StringValue("")) {
		t.Fatalf("Expected to not get contents for a image file but got content")
	}
}

func TestContentsForANonExistingFile(t *testing.T) {
	block := ContentsAttributeLazyEvaluationBlock{maxFileSizeInBytesSupported: twentyMb}
	value := block.evaluate("no-existing")

	if !reflect.DeepEqual(value, StringValue("")) {
		t.Fatalf("Expected to not get contents for a non existing file but got content")
	}
}

func TestContentsForADirectory(t *testing.T) {
	block := ContentsAttributeLazyEvaluationBlock{maxFileSizeInBytesSupported: twentyMb}
	value := block.evaluate("../test/resources/textfiles")

	if !reflect.DeepEqual(value, StringValue("")) {
		t.Fatalf("Expected to not get contents for a directory but got content")
	}
}

func TestContentsForATextFileHavingSizeMoreThanTwentyMB(t *testing.T) {
	block := ContentsAttributeLazyEvaluationBlock{maxFileSizeInBytesSupported: 1}
	value := block.evaluate("../test/resources/textfiles/sample.txt")

	if !reflect.DeepEqual(value, StringValue("")) {
		t.Fatalf("Expected to not get contents for a file which exceed the file max size limit but got content")
	}

}
