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
1. goselect listTimeFormats --sorted=true
2. goselect formats --sorted=true
3. goselect fmts --sorted=true
`,
		Run: func(cmd *cobra.Command, args []string) {
			buffer := new(bytes.Buffer)

			tableWriter := table.NewWriter()
			tableWriter.SetOutputMirror(buffer)
			tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
			tableWriter.Style().Options.SeparateColumns = true

			printHeader := func() {
				tableWriter.AppendHeader(table.Row{"Format", "Id"})
			}
			printFormat := func(format context.FormatDefinition) {
				tableWriter.AppendRow(table.Row{format.Format, format.Id})
			}
			printFormats := func(formatDefinitions map[string]context.FormatDefinition) {
				for _, definition := range formatDefinitions {
					printFormat(definition)
				}
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
			isSorted, _ := cmd.Flags().GetBool("sorted")
			printHeader()

			if !isSorted {
				printFormats(formatDefinitionById)
				tableWriter.Render()
				cmd.Print(buffer.String())
			} else {
				for _, format := range sortFormats(formatDefinitionById) {
					printFormat(formatDefinitionById[format])
				}
				tableWriter.Render()
				cmd.Print(buffer.String())
			}
		},
	}
}

func init() {
	listTimeFormatsCmd := newListTimeFormatsCommand()
	rootCmd.AddCommand(listTimeFormatsCmd)
	listTimeFormatsCmd.PersistentFlags().Bool("sorted", true, "display the formats in sorted order. Use --sorted=true or --sorted=false")
}
