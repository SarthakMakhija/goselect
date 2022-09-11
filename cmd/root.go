package cmd

import (
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goselect",
	Short: "SQL like select interface for file system",
	Example: `
1. goselect execute -q='select filename, absolutepath from .'
2. goselect execute -q='select name, size, fmtsize from . where like(name, results.*) order by 2'
3. goselect execute -q='select name, size, fmtsize from . where or(like(name, results.*), gt(size, 2048)) order by 2 limit 5'
`,
	Long: `goselect provides SQL like 'select' interface for file systems. The syntax for select query is: select <attributes> from <directory> [where condition] [order by] [limit].
Queries are case-insensitive in nature. 

goselect provides various features including:
1. Support for attribute aliases. For example, filename is same as fname  
2. Support for function aliases. For example, lower is same as low 
3. Support for various scalar functions like lower, upper, now, concat etc
4. Support for various aggregate functions like count, countdistinct, average etc
5. Support for exporting the results in table, json and html format

Features that are different from SQL:
1. goselect does not support 'group by'. All the aggregating functions return results that repeat for each row
2. goselect does not support expressions like 1+2 or 1*2. Instead, goselect gives functions like 'add' and 'mul' etc to write such expressions
3. goselect does not support expressions like name=sample.log in 'where' clause. Instead, various functions are given to represent such expressions. These functions include: 'eq', 'ne', 'lt' etc
4. goselect has a weak grammar. For example, a query like: select 1+2, name from /home/projects will ignore 1+2 and return file names
5. goselect's 'order by' clause supports only attribute positions. For example, a query like: select name, size from /home/projects order by 1 will order the results by first attribute
6. goselect does not support single quote ['] and double quotes["]. For example, to match a file name, one could simply write a query: select * from . where eq(name, sample) 

goselect is available here: https://github.com/SarthakMakhija/goselect
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cc.Init(&cc.Config{
		RootCmd:       rootCmd,
		Headings:      cc.HiCyan + cc.Bold + cc.Underline,
		Commands:      cc.HiYellow + cc.Bold,
		Example:       cc.Italic,
		ExecName:      cc.HiRed,
		Flags:         cc.Bold,
		FlagsDescr:    cc.Cyan,
		CmdShortDescr: cc.Red,
	})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
