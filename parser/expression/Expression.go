package expression

import (
	"goselect/parser/context"
)

type Expressions struct {
	Expressions []*Expression
}

type expressionType int

const (
	TypeValue     = 0
	TypeAttribute = 1
	TypeFunction  = 2
)

type Expression struct {
	eType     expressionType
	value     context.Value
	attribute string
	function  *FunctionInstance
}

type FunctionInstance struct {
	name        string
	args        []*Expression
	state       *context.FunctionState
	isAggregate bool
}

func FunctionInstanceWith(name string, args []*Expression, state *context.FunctionState, isAggregate bool) *FunctionInstance {
	return &FunctionInstance{
		name:        name,
		args:        args,
		state:       state,
		isAggregate: isAggregate,
	}
}

func WithAttribute(attribute string) *Expression {
	return &Expression{
		eType:     TypeAttribute,
		attribute: attribute,
	}
}

func WithFunctionInstance(fn *FunctionInstance) *Expression {
	return &Expression{
		eType:    TypeFunction,
		function: fn,
	}
}

func WithValue(value context.Value) *Expression {
	return &Expression{
		eType: TypeValue,
		value: value,
	}
}

func WithAttributes(attributes []string) []*Expression {
	var expressions = make([]*Expression, len(attributes))
	for index, attribute := range attributes {
		expressions[index] = WithAttribute(attribute)
	}
	return expressions
}

func (expressions Expressions) Count() int {
	return len(expressions.Expressions)
}

func (expressions Expressions) ExpressionAt(index int) *Expression {
	if index < len(expressions.Expressions) {
		return expressions.Expressions[index]
	}
	return nil
}

func (expressions Expressions) DisplayableAttributes() []string {
	var functionAsString func(expression *Expression) string
	functionAsString = func(expression *Expression) string {
		if !expression.isAFunction() {
			if expression.eType == TypeAttribute {
				return expression.attribute
			}
			return expression.value.GetAsString()
		}
		var result = expression.function.name + "("
		for _, arg := range expression.function.args {
			result = result + functionAsString(arg) + ","
		}
		if len(expression.function.args) > 0 {
			result = result[0:len(result)-1] + ")"
		} else {
			result = result + ")"
		}
		return result
	}

	var attributes []string
	for _, expression := range expressions.Expressions {
		if expression.eType == TypeAttribute {
			attributes = append(attributes, expression.attribute)
		} else if expression.eType == TypeValue {
			attributes = append(attributes, expression.value.GetAsString())
		} else {
			attributes = append(attributes, functionAsString(expression))
		}
	}
	return attributes
}

func (expressions Expressions) AggregationCount() int {
	count := 0
	for _, expression := range expressions.Expressions {
		if expression.HasAnAggregate() {
			count = count + 1
		}
	}
	return count
}

func (expressions Expressions) EvaluateWith(
	fileAttributes *context.FileAttributes,
	functions *context.AllFunctions,
) ([]context.Value, []bool, []*Expression, error) {

	var values []context.Value
	var fullyEvaluated []bool
	var resultingExpressions []*Expression

	for _, expression := range expressions.Expressions {
		resultingExpressions = append(resultingExpressions, expression)
		if !expression.isAFunction() {
			values = append(values, expression.getNonFunctionValue(fileAttributes))
			fullyEvaluated = append(fullyEvaluated, true)
		} else {
			value, err, isAggregate := expression.Evaluate(fileAttributes, functions)
			if err != nil {
				return nil, nil, nil, err
			}
			values = append(values, value)
			if isAggregate {
				fullyEvaluated = append(fullyEvaluated, false)
			} else {
				fullyEvaluated = append(fullyEvaluated, true)
			}
		}
	}
	return values, fullyEvaluated, resultingExpressions, nil
}

func (expression Expression) Evaluate(
	fileAttributes *context.FileAttributes,
	functions *context.AllFunctions,
) (context.Value, error, bool) {

	if !expression.isAFunction() {
		return expression.getNonFunctionValue(fileAttributes), nil, false
	}

	var values []context.Value
	isAtleastOneExpressionAnAggregateFunction := false
	for _, arg := range expression.function.args {
		v, err, isAggregate := arg.Evaluate(fileAttributes, functions)
		if err != nil {
			return context.EmptyValue, err, isAggregate
		}
		if isAggregate {
			isAtleastOneExpressionAnAggregateFunction = true
		}
		values = append(values, v)
	}

	isAnAggregateFunction := functions.IsAnAggregateFunction(expression.function.name)
	if isAnAggregateFunction && !isAtleastOneExpressionAnAggregateFunction {
		state, err := functions.ExecuteAggregate(expression.function.name, expression.function.state, values...)
		if err != nil {
			return context.EmptyValue, err, true
		}
		expression.function.state = state
		return state.Initial, err, true
	}
	if !isAnAggregateFunction && !isAtleastOneExpressionAnAggregateFunction {
		v, err := functions.Execute(expression.function.name, values...)
		return v, err, isAtleastOneExpressionAnAggregateFunction
	}
	return context.EmptyValue, nil, isAtleastOneExpressionAnAggregateFunction
}

func (expression *Expression) FullyEvaluate(functions *context.AllFunctions) (context.Value, error) {
	var execute func(expression *Expression) (context.Value, error)
	execute = func(expression *Expression) (context.Value, error) {
		if !expression.isAFunction() {
			return context.EmptyValue, nil
		}
		var values []context.Value
		for _, arg := range expression.function.args {
			if arg.isAFunction() && functions.IsAnAggregateFunction(arg.function.name) {
				v, err := execute(arg)
				if err != nil {
					return context.EmptyValue, err
				}
				values = append(values, v)
			} else {
				if arg.eType == TypeValue {
					values = append(values, arg.value)
				}
			}
		}
		if functions.IsAnAggregateFunction(expression.function.name) {
			return functions.FinalValue(expression.function.name, expression.function.state, values)
		}
		return functions.Execute(expression.function.name, values...)
	}
	return execute(expression)
}

func (expression Expression) HasAnAggregate() bool {
	isAnyArgumentAnAggregate := func(fn *FunctionInstance) bool {
		if fn != nil {
			for _, arg := range fn.args {
				if arg.isAFunction() && arg.function.isAggregate {
					return true
				}
			}
			return false
		}
		return false
	}
	return (expression.isAFunction() && expression.function.isAggregate) || isAnyArgumentAnAggregate(expression.function)
}

func (expression Expression) isAFunction() bool {
	return expression.function != nil
}

func (expression Expression) getNonFunctionValue(fileAttributes *context.FileAttributes) context.Value {
	if expression.eType == TypeAttribute {
		return fileAttributes.Get(expression.attribute)
	}
	return expression.value
}
