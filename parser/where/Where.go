package where

import (
	"errors"
	"goselect/parser/context"
	"goselect/parser/error/messages"
	"goselect/parser/expression"
	"goselect/parser/tokenizer"
)

type Where struct {
	expressions expression.Expressions
}

func NewWhere(
	tokenIterator *tokenizer.TokenIterator,
	context *context.ParsingApplicationContext,
) (*Where, error) {
	if expressions, isWhereSpecified, err := all(tokenIterator, context); err != nil {
		return nil, err
	} else {
		if isWhereSpecified && expressions.Count() == 0 {
			return nil, errors.New(messages.ErrorMessageExpectedExpressionInWhere)
		}
		if isWhereSpecified && expressions.Count() > 1 {
			return nil, errors.New(messages.ErrorMessageExpectedSingleExpressionInWhere)
		}
		return &Where{expressions: expressions}, nil
	}
}

func (where Where) Display() string {
	if attributes := where.expressions.DisplayableAttributes(); len(attributes) >= 1 {
		return attributes[0]
	}
	return ""
}

func (where Where) EvaluateWith(
	fileAttributes *context.FileAttributes,
	functions *context.AllFunctions,
) (bool, error) {

	if where.expressions.Count() == 1 {
		if expr := where.expressions.ExpressionAt(0); expr != nil {
			value, err, _ := expr.Evaluate(fileAttributes, functions)
			if err != nil {
				return false, err
			}
			return value.GetBoolean()
		}
	}
	return true, nil
}

func all(
	tokenIterator *tokenizer.TokenIterator,
	ctx *context.ParsingApplicationContext,
) (expression.Expressions, bool, error) {

	var expressions []*expression.Expression
	if !tokenIterator.HasNext() {
		return expression.Expressions{}, false, nil
	}
	if tokenIterator.HasNext() && !tokenIterator.Peek().Equals("where") {
		return expression.Expressions{}, false, nil
	}
	if tokenIterator.HasNext() && tokenIterator.Peek().Equals("where") {
		tokenIterator.Next()
	}
	if tokenIterator.HasNext() && !tokenIterator.Peek().Equals("order") {
		token := tokenIterator.Next()
		switch {
		case ctx.IsASupportedFunction(token.TokenValue) && ctx.FunctionContainsATag(token.TokenValue, "where"):
			if function, err := function(token, tokenIterator, ctx); err != nil {
				return expression.Expressions{}, true, err
			} else {
				expressions = append(expressions, expression.ExpressionWithFunctionInstance(function))
			}
		default:
			return expression.Expressions{}, true, errors.New(messages.ErrorMessageInvalidWhereFunctionUsed)
		}
	}
	return expression.Expressions{Expressions: expressions}, true, nil
}

func function(
	functionNameToken tokenizer.Token,
	tokenIterator *tokenizer.TokenIterator,
	ctx *context.ParsingApplicationContext,
) (*expression.FunctionInstance, error) {

	var parseFunction func(functionNameToken tokenizer.Token) (*expression.FunctionInstance, error)
	parseFunction = func(functionNameToken tokenizer.Token) (*expression.FunctionInstance, error) {
		var functionArgs []*expression.Expression
		expectOpeningParentheses := true

		for tokenIterator.HasNext() && !tokenIterator.Peek().Equals("order") {
			token := tokenIterator.Next()
			switch {
			case expectOpeningParentheses:
				if !token.Equals("(") {
					return nil, errors.New(messages.ErrorMessageOpeningParenthesesProjection)
				}
				expectOpeningParentheses = false
			case token.Equals(")"):
				return expression.FunctionInstanceWith(
					functionNameToken.TokenValue,
					functionArgs,
					nil,
				), nil
			case ctx.IsASupportedFunction(token.TokenValue):
				if ctx.IsAnAggregateFunction(token.TokenValue) {
					return nil, errors.New(messages.ErrorMessageAggregateFunctionInsideWhere)
				}
				fn, err := parseFunction(token)
				if err != nil {
					return nil, err
				}
				functionArgs = append(functionArgs, expression.ExpressionWithFunctionInstance(fn))
			case ctx.IsASupportedAttribute(token.TokenValue):
				functionArgs = append(functionArgs, expression.ExpressionWithAttribute(token.TokenValue))
				expectOpeningParentheses = false
			default:
				if !token.Equals(",") {
					functionArgs = append(functionArgs, expression.ExpressionWithValue(token.TokenValue))
				}
			}
		}
		return nil, nil
	}

	if fn, err := parseFunction(functionNameToken); err != nil {
		return nil, err
	} else if fn == nil {
		return nil, errors.New(messages.ErrorMessageInvalidWhere)
	} else {
		return fn, nil
	}
}
