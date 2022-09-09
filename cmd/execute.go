package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"goselect/parser/writer"
	"strings"
)

var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute Select SQL query",
	Long:  `Select SQL Query syntax: select <columns> from <source directory> [where <condition>] [order by] [limit]`,
	Run: func(cmd *cobra.Command, args []string) {
		buildOptions := func() *executor.Options {
			nestedTraversal, _ := cmd.Flags().GetBool("nestedTraversal")
			ignoreTraversal, _ := cmd.Flags().GetStringSlice("ignoreTraversal")

			options := executor.NewDefaultOptions()
			if nestedTraversal {
				options.EnableNestedTraversal()
			} else {
				options.DisableNestedTraversal()
			}
			options.DirectoriesToIgnoreTraversal(ignoreTraversal)
			return options
		}
		formatter := func(cmd *cobra.Command) (writer.Formatter, error) {
			exportFormat, _ := cmd.Flags().GetString("format")
			switch strings.ToLower(exportFormat) {
			case "json":
				return writer.NewJsonFormatter(), nil
			case "html":
				return writer.NewHtmlFormatter(), nil
			case "table":
				return writer.NewTableFormatter(), nil
			default:
				return nil, errors.New("unsupported export format")
			}
		}

		rawQuery, _ := cmd.Flags().GetString("query")
		errorColor := "\033[31m"
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
		rows, err := executor.NewSelectQueryExecutor(query, newContext, buildOptions()).Execute()
		if err != nil {
			fmt.Println(errorColor, err)
			return
		}
		exportFormatter, err := formatter(cmd)
		if err != nil {
			fmt.Println(errorColor, err)
			return
		}
		res := exportFormatter.Format(query.Projections, rows)
		_ = writer.NewConsoleWriter().Write(res)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
	rootCmd.PersistentFlags().StringP(
		"query",
		"q",
		"",
		"specify the query. Use --query=<query> or -q=<query>",
	)
	rootCmd.PersistentFlags().BoolP(
		"nestedTraversal",
		"n",
		true,
		"specify the if nested directories should be traversed. Use --nestedTraversal=<true/false> or -n=<true/false>",
	)
	rootCmd.PersistentFlags().StringSliceP(
		"ignoreTraversal",
		"i",
		[]string{".git", ".github"},
		"specify the directory names that should not be traversed. Use --ignoreTraversal=<directory> or -i=<directory>. Multiple directory names can be passed by using --ignoreTraversal=.git --ignoreTraversal=.github",
	)
	rootCmd.PersistentFlags().StringP(
		"format",
		"f",
		"table",
		"specify the export format. Supported values include: json, html and table. Use --format=<format>",
	)
}
