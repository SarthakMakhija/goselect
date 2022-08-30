package writer

import (
	"goselect/parser/context"
	"goselect/parser/projection"
	"strings"
)

type JsonFormatter struct{}

func NewJsonFormatter() *JsonFormatter {
	return &JsonFormatter{}
}

func (jsonFormatter JsonFormatter) Format(projections *projection.Projections, rows [][]context.Value) string {

	attributeNameAsString := func(attribute string) string {
		var value strings.Builder
		value.WriteString("\"")
		value.WriteString(attribute)
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
		attributes := projections.DisplayableAttributes()
		var json = new(strings.Builder)
		jsonFormatter.begin(json)

		for rowIndex, row := range rows {
			jsonFormatter.beginRow(json)
			for attributeIndex, attributeValue := range row {
				jsonFormatter.writeAttribute(json, attributeNameAsString(attributes[attributeIndex]), attributeValueAsString(attributeValue))
				if attributeIndex != len(row)-1 {
					jsonFormatter.writeSeparator(json)
				}
			}
			jsonFormatter.closeRow(json)
			if rowIndex != len(rows)-1 {
				jsonFormatter.writeSeparator(json)
			}
		}
		jsonFormatter.end(json)
		return json.String()
	}
	return buildJson()
}

func (jsonFormatter JsonFormatter) begin(json *strings.Builder) {
	json.WriteString("[")
}

func (jsonFormatter JsonFormatter) beginRow(json *strings.Builder) {
	json.WriteString("{")
}

func (jsonFormatter JsonFormatter) writeAttribute(json *strings.Builder, attributeName, attributeValue string) {
	json.WriteString(attributeName)
	json.WriteString(" : ")
	json.WriteString(attributeValue)
}

func (jsonFormatter JsonFormatter) writeSeparator(json *strings.Builder) {
	json.WriteString(", ")
}

func (jsonFormatter JsonFormatter) closeRow(json *strings.Builder) {
	json.WriteString("}")
}

func (jsonFormatter JsonFormatter) end(json *strings.Builder) {
	json.WriteString("]")
}
