package cmd

import (
	"bytes"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/context"
)

func newDescribeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "describe",
		Aliases: []string{"desc"},
		Short:   "Describe an attribute or a function",
		Example: `
1. goselect describe -t=fname
2. goselect desc -t=lower
`,
		Run: func(cmd *cobra.Command, args []string) {
			buffer := new(bytes.Buffer)

			tableWriter := table.NewWriter()
			tableWriter.SetOutputMirror(buffer)
			tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
			tableWriter.Style().Options.SeparateColumns = true

			lookFor, _ := cmd.Flags().GetString("term")
			errorColor := "\033[31m"

			if len(lookFor) == 0 {
				cmd.Println(errorColor, ErrorMessageEmptyTerm)
				return
			}

			attributes, functions := context.NewAttributes(), context.NewFunctions()
			if attributes.IsASupportedAttribute(lookFor) {
				tableWriter.AppendHeader(table.Row{"Attribute", "Description"})
				tableWriter.AppendRow(table.Row{lookFor, attributes.DescriptionOf(lookFor)})
				tableWriter.SetColumnConfigs([]table.ColumnConfig{{Name: "Description", WidthMax: 150}})
				tableWriter.Render()
				cmd.Print(buffer.String())
				return
			}
			if functions.IsASupportedFunction(lookFor) {
				tableWriter.AppendHeader(table.Row{"Function", "Description"})
				tableWriter.AppendRow(table.Row{lookFor, functions.DescriptionOf(lookFor)})
				tableWriter.SetColumnConfigs([]table.ColumnConfig{{Name: "Description", WidthMax: 150}})
				tableWriter.Render()
				cmd.Print(buffer.String())
				return
			}
			cmd.Println(errorColor, ErrorMessageInvalidTerm)
		},
	}
}

func init() {
	describeCmd := newDescribeCommand()
	rootCmd.AddCommand(describeCmd)
	describeCmd.PersistentFlags().StringP("term", "t", "", "pass the name of an attribute or a function. Use --term=<attribute> or --term=<function>")
}
