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

func (expressions Expressions) DisplayableColumns() []string {
	var functionAsString func(expression *Expression) string
	functionAsString = func(expression *Expression) string {
		if !expression.isAFunction() {
			return expression.column
		}
		return expression.function.name + "(" + functionAsString(expression.function.left) + ")"
	}

	var columns []string
	for _, expression := range expressions.expressions {
		if len(expression.column) != 0 {
			columns = append(columns, expression.column)
		} else {
			columns = append(columns, functionAsString(expression))
		}
	}
	return columns
}

func (expressions Expressions) count() int {
	return len(expressions.expressions)
}

func (expressions Expressions) ExecuteWith(fileAttributes *FileAttributes) []interface{} {
	var values []interface{}

	var execute func(expression *Expression) interface{}
	execute = func(expression *Expression) interface{} {
		if !expression.isAFunction() {
			return fileAttributes.Get(expression.column)
		}
		v := execute(expression.function.left)
		return ExecuteFn(expression.function.name, v)
	}
	for _, expression := range expressions.expressions {
		if !expression.isAFunction() {
			values = append(values, fileAttributes.Get(expression.column))
		} else {
			values = append(values, execute(expression))
		}
	}
	return values
}

func (expression Expression) isAFunction() bool {
	return expression.function != nil
}
