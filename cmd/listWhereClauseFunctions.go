package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"sort"
	"strings"
)

var listWhereClauseFunctionsCmd = &cobra.Command{
	Use:   "listWhereClauseFunctions",
	Short: "List all the functions supported by goselect in 'where' clause",
	Long:  `List all the functions along with their aliases supported by goselect in 'where' clause`,
	Run: func(cmd *cobra.Command, args []string) {
		asString := func(aliases []string) string {
			return strings.Join(aliases, ", ")
		}
		printHeader := func() {
			fmt.Printf("%v%-18v %-12v\n", headerColor, "Function", "Aliases")
		}
		printFunction := func(function string, aliases []string) {
			fmt.Printf("%v%-18v %-18v\n", contentColor, function, asString(aliases))
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
		} else {
			for _, function := range sortFunctions(aliasesByFunction) {
				printFunction(function, aliasesByFunction[function])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listWhereClauseFunctionsCmd)
	listWhereClauseFunctionsCmd.PersistentFlags().Bool("sorted", false, "display the functions supported 'where' clause in sorted order. Use --sorted=true or --sorted=false")
}
