package writer

import (
	"fmt"
	"github.com/bndr/gotabulate"
	"goselect/parser/executor"
	"goselect/parser/projection"
)

type TableFormatter struct{}

func NewTableFormatter() *TableFormatter {
	return &TableFormatter{}
}

func (tableFormatter TableFormatter) Format(projections *projection.Projections, rows *executor.EvaluatingRows) string {
	var displayableRows [][]string

	iterator := rows.RowIterator()
	for iterator.HasNext() {
		var attributes []string
		for _, attribute := range iterator.Next().AllAttributes() {
			attributes = append(attributes, attribute.GetAsString())
		}
		displayableRows = append(displayableRows, attributes)
	}
	if len(displayableRows) == 0 {
		return "no records found\n"
	}
	return tableFormatter.renderContentTable(projections.DisplayableAttributes(), displayableRows) +
		tableFormatter.renderFooterTable(displayableRows)
}

func (tableFormatter TableFormatter) renderContentTable(attributes []string, displayableRows [][]string) string {
	cellSize := 100 / len(attributes)
	table := gotabulate.Create(displayableRows)
	table.SetHeaders(attributes)
	table.SetMaxCellSize(cellSize)
	table.SetWrapStrings(true)
	table.SetAlign("left")

	return table.Render("grid")
}

func (tableFormatter TableFormatter) renderFooterTable(displayableRows [][]string) string {
	table := gotabulate.Create([]string{fmt.Sprintf("Total Rows: %v", len(displayableRows))})
	table.SetHeaders([]string{""})
	table.SetAlign("left")
	return table.Render("grid")
}
