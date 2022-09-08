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
		attributes := context.NewAttributes()
		if attributes.IsASupportedAttribute(lookFor) {
			fmt.Println(attributes.DescriptionOf(lookFor))
		}
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
	describeCmd.PersistentFlags().String("term", "", "pass the name of an attribute or a function. Use --term=<attribute> or --term=<function>")
}
