package cmd

import (
	"bytes"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"goselect/parser/alias"
	"sort"
)

var listQueryAliasesCmd = &cobra.Command{
	Use:     "listQueryAliases",
	Aliases: []string{"qAliases"},
	Short:   "List all the saved query aliases",
	Example: `
1. goselect listQueryAliases
2. goselect qAliases
`,
	Run: func(cmd *cobra.Command, args []string) {
		errorColor := "\033[31m"

		buffer := new(bytes.Buffer)

		tableWriter := table.NewWriter()
		tableWriter.SetOutputMirror(buffer)
		tableWriter.SetStyle(table.StyleColoredBlackOnCyanWhite)
		tableWriter.Style().Options.SeparateColumns = true

		appendHeader := func() {
			tableWriter.AppendHeader(table.Row{"Alias", "Query"})
		}
		appendAlias := func(alias string, query string) {
			tableWriter.AppendRow(table.Row{alias, query})
		}
		sortAliases := func(queryByAlias map[string]string) []string {
			aliases := make([]string, 0, len(queryByAlias))
			for alias := range queryByAlias {
				aliases = append(aliases, alias)
			}
			sort.Strings(aliases)
			return aliases
		}

		queryAlias := alias.NewQueryAlias()
		aliases, err := queryAlias.All()
		if err != nil {
			cmd.Println(errorColor, err)
			return
		}
		appendHeader()
		for _, alias := range sortAliases(aliases) {
			appendAlias(alias, aliases[alias])
		}
		tableWriter.Render()
		cmd.Print(buffer.String())
	},
}

func init() {
	rootCmd.AddCommand(listQueryAliasesCmd)
}
