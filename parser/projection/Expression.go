package projection

import (
	"goselect/parser/context"
)

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

func (expressions Expressions) evaluateWith(fileAttributes *context.FileAttributes, functions *context.AllFunctions) ([]context.Value, error) {
	var values []context.Value

	var execute func(expression *Expression) (context.Value, error)
	execute = func(expression *Expression) (context.Value, error) {
		if !expression.isAFunction() {
			return fileAttributes.Get(expression.attribute), nil
		}
		v, err := execute(expression.function.left)
		if err != nil {
			return context.EmptyValue(), err
		}
		return functions.Execute(expression.function.name, v)
	}
	for _, expression := range expressions.expressions {
		if !expression.isAFunction() {
			values = append(values, fileAttributes.Get(expression.attribute))
		} else {
			value, err := execute(expression)
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
	}
	return values, nil
}

func (expression Expression) isAFunction() bool {
	return expression.function != nil
}
