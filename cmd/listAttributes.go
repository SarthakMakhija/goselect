package cmd

import (
	"bytes"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/context"
	"sort"
	"strings"
)

func newListAttributesCommand() *cobra.Command {
	return &cobra.Command{
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
			buffer := new(bytes.Buffer)

			tableWriter := table.NewWriter()
			tableWriter.SetOutputMirror(buffer)
			tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
			tableWriter.Style().Options.SeparateColumns = true

			asString := func(aliases []string) string {
				return strings.Join(aliases, ", ")
			}
			appendHeader := func() {
				tableWriter.AppendHeader(table.Row{"Attribute", "Aliases"})
			}
			appendAttribute := func(attribute string, aliases []string) {
				tableWriter.AppendRow(table.Row{attribute, asString(aliases)})
			}
			appendAttributes := func(aliasesByAttribute map[string][]string) {
				for attribute, aliases := range aliasesByAttribute {
					appendAttribute(attribute, aliases)
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
			appendHeader()

			if !isSorted {
				appendAttributes(aliasesByAttribute)
				tableWriter.Render()
				cmd.Print(buffer.String())
			} else {
				for _, attribute := range sortAttributes(aliasesByAttribute) {
					appendAttribute(attribute, aliasesByAttribute[attribute])
				}
				tableWriter.Render()
				cmd.Print(buffer.String())
			}
		},
	}
}

func init() {
	listAttributesCmd := newListAttributesCommand()
	rootCmd.AddCommand(listAttributesCmd)
	listAttributesCmd.PersistentFlags().Bool("sorted", true, "display the attributes in sorted order. Use --sorted=true or --sorted=false")
}
