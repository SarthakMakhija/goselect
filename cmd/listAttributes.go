package cmd

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"os"
	"sort"
	"strings"
)

var listAttributesCmd = &cobra.Command{
	Use:     "listAttributes",
	Aliases: []string{"attributes", "attrs"},
	Short:   "List all the attributes supported by goselect",
	Long:    `List all the attributes along with their aliases supported by goselect`,
	Example: `
1. goselect listAttributes --sorted=true
2. goselect attributes --sorted=true
3. goselect attrs --sorted=true
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
			tableWriter.AppendHeader(table.Row{"Attribute", "Aliases"})
		}
		printAttribute := func(attribute string, aliases []string) {
			tableWriter.AppendRow(table.Row{attribute, asString(aliases)})
		}
		printAttributes := func(aliasesByAttribute map[string][]string) {
			for attribute, aliases := range aliasesByAttribute {
				printAttribute(attribute, aliases)
			}
		}
		sortAttributes := func(aliasesByAttribute map[string][]string) []string {
			attributes := make([]string, 0, len(aliasesByAttribute))
			for attribute := range aliasesByAttribute {
				attributes = append(attributes, attribute)
			}
			sort.Strings(attributes)
			return attributes
		}

		isSorted, _ := cmd.Flags().GetBool("sorted")
		aliasesByAttribute := context.NewAttributes().AllAttributeWithAliases()
		printHeader()

		if !isSorted {
			printAttributes(aliasesByAttribute)
			tableWriter.Render()
		} else {
			for _, attribute := range sortAttributes(aliasesByAttribute) {
				printAttribute(attribute, aliasesByAttribute[attribute])
			}
			tableWriter.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(listAttributesCmd)
	listAttributesCmd.LocalFlags().Bool("sorted", true, "display the attributes in sorted order. Use --sorted=true or --sorted=false")
}
