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

func (context *ParsingApplicationContext) AllFunctions() *AllFunctions {
	return context.allFunctions
}