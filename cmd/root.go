package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goselect",
	Short: "SQL like select interface for file system",
	Long: `goselect provides SQL like 'select' interface for file system. The syntax for select query is: select <attributes> from <directory> [where condition] [order by] [limit].
Queries are case-insensitive in nature. goselect is available here: https://github.com/SarthakMakhija/goselect. goselect provides various features including:

1. Support for attribute aliases. Example, filename is same as fname  
2. Support for function aliases. Example, lower is same as low 
3. Support for various scalar functions like lower, upper, now, concat etc
4. Support for various aggregate functions like count, countdistinct, average etc
5. Support for exporting the results in table, json and html format

Features that are different from SQL:
1. goselect does not support 'group by'. All the aggregating functions return results that repeat for each row
2. goselect does not support expressions like '1+2', '1*2'. Instead, goselect gives functions like 'add', 'mul' etc to write such expressions
3. goselect does not support expressions like name='sample.log' in 'where' clause. Instead, various functions are given to represent such expressions. These functions include: 'eq', 'ne', 'lt' etc
4. goselect has a weak grammar. For example, a query like 'select 1+2, name from /home/projects' will ignore 1+2
5. goselect's 'order by' clause supports only attribute positions. For example, a query like 'select name, size from /home/projects order by 1'`,
}

const (
	headerColor  = "\033[34m"
	contentColor = "\033[0m"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goselect.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
