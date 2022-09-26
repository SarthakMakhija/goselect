package cmd

import (
	"bytes"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"sort"
	"strings"
)

func newListWhereClauseFunctionsCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "listWhereClauseFunctions",
		Aliases: []string{"wherefunctions", "wherefns"},
		Short:   "List all the functions supported by goselect in 'where' clause",
		Long:    `List all the functions along with their aliases supported by goselect in 'where' clause`,
		Example: `
1. goselect listWhereClauseFunctions --sorted=true
2. goselect wherefunctions --sorted=true
3. goselect wherefns --sorted=true
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
			aliasesByFunction := context.NewFunctions().AllFunctionsWithAliasesHavingTag("where")
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
	listWhereClauseFunctionsCmd := newListWhereClauseFunctionsCommand()
	rootCmd.AddCommand(listWhereClauseFunctionsCmd)
	listWhereClauseFunctionsCmd.PersistentFlags().Bool("sorted", true, "display the functions supported 'where' clause in sorted order. Use --sorted=true or --sorted=false")
}
