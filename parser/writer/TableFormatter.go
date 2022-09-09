package writer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"goselect/parser/executor"
	"goselect/parser/projection"
)

const (
	minWidth          = 10
	maxAvailableWidth = 600
)

type TableFormatter struct {
	tableWriter table.Writer
}

func NewTableFormatter() *TableFormatter {
	var buffer bytes.Buffer
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(bufio.NewWriter(&buffer))
	tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
	tableWriter.Style().Options.SeparateColumns = true
	return &TableFormatter{
		tableWriter: tableWriter,
	}
}

func (tableFormatter *TableFormatter) Format(projections *projection.Projections, rows *executor.EvaluatingRows) string {
	tableFormatter.addHeader(projections)
	tableFormatter.addContent(rows)
	tableFormatter.addFooter(rows)

	return tableFormatter.tableWriter.Render()
}

func (tableFormatter *TableFormatter) addHeader(projections *projection.Projections) {
	maxWidth := maxAvailableWidth / projections.Count()

	var attributes []interface{}
	var columnConfigs []table.ColumnConfig

	for _, headerAttribute := range projections.DisplayableAttributes() {
		attributes = append(attributes, headerAttribute)
		columnConfigs = append(columnConfigs, table.ColumnConfig{WidthMin: minWidth, WidthMax: maxWidth, Name: headerAttribute})
	}
	tableFormatter.tableWriter.AppendHeader(attributes)
	tableFormatter.tableWriter.SetColumnConfigs(columnConfigs)
}

func (tableFormatter *TableFormatter) addContent(rows *executor.EvaluatingRows) {
	iterator := rows.RowIterator()
	for iterator.HasNext() {
		var attributes []interface{}
		for _, attribute := range iterator.Next().AllAttributes() {
			attributes = append(attributes, attribute.GetAsString())
		}
		tableFormatter.tableWriter.AppendRow(attributes)
	}
}

func (tableFormatter *TableFormatter) addFooter(rows *executor.EvaluatingRows) {
	tableFormatter.tableWriter.AppendFooter(table.Row{fmt.Sprintf("Total Rows: %v", rows.Count())})
}
