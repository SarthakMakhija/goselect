package projection

import (
	"errors"
	"fmt"
	"goselect/parser/context"
	"goselect/parser/error/messages"
	"goselect/parser/expression"
	"goselect/parser/tokenizer"
)

type Projections struct {
	expressions expression.Expressions
}

func NewProjections(
	tokenIterator *tokenizer.TokenIterator,
	context *context.ParsingApplicationContext,
) (*Projections, error) {

	expressions, err := all(tokenIterator, context)
	if err != nil {
		return nil, err
	}
	if expressions.Count() == 0 {
		return nil, errors.New(messages.ErrorMessageExpectedExpressionInProjection)
	}
	return &Projections{expressions: expressions}, nil
}

func (projections Projections) Count() int {
	return projections.expressions.Count()
}

func (projections Projections) AggregationCount() int {
	return projections.expressions.AggregationCount()
}

func (projections Projections) HasAllAggregates() bool {
	return projections.Count() == projections.AggregationCount()
}

func (projections Projections) DisplayableAttributes() []string {
	return projections.expressions.DisplayableAttributes()
}

func (projections Projections) EvaluateWith(
	fileAttributes *context.FileAttributes,
	functions *context.AllFunctions,
) ([]context.Value, []bool, []*expression.Expression, error) {
	return projections.expressions.EvaluateWith(fileAttributes, functions)
}

/*
projection:  attributes Or functions Or expressions
attributes:  name, size etc
functions: 	 min(size), lower(name), min(Count(size)) etc
expressions: add(..), mul(..), gt(..)
*/
func all(
	tokenIterator *tokenizer.TokenIterator,
	ctx *context.ParsingApplicationContext,
) (expression.Expressions, error) {

	var expressions []*expression.Expression
	var expectComma bool

	for tokenIterator.HasNext() && !tokenIterator.Peek().Equals("from") {
		token := tokenIterator.Next()
		switch {
		case expectComma:
			if !token.Equals(",") {
				return expression.Expressions{}, errors.New(messages.ErrorMessageMissingCommaProjection)
			}
			expectComma = false
		case context.IsAWildcardAttribute(token.TokenValue):
			expressions = append(expressions, expression.WithAttributes(context.AttributesOnWildcard())...)
			expectComma = true
		case ctx.IsASupportedAttribute(token.TokenValue):
			expressions = append(expressions, expression.WithAttribute(token.TokenValue))
			expectComma = true
		case ctx.IsASupportedFunction(token.TokenValue):
			function, err := function(token, tokenIterator, ctx)
			if err != nil {
				return expression.Expressions{}, err
			}
			expressions = append(expressions, expression.WithFunctionInstance(function))
			expectComma = true
		}
	}
	return expression.Expressions{Expressions: expressions}, nil
}

func function(
	functionNameToken tokenizer.Token,
	tokenIterator *tokenizer.TokenIterator,
	ctx *context.ParsingApplicationContext,
) (*expression.FunctionInstance, error) {

	var parseFunction func(functionNameToken tokenizer.Token) (*expression.FunctionInstance, error)

	aggregateFunctionStateOrNil := func(fn string) *context.FunctionState {
		if ctx.IsAnAggregateFunction(fn) {
			return ctx.InitialState(fn)
		}
		return nil
	}
	parseFunction = func(functionNameToken tokenizer.Token) (*expression.FunctionInstance, error) {
		var functionArgs []*expression.Expression
		expectOpeningParentheses := true

		for tokenIterator.HasNext() && !tokenIterator.Peek().Equals("from") {
			token := tokenIterator.Next()
			switch {
			case expectOpeningParentheses:
				if !token.Equals("(") {
					return nil, fmt.Errorf(messages.ErrorMessageOpeningParenthesesProjection, functionNameToken.TokenValue)
				}
				expectOpeningParentheses = false
			case token.Equals(")"):
				var isAggregate bool
				if ctx.IsAnAggregateFunction(functionNameToken.TokenValue) {
					isAggregate = true
				} else {
					isAggregate = false
				}
				return expression.FunctionInstanceWith(
					functionNameToken.TokenValue,
					functionArgs,
					aggregateFunctionStateOrNil(functionNameToken.TokenValue),
					isAggregate,
				), nil
			case ctx.IsASupportedFunction(token.TokenValue):
				fn, err := parseFunction(token)
				if err != nil {
					return nil, err
				}
				functionArgs = append(functionArgs, expression.WithFunctionInstance(fn))
			case ctx.IsASupportedAttribute(token.TokenValue):
				functionArgs = append(functionArgs, expression.WithAttribute(token.TokenValue))
				expectOpeningParentheses = false
			default:
				if !token.Equals(",") {
					value, err := context.ToValue(token)
					if err != nil {
						value = context.StringValue(token.TokenValue)
					}
					functionArgs = append(functionArgs, expression.WithValue(value))
				}
			}
		}
		return nil, nil
	}
	fn, err := parseFunction(functionNameToken)
	if err != nil {
		return nil, err
	}
	if fn == nil {
		return nil, errors.New(messages.ErrorMessageInvalidProjection)
	}
	return fn, nil
}
