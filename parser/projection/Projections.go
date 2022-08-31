package projection

import (
	"errors"
	"goselect/parser/context"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
)

type Projections struct {
	expressions Expressions
}

func NewProjections(tokenIterator *tokenizer.TokenIterator, context *context.ParsingApplicationContext) (*Projections, error) {
	if expressions, err := all(tokenIterator, context); err != nil {
		return nil, err
	} else {
		if expressions.count() == 0 {
			return nil, errors.New("expected at least one attribute in projection list")
		}
		return &Projections{expressions: expressions}, nil
	}
}

func (projections Projections) Count() int {
	return projections.expressions.count()
}

func (projections Projections) DisplayableAttributes() []string {
	return projections.expressions.displayableAttributes()
}

func (projections Projections) EvaluateWith(fileAttributes *context.FileAttributes, functions *context.AllFunctions) ([]context.Value, []bool, []*Expression, error) {
	return projections.expressions.evaluateWith(fileAttributes, functions)
}

/*
projection:  attributes Or functions Or expressions
attributes:  name, size etc
functions: 	 min(size), lower(name), min(Count(size)) etc
expressions: 2 + 3, 2 > 3 etc
*/
func all(tokenIterator *tokenizer.TokenIterator, ctx *context.ParsingApplicationContext) (Expressions, error) {
	var expressions []*Expression
	var expectComma bool

	for tokenIterator.HasNext() && !tokenIterator.Peek().Equals("from") {
		token := tokenIterator.Next()
		switch {
		case expectComma:
			if !token.Equals(",") {
				return Expressions{}, errors.New(messages.ErrorMessageMissingCommaProjection)
			}
			expectComma = false
		case context.IsAWildcardAttribute(token.TokenValue):
			expressions = append(expressions, expressionsWithAttributes(context.AttributesOnWildcard())...)
			expectComma = true
		case ctx.IsASupportedAttribute(token.TokenValue):
			expressions = append(expressions, expressionWithAttribute(token.TokenValue))
			expectComma = true
		case ctx.IsASupportedFunction(token.TokenValue):
			if function, err := function(token, tokenIterator, ctx); err != nil {
				return Expressions{}, err
			} else {
				expressions = append(expressions, expressionWithFunctionInstance(function))
			}
			expectComma = true
		}
	}
	return Expressions{expressions: expressions}, nil
}

func function(functionNameToken tokenizer.Token, tokenIterator *tokenizer.TokenIterator, ctx *context.ParsingApplicationContext) (*FunctionInstance, error) {
	var parseFunction func(functionNameToken tokenizer.Token) (*FunctionInstance, error)

	aggregateFunctionStateOrNil := func(fn string) *context.AggregateFunctionState {
		if ctx.IsAnAggregateFunction(fn) {
			return ctx.InitialState(fn)
		}
		return nil
	}
	parseFunction = func(functionNameToken tokenizer.Token) (*FunctionInstance, error) {
		var functionArgs []*Expression
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
				return &FunctionInstance{
					name:  functionNameToken.TokenValue,
					args:  functionArgs,
					state: aggregateFunctionStateOrNil(functionNameToken.TokenValue),
				}, nil
			case ctx.IsASupportedFunction(token.TokenValue):
				fn, err := parseFunction(token)
				if err != nil {
					return nil, err
				}
				functionArgs = append(functionArgs, expressionWithFunctionInstance(fn))
			case ctx.IsASupportedAttribute(token.TokenValue):
				functionArgs = append(functionArgs, expressionWithAttribute(token.TokenValue))
				expectOpeningParentheses = false
			default:
				if !token.Equals(",") {
					functionArgs = append(functionArgs, expressionWithValue(token.TokenValue))
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
