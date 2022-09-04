package writer

import (
	"fmt"
	"goselect/parser/executor"
	"goselect/parser/projection"
	"strings"
)

type TableFormatter struct{}

func NewTableFormatter() *TableFormatter {
	return &TableFormatter{}
}

func (tableFormatter TableFormatter) Format(projections *projection.Projections, rows *executor.EvaluatingRows) string {
	var result = new(strings.Builder)
	headerColor := "\033[34m"
	contentColor := "\033[0m"

	for _, attributeHeader := range projections.DisplayableAttributes() {
		result.WriteString(fmt.Sprintf("%v%-24v", headerColor, attributeHeader))
	}
	result.WriteString("\n")
	iterator := rows.RowIterator()
	for iterator.HasNext() {
		for _, attribute := range iterator.Next().AllAttributes() {
			result.WriteString(fmt.Sprintf("%v%-24v", contentColor, attribute.GetAsString()))
		}
		result.WriteString("\n")
	}
	return result.String()
}
