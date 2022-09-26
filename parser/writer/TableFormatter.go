package writer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"goselect/parser/executor"
	"goselect/parser/projection"
)

type TableFormatter struct {
	tableWriter table.Writer
	options     *AttributeWidthOptions
}

type AttributeWidthOptions struct {
	minCharacters int
	maxCharacters int
}

func NewAttributeWidthOptions(minCharacters, maxCharacters int) *AttributeWidthOptions {
	return &AttributeWidthOptions{
		minCharacters: minCharacters,
		maxCharacters: maxCharacters,
	}
}

func NewTableFormatter() *TableFormatter {
	return NewTableFormatterWithWidthOptions(NewAttributeWidthOptions(10, 75))
}

func NewTableFormatterWithWidthOptions(attributeWidthOptions *AttributeWidthOptions) *TableFormatter {
	var buffer bytes.Buffer
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(bufio.NewWriter(&buffer))
	tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
	tableWriter.Style().Options.SeparateColumns = true

	return &TableFormatter{
		tableWriter: tableWriter,
		options:     attributeWidthOptions,
	}
}

func (tableFormatter *TableFormatter) Format(projections *projection.Projections, rows *executor.EvaluatingRows) string {
	tableFormatter.addHeader(projections)
	tableFormatter.addContent(rows)
	tableFormatter.addFooter(rows)

	return tableFormatter.tableWriter.Render()
}

func (tableFormatter *TableFormatter) addHeader(projections *projection.Projections) {
	var attributes []interface{}
	var columnConfigs []table.ColumnConfig

	for _, headerAttribute := range projections.DisplayableAttributes() {
		attributes = append(attributes, headerAttribute)
		columnConfigs = append(
			columnConfigs,
			table.ColumnConfig{
				Name:     headerAttribute,
				WidthMin: tableFormatter.options.minCharacters,
				WidthMax: tableFormatter.options.maxCharacters,
			},
		)
	}
	tableFormatter.setHeaderFooterDefaultCase()
	tableFormatter.tableWriter.SetColumnConfigs(columnConfigs)
	tableFormatter.tableWriter.AppendHeader(attributes)
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

func (tableFormatter *TableFormatter) setHeaderFooterDefaultCase() {
	tableFormatter.tableWriter.Style().Format.Header = 0
	tableFormatter.tableWriter.Style().Format.Footer = 0
}
