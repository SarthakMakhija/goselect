package cmd

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"os"
	"sort"
	"strings"
)

var listWhereClauseFunctionsCmd = &cobra.Command{
	Use:     "listWhereClauseFunctions",
	Aliases: []string{"wherefunctions", "wherefunc", "wherefns"},
	Short:   "List all the functions supported by goselect in 'where' clause",
	Long:    `List all the functions along with their aliases supported by goselect in 'where' clause`,
	Example: `
1. goselect listWhereClauseFunctions --sorted=true
2. goselect wherefunctions --sorted=true
3. goselect wherefunc --sorted=true
4. goselect wherefns --sorted=true
`,
	Run: func(cmd *cobra.Command, args []string) {
		tableWriter := table.NewWriter()
		tableWriter.SetOutputMirror(os.Stdout)
		tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
		tableWriter.Style().Options.SeparateColumns = true

		asString := func(aliases []string) string {
			return strings.Join(aliases, ", ")
		}
		printHeader := func() {
			tableWriter.AppendHeader(table.Row{"Function", "Aliases"})
		}
		printFunction := func(function string, aliases []string) {
			tableWriter.AppendRow(table.Row{function, asString(aliases)})
		}
		printFunctions := func(aliasesByFunction map[string][]string) {
			for function, aliases := range aliasesByFunction {
				printFunction(function, aliases)
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
		printHeader()

		if !isSorted {
			printFunctions(aliasesByFunction)
			tableWriter.Render()
		} else {
			for _, function := range sortFunctions(aliasesByFunction) {
				printFunction(function, aliasesByFunction[function])
			}
			tableWriter.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(listWhereClauseFunctionsCmd)
	listWhereClauseFunctionsCmd.PersistentFlags().Bool("sorted", true, "display the functions supported 'where' clause in sorted order. Use --sorted=true or --sorted=false")
}
