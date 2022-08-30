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
	return buildJson()
}
