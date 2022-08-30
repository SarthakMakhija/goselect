package writer

import (
	"fmt"
	"goselect/parser/context"
	"goselect/parser/projection"
	"io"
	"os"
	"strings"
)

type Writer interface {
	Write(projections *projection.Projections, rows [][]context.Value) error
}

type JsonWriter struct {
	backingWriter io.Writer
}

func NewJsonWriter(backingWriter io.Writer) *JsonWriter {
	return &JsonWriter{
		backingWriter: backingWriter,
	}
}

func NewJsonConsoleWriter() *JsonWriter {
	return &JsonWriter{
		backingWriter: os.Stdout,
	}
}

func (writer JsonWriter) Write(projections *projection.Projections, rows [][]context.Value) error {
	attributes := projections.DisplayableAttributes()

	attributeNameAsString := func(attributeIndex int) string {
		var value strings.Builder
		value.WriteString("\"")
		value.WriteString(attributes[attributeIndex])
		value.WriteString("\"")

		return value.String()
	}
	attributeValueAsString := func(attributeValue context.Value) string {
		var value strings.Builder
		value.WriteString("\"")
		value.WriteString(attributeValue.GetAsString())
		value.WriteString("\"")

		return value.String()
	}
	buildJson := func() string {
		var json strings.Builder
		json.WriteString("[")

		for rowIndex, row := range rows {
			json.WriteString("{")
			for attributeIndex, attributeValue := range row {
				json.WriteString(attributeNameAsString(attributeIndex))
				json.WriteString(" : ")
				json.WriteString(attributeValueAsString(attributeValue))
				if attributeIndex != len(row)-1 {
					json.WriteString(", ")
				}
			}
			json.WriteString("}")
			if rowIndex != len(rows)-1 {
				json.WriteString(", ")
			}
		}
		json.WriteString("]")
		return json.String()
	}

	if _, err := fmt.Fprint(writer.backingWriter, buildJson()); err != nil {
		return err
	}
	return nil
}
