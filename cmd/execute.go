package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"goselect/parser"
	"goselect/parser/context"
	"goselect/parser/executor"
	"goselect/parser/source"
	"goselect/parser/writer"
	"os"
	"strings"
)

var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute a select query",
	Long:  `Execute a select query. Select query syntax: select <columns> from <source directory> [where <condition>] [order by] [limit]`,
	Run: func(cmd *cobra.Command, args []string) {
		errorColor := "\033[31m"

		buildOptions := func() *executor.Options {
			nestedTraversal, _ := cmd.Flags().GetBool("nestedTraversal")
			ignoreTraversal, _ := cmd.Flags().GetStringSlice("skipDirectoryTraversal")

			options := executor.NewDefaultOptions()
			if nestedTraversal {
				options.EnableNestedTraversal()
			} else {
				options.DisableNestedTraversal()
			}
			options.DirectoriesToIgnoreTraversal(ignoreTraversal)
			return options
		}
		executeQuery := func(cmd *cobra.Command) (*executor.EvaluatingRows, *parser.SelectQuery, error) {
			rawQuery, _ := cmd.Flags().GetString("query")
			if len(rawQuery) == 0 {
				return nil, nil, errors.New(ErrorMessageEmptyQuery)
			}
			newContext := context.NewContext(context.NewFunctions(), context.NewAttributes())
			parser, err := parser.NewParser(rawQuery, newContext)
			if err != nil {
				return nil, nil, err
			}
			query, err := parser.Parse()
			if err != nil {
				return nil, nil, err
			}
			rows, err := executor.NewSelectQueryExecutor(query, newContext, buildOptions()).Execute()
			if err != nil {
				return nil, nil, err
			}
			return rows, query, nil
		}
		formatter := func(cmd *cobra.Command) (writer.Formatter, string, error) {
			exportFormat, _ := cmd.Flags().GetString("format")
			switch strings.ToLower(exportFormat) {
			case "json":
				return writer.NewJsonFormatter(), strings.ToLower(exportFormat), nil
			case "html":
				return writer.NewHtmlFormatter(), strings.ToLower(exportFormat), nil
			case "table":
				return writer.NewTableFormatter(), strings.ToLower(exportFormat), nil
			default:
				return nil, "", fmt.Errorf(ErrorMessageInvalidExportFormat, "json, html or table")
			}
		}
		writer := func(cmd *cobra.Command, format string) (writer.Writer, error) {
			directoryPath, _ := cmd.Flags().GetString("path")
			if len(directoryPath) == 0 {
				return writer.NewConsoleWriter(), nil
			}
			if strings.EqualFold(format, "table") {
				return nil, errors.New(ErrorMessageAttemptedToExportTableToFile)
			}
			directoryPath, err := source.ExpandDirectoryPath(directoryPath)
			if err != nil {
				return nil, err
			}
			if filePath, err := os.Stat(directoryPath); err != nil {
				return nil, err
			} else {
				if !filePath.IsDir() {
					return nil, errors.New(ErrorMessageExpectedFilePathToBeADirectory)
				}
				pathSeparator := string(os.PathSeparator)
				filePath := directoryPath + pathSeparator + fmt.Sprintf("results.%v", format)
				if strings.HasSuffix(directoryPath, pathSeparator) {
					filePath = directoryPath + fmt.Sprintf("results.%v", format)
				}
				writer, err := writer.NewFileWriter(filePath)
				if err != nil {
					return nil, err
				}
				return writer, nil
			}
		}
		run := func() {
			rows, query, err := executeQuery(cmd)
			if err != nil {
				fmt.Println(errorColor, err)
				return
			}
			exportFormatter, format, err := formatter(cmd)
			if err != nil {
				fmt.Println(errorColor, err)
				return
			}
			writer, err := writer(cmd, format)
			if err != nil {
				fmt.Println(errorColor, err)
				return
			}
			res := exportFormatter.Format(query.Projections, rows)
			if err := writer.Write(res); err != nil {
				fmt.Println(errorColor, err)
				return
			}
		}
		run()
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
		"specify if nested directories should be traversed. Use --nestedTraversal=<true/false> or -n=<true/false>",
	)
	rootCmd.PersistentFlags().StringSliceP(
		"skipDirectoryTraversal",
		"s",
		[]string{".git", ".github"},
		"specify the directory names that should not be traversed. Use --skipDirectoryTraversal=<directory> or -s=<directory>. Multiple directory names can be passed by using --skipDirectoryTraversal=.git --skipDirectoryTraversal=.github",
	)
	rootCmd.PersistentFlags().StringP(
		"format",
		"f",
		"table",
		"specify the export format. Supported values include: json, html and table. Use --format=<format>",
	)
	rootCmd.PersistentFlags().StringP(
		"path",
		"p",
		"",
		"specify the directory path. Use --path=<directoryPath>",
	)
}
