package cmd

import (
	"bytes"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"sort"
)

func newListTimeFormatsCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "listTimeFormats",
		Aliases: []string{"formats", "fmts"},
		Short:   "List the date/time formats supported by goselect",
		Example: `
1. goselect listTimeFormats
2. goselect formats
3. goselect fmts
`,
		Run: func(cmd *cobra.Command, args []string) {
			buffer := new(bytes.Buffer)

			tableWriter := table.NewWriter()
			tableWriter.SetOutputMirror(buffer)
			tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
			tableWriter.Style().Options.SeparateColumns = true

			appendHeader := func() {
				tableWriter.AppendHeader(table.Row{"Format", "Id"})
			}
			appendFormat := func(format context.FormatDefinition) {
				tableWriter.AppendRow(table.Row{format.Format, format.Id})
			}
			sortFormats := func(formatDefinitions map[string]context.FormatDefinition) []string {
				formats := make([]string, 0, len(formatDefinitions))
				for format := range formatDefinitions {
					formats = append(formats, format)
				}
				sort.Strings(formats)
				return formats
			}

			formatDefinitionById := context.SupportedFormats()
			appendHeader()
			for _, format := range sortFormats(formatDefinitionById) {
				appendFormat(formatDefinitionById[format])
			}
			tableWriter.Render()
			cmd.Print(buffer.String())
		},
	}
}

func init() {
	rootCmd.AddCommand(newListTimeFormatsCommand())
}
