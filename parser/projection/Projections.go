package projection

import (
	"errors"
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"goselect/parser/tokenizer"
)

type Projections struct {
	tokenIterator *tokenizer.TokenIterator
}

func newProjections(tokenIterator *tokenizer.TokenIterator) *Projections {
	return &Projections{tokenIterator: tokenIterator}
}

/*
projection: columns Or functions Or expressions
columns: 	 name, size etc
functions: 	 min(size), lower(name), min(count(size)) etc
expressions: 2 + 3, 2 > 3 etc
*/
func (projections *Projections) all() (Expressions, error) {
	var expressions []*Expression
	var expectComma bool

	for projections.tokenIterator.HasNext() && !projections.tokenIterator.Peek().Equals("from") {
		token := projections.tokenIterator.Next()
		switch {
		case expectComma:
			if !token.Equals(",") {
				return Expressions{}, errors.New("expected a comma in projection list")
			}
			expectComma = false
		case isAWildcard(token.TokenValue):
			expressions = append(expressions, expressionsWithColumns(columnsOnWildcard())...)
			expectComma = true
		case isASupportedColumn(token.TokenValue):
			expressions = append(expressions, expressionWithColumn(token.TokenValue))
			expectComma = true
		case isASupportedFunction(token.TokenValue):
			projections.tokenIterator.Drop()
			if function, err := projections.function(); err != nil {
				return Expressions{}, err
			} else {
				expressions = append(expressions, expressionWithFunction(function))
			}
			expectComma = true
		}
	}
	return Expressions{expressions: expressions}, nil
}

func (projections *Projections) function() (*Function, error) {
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

		for projections.tokenIterator.HasNext() && !projections.tokenIterator.Peek().Equals(",") {
			token := projections.tokenIterator.Next()
			switch {
			case expectClosingParentheses:
				if !token.Equals(")") {
					functionStack.Clear()
					return nil, tokenizer.Token{}, errors.New("expected a closing parentheses in projection")
				}
				closingParenthesesCount = closingParenthesesCount + 1
				if closingParenthesesCount == functionStack.Size() {
					expectClosingParentheses = false
				}
			case expectOpeningParentheses:
				if !token.Equals("(") {
					functionStack.Clear()
					return nil, tokenizer.Token{}, errors.New("expected an opening parentheses in projection")
				}
				expectOpeningParentheses = false
			case isASupportedFunction(token.TokenValue):
				functionStack.Push(token)
				expectOpeningParentheses = true
			case isASupportedColumn(token.TokenValue):
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
