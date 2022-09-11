package cmd

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"os"
	"sort"
)

var listTimeFormatsCmd = &cobra.Command{
	Use:   "listTimeFormats",
	Short: "List the date/time formats supported by goselect",
	Example: `
1. goselect listTimeFormats --sorted=true
`,
	Run: func(cmd *cobra.Command, args []string) {
		tableWriter := table.NewWriter()
		tableWriter.SetOutputMirror(os.Stdout)
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
		} else {
			for _, format := range sortFormats(formatDefinitionById) {
				printFormat(formatDefinitionById[format])
			}
			tableWriter.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(listTimeFormatsCmd)
	listTimeFormatsCmd.PersistentFlags().Bool("sorted", true, "display the formats in sorted order. Use --sorted=true or --sorted=false")
}
