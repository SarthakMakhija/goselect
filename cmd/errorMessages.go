package cmd

const (
	ErrorMessageEmptyTerm                      = "expected term to be non-empty. please use --term=<attribute> or --term=<function>"
	ErrorMessageInvalidTerm                    = "expected term to be one of the supported attributes or a function"
	ErrorMessageInvalidExportFormat            = "expected export format to be one of the supported exported formats: %v"
	ErrorMessageAttemptedToExportTableToFile   = "table can not be exported to a file"
	ErrorMessageExpectedFilePathToBeADirectory = "expected file path to be a directory"
	ErrorMessageExpectedAQueryForAnAlias       = "expected a query to exist for the alias %v, but none was found"
)
