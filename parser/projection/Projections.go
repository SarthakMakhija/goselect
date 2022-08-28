package projection

import (
	"errors"
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"goselect/parser/context"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
)

type Projections struct {
	expressions Expressions
}

func NewProjections(tokenIterator *tokenizer.TokenIterator, context *context.Context) (*Projections, error) {
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

func (projections Projections) AllExpressions() Expressions {
	return projections.expressions
}

func (projections Projections) DisplayableAttributes() []string {
	return projections.expressions.displayableAttributes()
}

func (projections Projections) EvaluateWith(fileAttributes *context.FileAttributes, functions *context.AllFunctions) []interface{} {
	return projections.expressions.evaluateWith(fileAttributes, functions)
}

/*
projection:  attributes Or functions Or expressions
attributes:  name, size etc
functions: 	 min(size), lower(name), min(Count(size)) etc
expressions: 2 + 3, 2 > 3 etc
*/
func all(tokenIterator *tokenizer.TokenIterator, ctx *context.Context) (Expressions, error) {
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
			tokenIterator.Drop()
			if function, err := function(tokenIterator, ctx); err != nil {
				return Expressions{}, err
			} else {
				expressions = append(expressions, expressionWithFunction(function))
			}
			expectComma = true
		}
	}
	return Expressions{expressions: expressions}, nil
}

func function(tokenIterator *tokenizer.TokenIterator, ctx *context.Context) (*Function, error) {
	buildFunction := func(functionStack *linkedliststack.Stack, operatingAttribute tokenizer.Token) *Function {
		functionToken, _ := functionStack.Pop()
		var rootFunction = &Function{
			name: (functionToken.(tokenizer.Token)).TokenValue,
			left: &Expression{attribute: operatingAttribute.TokenValue},
		}
		for !functionStack.Empty() {
			functionToken, _ := functionStack.Pop()
			rootFunction = &Function{
				name: (functionToken.(tokenizer.Token)).TokenValue,
				left: &Expression{function: rootFunction},
			}
		}
		return rootFunction
	}

	parseToken := func() (*linkedliststack.Stack, tokenizer.Token, error) {
		var operatingAttribute tokenizer.Token
		expectOpeningParentheses, expectClosingParentheses := false, false
		closingParenthesesCount, functionStack := 0, linkedliststack.New()

	loop:
		for tokenIterator.HasNext() && !tokenIterator.Peek().Equals("from") {
			token := tokenIterator.Next()
			switch {
			case expectOpeningParentheses:
				if !token.Equals("(") {
					functionStack.Clear()
					return nil, tokenizer.Token{}, errors.New(messages.ErrorMessageOpeningParenthesesProjection)
				}
				expectOpeningParentheses = false
			case expectClosingParentheses || token.Equals(")"):
				if !token.Equals(")") {
					functionStack.Clear()
					return nil, tokenizer.Token{}, errors.New(messages.ErrorMessageClosingParenthesesProjection)
				}
				closingParenthesesCount = closingParenthesesCount + 1
				if closingParenthesesCount == functionStack.Size() {
					expectClosingParentheses = false
					break loop
				}
			case ctx.IsASupportedFunction(token.TokenValue):
				functionStack.Push(token)
				expectOpeningParentheses = true
			case ctx.IsASupportedAttribute(token.TokenValue):
				operatingAttribute = token
				expectOpeningParentheses, expectClosingParentheses = false, true
			}
		}
		return functionStack, operatingAttribute, nil
	}

	if stack, operatingAttribute, err := parseToken(); err != nil {
		return nil, err
	} else {
		return buildFunction(stack, operatingAttribute), nil
	}
}
