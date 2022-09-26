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

type FileWriter struct {
	file *os.File
}

func NewWriter(backingWriter io.Writer) *ConsoleWriter {
	return &ConsoleWriter{
		backingWriter: backingWriter,
	}
}

func NewFileWriter(filePath string) (*FileWriter, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &FileWriter{file: file}, nil
}

func (writer ConsoleWriter) Write(result string) error {
	if _, err := fmt.Fprintln(writer.backingWriter, result); err != nil {
		return err
	}
	return nil
}

func (writer FileWriter) Write(result string) error {
	if _, err := writer.file.WriteString(result); err != nil {
		return err
	}
	_ = writer.file.Sync()
	return nil
}
