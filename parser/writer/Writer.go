package writer

import (
	"fmt"
	"io"
	"os"
)

type Writer interface {
	Write(result string) error
}

type ConsoleWriter struct {
	backingWriter io.Writer
}

func NewWriter(backingWriter io.Writer) *ConsoleWriter {
	return &ConsoleWriter{
		backingWriter: backingWriter,
	}
}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{
		backingWriter: os.Stdout,
	}
}

func (writer ConsoleWriter) Write(result string) error {
	if _, err := fmt.Fprintln(writer.backingWriter, result); err != nil {
		return err
	}
	return nil
}
