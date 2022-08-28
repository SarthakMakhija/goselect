package projection

import "goselect/parser/context"

type Expressions struct {
	expressions []*Expression
}

type Expression struct {
	attribute string
	function  *Function
}

type Function struct {
	name string
	left *Expression
}

func expressionWithAttribute(attribute string) *Expression {
	return &Expression{
		attribute: attribute,
	}
}

func expressionsWithAttributes(attributes []string) []*Expression {
	var expressions = make([]*Expression, len(attributes))
	for index, attribute := range attributes {
		expressions[index] = expressionWithAttribute(attribute)
	}
	return expressions
}

func expressionWithFunction(fn *Function) *Expression {
	return &Expression{
		function: fn,
	}
}

func (expressions Expressions) count() int {
	return len(expressions.expressions)
}

func (expressions Expressions) displayableAttributes() []string {
	var functionAsString func(expression *Expression) string
	functionAsString = func(expression *Expression) string {
		if !expression.isAFunction() {
			return expression.attribute
		}
		return expression.function.name + "(" + functionAsString(expression.function.left) + ")"
	}

	var attributes []string
	for _, expression := range expressions.expressions {
		if len(expression.attribute) != 0 {
			attributes = append(attributes, expression.attribute)
		} else {
			attributes = append(attributes, functionAsString(expression))
		}
	}
	return attributes
}

func (expressions Expressions) evaluateWith(fileAttributes *context.FileAttributes, functions *context.AllFunctions) []interface{} {
	var values []interface{}

	var execute func(expression *Expression) interface{}
	execute = func(expression *Expression) interface{} {
		if !expression.isAFunction() {
			return fileAttributes.Get(expression.attribute)
		}
		v := execute(expression.function.left)
		return functions.Execute(expression.function.name, v)
	}
	for _, expression := range expressions.expressions {
		if !expression.isAFunction() {
			values = append(values, fileAttributes.Get(expression.attribute))
		} else {
			values = append(values, execute(expression))
		}
	}
	return values
}

func (expression Expression) isAFunction() bool {
	return expression.function != nil
}
