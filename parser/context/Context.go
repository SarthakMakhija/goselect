package context

type Context struct {
	allFunctions  *AllFunctions
	allAttributes *AllAttributes
}

func NewContext(functions *AllFunctions, attributes *AllAttributes) *Context {
	return &Context{allFunctions: functions, allAttributes: attributes}
}

func (context *Context) IsASupportedAttribute(attribute string) bool {
	return context.allAttributes.IsASupportedAttribute(attribute)
}

func (context *Context) IsASupportedFunction(functionName string) bool {
	return context.allFunctions.IsASupportedFunction(functionName)
}

func (context *Context) AllFunctions() *AllFunctions {
	return context.allFunctions
}
