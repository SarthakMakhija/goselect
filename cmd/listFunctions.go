package cmd

import (
	"bytes"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"sort"
	"strings"
)

func newListFunctionsCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "listFunctions",
		Aliases: []string{"functions", "fns"},
		Short:   "List all the functions supported by goselect",
		Long:    `List all the functions along with their aliases supported by goselect`,
		Example: `
1. goselect listFunctions --sorted=true
2. goselect functions --sorted=true
3. goselect fns --sorted=true
`,
		Run: func(cmd *cobra.Command, args []string) {
			buffer := new(bytes.Buffer)

			tableWriter := table.NewWriter()
			tableWriter.SetOutputMirror(buffer)
			tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
			tableWriter.Style().Options.SeparateColumns = true

			asString := func(aliases []string) string {
				return strings.Join(aliases, ", ")
			}
			appendHeader := func() {
				tableWriter.AppendHeader(table.Row{"Function", "Aliases"})
			}
			appendFunction := func(function string, aliases []string) {
				tableWriter.AppendRow(table.Row{function, asString(aliases)})
			}
			appendFunctions := func(aliasesByFunction map[string][]string) {
				for function, aliases := range aliasesByFunction {
					appendFunction(function, aliases)
				}
			}
			sortFunctions := func(aliasesByFunction map[string][]string) []string {
				functions := make([]string, 0, len(aliasesByFunction))
				for function := range aliasesByFunction {
					functions = append(functions, function)
				}
				sort.Strings(functions)
				return functions
			}

			isSorted, _ := cmd.Flags().GetBool("sorted")
			aliasesByFunction := context.NewFunctions().AllFunctionsWithAliases()
			appendHeader()

			if !isSorted {
				appendFunctions(aliasesByFunction)
				tableWriter.Render()
				cmd.Print(buffer.String())
			} else {
				for _, function := range sortFunctions(aliasesByFunction) {
					appendFunction(function, aliasesByFunction[function])
				}
				tableWriter.Render()
				cmd.Print(buffer.String())
			}
		},
	}
}

func init() {
	listFunctionsCmd := newListFunctionsCommand()
	rootCmd.AddCommand(listFunctionsCmd)
	listFunctionsCmd.PersistentFlags().Bool("sorted", true, "display the functions in sorted order. Use --sorted=true or --sorted=false")
}
