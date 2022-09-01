package projection

import (
	"errors"
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
	if expressions, err := all(tokenIterator, context); err != nil {
		return nil, err
	} else {
		if expressions.Count() == 0 {
			return nil, errors.New("expected at least one attribute in projection list")
		}
		return &Projections{expressions: expressions}, nil
	}
}

func (projections Projections) Count() int {
	return projections.expressions.Count()
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
			expressions = append(expressions, expression.ExpressionsWithAttributes(context.AttributesOnWildcard())...)
			expectComma = true
		case ctx.IsASupportedAttribute(token.TokenValue):
			expressions = append(expressions, expression.ExpressionWithAttribute(token.TokenValue))
			expectComma = true
		case ctx.IsASupportedFunction(token.TokenValue):
			if function, err := function(token, tokenIterator, ctx); err != nil {
				return expression.Expressions{}, err
			} else {
				expressions = append(expressions, expression.ExpressionWithFunctionInstance(function))
			}
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
					return nil, errors.New(messages.ErrorMessageOpeningParenthesesProjection)
				}
				expectOpeningParentheses = false
			case token.Equals(")"):
				return expression.FunctionInstanceWith(
					functionNameToken.TokenValue,
					functionArgs,
					aggregateFunctionStateOrNil(functionNameToken.TokenValue),
				), nil
			case ctx.IsASupportedFunction(token.TokenValue):
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
		return nil, errors.New(messages.ErrorMessageInvalidProjection)
	} else {
		return fn, nil
	}
}
