package context

type ParsingApplicationContext struct {
	allFunctions  *AllFunctions
	allAttributes *AllAttributes
}

func NewContext(functions *AllFunctions, attributes *AllAttributes) *ParsingApplicationContext {
	return &ParsingApplicationContext{allFunctions: functions, allAttributes: attributes}
}

func (context *ParsingApplicationContext) IsASupportedAttribute(attribute string) bool {
	return context.allAttributes.IsASupportedAttribute(attribute)
}

func (context *ParsingApplicationContext) IsASupportedFunction(functionName string) bool {
	return context.allFunctions.IsASupportedFunction(functionName)
}

func (context *ParsingApplicationContext) FunctionContainsATag(function string, tag string) bool {
	return context.allFunctions.ContainsATag(function, tag)
}

func (context *ParsingApplicationContext) IsAnAggregateFunction(functionName string) bool {
	return context.allFunctions.IsAnAggregateFunction(functionName)
}

func (context *ParsingApplicationContext) InitialState(functionName string) *FunctionState {
	return context.allFunctions.InitialState(functionName)
}

func (context *ParsingApplicationContext) AllFunctions() *AllFunctions {
	return context.allFunctions
}
