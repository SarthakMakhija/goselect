package projection

import (
	"errors"
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
)

type Projections struct {
	expressions Expressions
}

func NewProjections(tokenIterator *tokenizer.TokenIterator) (*Projections, error) {
	if expressions, err := all(tokenIterator); err != nil {
		return nil, err
	} else {
		if expressions.count() == 0 {
			return nil, errors.New("expected at least one column in projection list")
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

/*
projection: columns Or functions Or expressions
columns: 	 name, size etc
functions: 	 min(size), lower(name), min(Count(size)) etc
expressions: 2 + 3, 2 > 3 etc
*/
func all(tokenIterator *tokenizer.TokenIterator) (Expressions, error) {
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
		case isAWildcard(token.TokenValue):
			expressions = append(expressions, expressionsWithColumns(columnsOnWildcard())...)
			expectComma = true
		case IsASupportedColumn(token.TokenValue):
			expressions = append(expressions, expressionWithColumn(token.TokenValue))
			expectComma = true
		case isASupportedFunction(token.TokenValue):
			tokenIterator.Drop()
			if function, err := function(tokenIterator); err != nil {
				return Expressions{}, err
			} else {
				expressions = append(expressions, expressionWithFunction(function))
			}
			expectComma = true
		}
	}
	return Expressions{expressions: expressions}, nil
}

func function(tokenIterator *tokenizer.TokenIterator) (*Function, error) {
	buildFunction := func(functionStack *linkedliststack.Stack, operatingColumn tokenizer.Token) *Function {
		functionToken, _ := functionStack.Pop()
		var rootFunction = &Function{
			name: (functionToken.(tokenizer.Token)).TokenValue,
			left: &Expression{column: operatingColumn.TokenValue},
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
		var operatingColumn tokenizer.Token
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
			case isASupportedFunction(token.TokenValue):
				functionStack.Push(token)
				expectOpeningParentheses = true
			case IsASupportedColumn(token.TokenValue):
				operatingColumn = token
				expectOpeningParentheses, expectClosingParentheses = false, true
			}
		}
		return functionStack, operatingColumn, nil
	}

	if stack, operatingColumn, err := parseToken(); err != nil {
		return nil, err
	} else {
		return buildFunction(stack, operatingColumn), nil
	}
}
