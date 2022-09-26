package writer

import (
	"fmt"
	"goselect/parser/executor"
	"goselect/parser/projection"
	"strings"
)

type HtmlFormatter struct{}

func NewHtmlFormatter() *HtmlFormatter {
	return &HtmlFormatter{}
}

func (htmlFormatter HtmlFormatter) Format(projections *projection.Projections, rows *executor.EvaluatingRows) string {
	var result = new(strings.Builder)
	htmlFormatter.beginHtml(result)
	htmlFormatter.beginBody(result)
	htmlFormatter.beginTable(result)
	htmlFormatter.beginTableHeader(result, projections)
	htmlFormatter.beginTableContent(result, rows)
	htmlFormatter.beginFooterRow(result, projections, rows)
	htmlFormatter.closeTable(result)
	htmlFormatter.closeBody(result)
	htmlFormatter.closeHtml(result)

	return result.String()
}

func (htmlFormatter HtmlFormatter) beginHtml(html *strings.Builder) {
	html.WriteString("<html>")
}

func (htmlFormatter HtmlFormatter) beginBody(html *strings.Builder) {
	html.WriteString("<body>")
}

func (htmlFormatter HtmlFormatter) beginTable(html *strings.Builder) {
	html.WriteString("<table style=\"width:100%; border: 1px solid black\">")
}

func (htmlFormatter HtmlFormatter) beginTableHeader(html *strings.Builder, projections *projection.Projections) {
	htmlFormatter.beginRow(html)
	for _, attribute := range projections.DisplayableAttributes() {
		htmlFormatter.writeColumnHeader(html, attribute)
	}
	htmlFormatter.closeRow(html)
}

func (htmlFormatter HtmlFormatter) beginTableContent(html *strings.Builder, rows *executor.EvaluatingRows) {
	iterator := rows.RowIterator()
	for iterator.HasNext() {
		row := iterator.Next()
		htmlFormatter.beginRow(html)
		for _, attribute := range row.AllAttributes() {
			htmlFormatter.writeColumnContent(html, attribute.GetAsString())
		}
		htmlFormatter.closeRow(html)
	}
}

func (htmlFormatter HtmlFormatter) beginFooterRow(html *strings.Builder, projections *projection.Projections, rows *executor.EvaluatingRows) {
	htmlFormatter.beginRow(html)
	html.WriteString(fmt.Sprintf("<td colspan=\"%v\" style=\"border: 1px solid black\">", projections.Count()))
	html.WriteString(fmt.Sprintf("Rows: %v", rows.Count()))
	html.WriteString("</td>")
	htmlFormatter.closeRow(html)
}

func (htmlFormatter HtmlFormatter) beginRow(html *strings.Builder) {
	html.WriteString("<tr>")
}

func (htmlFormatter HtmlFormatter) writeColumnHeader(html *strings.Builder, column string) {
	html.WriteString("<th style=\"border: 1px solid black\">")
	html.WriteString(column)
	html.WriteString("</th>")
}

func (htmlFormatter HtmlFormatter) writeColumnContent(html *strings.Builder, column string) {
	html.WriteString("<td style=\"border: 1px solid black\">")
	html.WriteString(column)
	html.WriteString("</td>")
}

func (htmlFormatter HtmlFormatter) closeRow(html *strings.Builder) {
	html.WriteString("</tr>")
}

func (htmlFormatter HtmlFormatter) closeTable(html *strings.Builder) {
	html.WriteString("</table>")
}

func (htmlFormatter HtmlFormatter) closeBody(html *strings.Builder) {
	html.WriteString("</body>")
}

func (htmlFormatter HtmlFormatter) closeHtml(html *strings.Builder) {
	html.WriteString("</html>")
}
