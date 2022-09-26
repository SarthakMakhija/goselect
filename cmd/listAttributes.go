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
1. goselect listAttributes
2. goselect attributes
3. goselect attrs
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
			sortAttributes := func(aliasesByAttribute map[string][]string) []string {
				attributes := make([]string, 0, len(aliasesByAttribute))
				for attribute := range aliasesByAttribute {
					attributes = append(attributes, attribute)
				}
				sort.Strings(attributes)
				return attributes
			}

			aliasesByAttribute := context.NewAttributes().AllAttributeWithAliases()
			appendHeader()
			for _, attribute := range sortAttributes(aliasesByAttribute) {
				appendAttribute(attribute, aliasesByAttribute[attribute])
			}
			tableWriter.Render()
			cmd.Print(buffer.String())
		},
	}
}

func init() {
	rootCmd.AddCommand(newListAttributesCommand())
}
