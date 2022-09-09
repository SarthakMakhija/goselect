package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goselect/parser/context"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe an attribute or a function",
	Run: func(cmd *cobra.Command, args []string) {
		lookFor, _ := cmd.Flags().GetString("term")
		errorColor := "\033[31m"
		if len(lookFor) == 0 {
			fmt.Println(errorColor, "term is mandatory. please use --term=<attribute> or --term=<function>")
			return
		}
		attributes, functions := context.NewAttributes(), context.NewFunctions()
		if attributes.IsASupportedAttribute(lookFor) {
			fmt.Println(attributes.DescriptionOf(lookFor))
			return
		}
		if functions.IsASupportedFunction(lookFor) {
			fmt.Println(functions.DescriptionOf(lookFor))
			return
		}
		fmt.Println(errorColor, "unsupported attribute or function.")
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
	describeCmd.PersistentFlags().String("term", "", "pass the name of an attribute or a function. Use --term=<attribute> or --term=<function>")
}
