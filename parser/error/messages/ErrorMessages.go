package messages

const (
	ErrorMessageEmptyQuery                   = "expected query to be non-empty"
	ErrorMessageNonSelectQuery               = "expected a select query statement"
	ErrorMessageLimitValue                   = "expected a limit value"
	ErrorMessageLimitValueInt                = "expected limit to be an integer"
	ErrorMessageMissingBy                    = "expected 'by' after order"
	ErrorMessageMissingCommaOrderBy          = "expected a comma after 'order by' in column separator"
	ErrorMessageMissingOrderByColumns        = "expected a column after 'order by'"
	ErrorMessageMissingSource                = "expected a source path after 'from`"
	ErrorMessageMissingCommaProjection       = "expected a comma in projection list or from keyword"
	ErrorMessageOpeningParenthesesProjection = "expected an opening parentheses in projection"
	ErrorMessageClosingParenthesesProjection = "expected a closing parentheses in projection"
)
