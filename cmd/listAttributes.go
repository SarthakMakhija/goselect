package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"sort"
	"strings"
)

var listAttributesCmd = &cobra.Command{
	Use:   "listAttributes",
	Short: "List all the attributes supported by goselect.",
	Long:  `List all the attributes along with their aliases supported by goselect.`,
	Run: func(cmd *cobra.Command, args []string) {
		asString := func(aliases []string) string {
			return strings.Join(aliases, ", ")
		}
		printHeader := func() {
			fmt.Printf("%v%-14v %-12v\n", headerColor, "Attribute", "Aliases")
		}
		printAttribute := func(attribute string, aliases []string) {
			fmt.Printf("%v%-14v %-18v\n", contentColor, attribute, asString(aliases))
		}
		printAttributes := func(aliasesByAttributes map[string][]string) {
			for attribute, aliases := range aliasesByAttributes {
				printAttribute(attribute, aliases)
			}
		}
		sortAttributes := func(aliasesByAttributes map[string][]string) []string {
			attributes := make([]string, 0, len(aliasesByAttributes))
			for attribute := range aliasesByAttributes {
				attributes = append(attributes, attribute)
			}
			sort.Strings(attributes)
			return attributes
		}

		isSorted, _ := cmd.Flags().GetBool("sorted")
		aliasesByAttributes := context.NewAttributes().AllAttributeWithAliases()
		printHeader()

		if !isSorted {
			printAttributes(aliasesByAttributes)
		} else {
			for _, attribute := range sortAttributes(aliasesByAttributes) {
				printAttribute(attribute, aliasesByAttributes[attribute])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listAttributesCmd)
	listAttributesCmd.PersistentFlags().Bool("sorted", false, "Display the attributes in sorted order. Use --sorted=true or --sorted=false.")
}
