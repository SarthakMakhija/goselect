package context

import "testing"

func TestMimeTypeAsText(t *testing.T) {
	block := MimeTypeAttributeEvaluationBlock{}
	value, _ := block.evaluate("./Value.go")
	expected := "text/plain; charset=utf-8"

	if value.GetAsString() != expected {
		t.Fatalf("Expected mime type to be %v, received %v", expected, value.GetAsString())
	}
}

func TestMimeTypeAsImage(t *testing.T) {
	block := MimeTypeAttributeEvaluationBlock{}
	value, _ := block.evaluate("../../images/any_10.png")
	expected := "image/png"

	if value.GetAsString() != expected {
		t.Fatalf("Expected mime type to be %v, received %v", expected, value.GetAsString())
	}
}