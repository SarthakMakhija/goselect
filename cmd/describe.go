package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"os"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe an attribute or a function",
	Run: func(cmd *cobra.Command, args []string) {
		tableWriter := table.NewWriter()
		tableWriter.SetOutputMirror(os.Stdout)
		tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
		tableWriter.Style().Options.SeparateColumns = true

		lookFor, _ := cmd.Flags().GetString("term")
		errorColor := "\033[31m"

		if len(lookFor) == 0 {
			fmt.Println(errorColor, "term is mandatory. please use --term=<attribute> or --term=<function>")
			return
		}

		attributes, functions := context.NewAttributes(), context.NewFunctions()
		if attributes.IsASupportedAttribute(lookFor) {
			tableWriter.AppendHeader(table.Row{"Attribute", "Description"})
			tableWriter.AppendRow(table.Row{lookFor, attributes.DescriptionOf(lookFor)})
			tableWriter.SetColumnConfigs([]table.ColumnConfig{{Name: "Description", WidthMax: 150}})
			tableWriter.Render()
			return
		}
		if functions.IsASupportedFunction(lookFor) {
			tableWriter.AppendHeader(table.Row{"Function", "Description"})
			tableWriter.AppendRow(table.Row{lookFor, functions.DescriptionOf(lookFor)})
			tableWriter.SetColumnConfigs([]table.ColumnConfig{{Name: "Description", WidthMax: 150}})
			tableWriter.Render()
			return
		}
		fmt.Println(errorColor, "unsupported attribute or function")
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
	describeCmd.PersistentFlags().String("term", "", "pass the name of an attribute or a function. Use --term=<attribute> or --term=<function>")
}
