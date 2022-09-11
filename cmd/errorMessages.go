package cmd

const (
	ErrorMessageEmptyTerm                      = "expected term to be non-empty. please use --term=<attribute> or --term=<function>"
	ErrorMessageInvalidTerm                    = "expected term to be one of the supported attributes or a function"
	ErrorMessageEmptyQuery                     = "expected query to be non-empty. please use --query to specify the query"
	ErrorMessageInvalidExportFormat            = "expected export format to be one of the supported exported formats: %v"
	ErrorMessageAttemptedToExportTableToFile   = "table can not be exported to a file"
	ErrorMessageExpectedFilePathToBeADirectory = "expected file path to be a directory"
)
