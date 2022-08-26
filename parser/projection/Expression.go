package projection

type Expressions struct {
	expressions []*Expression
}

type Expression struct {
	column   string
	function *Function
}

type Function struct {
	name string
	left *Expression
}

func (expressions Expressions) allColumns() []string {
	var columns []string
	for _, expression := range expressions.expressions {
		columns = append(columns, expression.column)
	}
	return columns
}

func (expressions Expressions) functions() []string {
	var functionAsString func(expression *Expression) string
	functionAsString = func(expression *Expression) string {
		if expression.function == nil {
			return expression.column
		}
		return expression.function.name + "(" + functionAsString(expression.function.left) + ")"
	}

	var functions []string
	for _, expression := range expressions.expressions {
		functions = append(functions, functionAsString(expression))
	}
	return functions
}

func expressionWithColumn(column string) *Expression {
	return &Expression{
		column: column,
	}
}

func expressionsWithColumns(columns []string) []*Expression {
	var expressions = make([]*Expression, len(columns))
	for index, column := range columns {
		expressions[index] = expressionWithColumn(column)
	}
	return expressions
}

func expressionWithFunction(fn *Function) *Expression {
	return &Expression{
		function: fn,
	}
}
