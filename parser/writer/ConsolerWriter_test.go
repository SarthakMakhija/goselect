package writer

import (
	"bytes"
	"errors"
	"testing"
)

func TestConsoleWriter(t *testing.T) {
	backingWriter := new(bytes.Buffer)
	_ = NewWriter(backingWriter).Write("[{\"a\": \"b\"}]")

	op := backingWriter.String()
	expected := "[{\"a\": \"b\"}]\n"

	if expected != op {
		t.Fatalf("Expected console writer to write %v, received %v", expected, op)
	}
}

func TestConsoleWriterWithAnError(t *testing.T) {
	err := NewWriter(alwaysThrowErrorWriter{}).Write("[{\"a\": \"b\"}]")

	if err == nil {
		t.Fatalf("Expected an error while writing using console writer but received none")
	}
}

type alwaysThrowErrorWriter struct {
}

func (a alwaysThrowErrorWriter) Write(p []byte) (n int, err error) {
	return -1, errors.New("test throws an error")
}
