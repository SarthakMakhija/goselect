package messages

const (
	ErrorMessageEmptyQuery                         = "expected query to be non-empty"
	ErrorMessageNonSelectQuery                     = "expected a select query statement"
	ErrorMessageLimitValue                         = "expected a limit value"
	ErrorMessageLimitValueInt                      = "expected limit to be an integer"
	ErrorMessageMissingBy                          = "expected 'by' after order"
	ErrorMessageMissingCommaOrderBy                = "expected a comma after 'order by' in attribute separator"
	ErrorMessageMissingOrderByAttributes           = "expected a attribute after 'order by'"
	ErrorMessageNonZeroPositivePositions           = "expected non-zero & positive order by positions"
	ErrorMessageMissingSource                      = "expected a source path after 'from`"
	ErrorMessageMissingCommaProjection             = "expected a comma in projection list or from keyword"
	ErrorMessageOpeningParenthesesProjection       = "expected an opening parentheses in projection"
	ErrorMessageInvalidProjection                  = "invalid projection"
	ErrorMessageMissingParameterInScalarFunctions  = "expected %v parameters in the function %v"
	ErrorMessageIncorrectValueType                 = "expected a %v value type but was not"
	ErrorMessageIncorrectEndIndexInSubstring       = "expected the end index to be greater than the from index in substr"
	ErrorMessageIllegalFromToIndexInSubstring      = "expected the from and to index to be positive integers"
	ErrorMessageFinalStateCalledOnScalarFunction   = "expected final state to be called on aggregate function"
	ErrorMessageInitialStateCalledOnScalarFunction = "expected initial state to be called on aggregate function"
)
