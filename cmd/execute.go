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
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(rawQuery) == 0 {
			fmt.Println("select query is mandatory. please use --query to specify the query.")
			return
		}
		newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
		parser, err := parser.NewParser(rawQuery, newContext)
		if err != nil {
			fmt.Println(err)
			return
		}
		query, err := parser.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}
		rows, err := executor.NewSelectQueryExecutor(query, newContext).Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
		res := writer.NewJsonFormatter().Format(query.Projections, rows)
		_ = writer.NewConsoleWriter().Write(res)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
	rootCmd.PersistentFlags().String("query", "", "Specify the query. Use --query=<query>")
}
