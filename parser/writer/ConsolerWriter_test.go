package writer

import (
	"bytes"
	"testing"
)

func TestConsoleWriter(t *testing.T) {
	backingWriter := new(bytes.Buffer)
	_ = NewWriter(backingWriter).Write("[{\"a\": \"b\"}]")

	op := backingWriter.String()
	expected := "[{\"a\": \"b\"}]"

	if expected != op {
		t.Fatalf("Expected console writer to write %v, received %v", expected, op)
	}
}
