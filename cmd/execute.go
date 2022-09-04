package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"goselect/parser/writer"
)

var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute Select SQL query.",
	Long:  `Select SQL Query syntax: select <columns> from <source directory> [where <condition>] [order by] [limit]`,
	Run: func(cmd *cobra.Command, args []string) {
		rawQuery, err := cmd.Flags().GetString("query")
		errorColor := "\033[31m"
		if err != nil {
			fmt.Println(errorColor, err)
			return
		}
		if len(rawQuery) == 0 {
			fmt.Println(errorColor, "select query is mandatory. please use --query to specify the query.")
			return
		}
		newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
		parser, err := parser.NewParser(rawQuery, newContext)
		if err != nil {
			fmt.Println(errorColor, err)
			return
		}
		query, err := parser.Parse()
		if err != nil {
			fmt.Println(errorColor, err)
			return
		}
		rows, err := executor.NewSelectQueryExecutor(query, newContext).Execute()
		if err != nil {
			fmt.Println(errorColor, err)
			return
		}
		res := writer.NewTableFormatter().Format(query.Projections, rows)
		_ = writer.NewConsoleWriter().Write(res)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
	rootCmd.PersistentFlags().StringP("query", "q", "", "specify the query. Use --query=<query> or --q=<query>")
}
