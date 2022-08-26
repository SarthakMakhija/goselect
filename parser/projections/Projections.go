package projections

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
projections: columns Or functions Or expressions
columns: 	 name, size etc
functions: 	 min(size), lower(name), min(count(size)) etc
expressions: 2 + 3, 2 > 3 etc
*/
func (projections *Projections) all() (Expressions, error) {
	var columns []Expression
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
			columns = append(columns, expressionsWithColumns(columnsOnWildcard())...)
			expectComma = true
		case isASupportedColumn(token.TokenValue):
			columns = append(columns, expressionWithColumn(token.TokenValue))
			expectComma = true
		case isASupportedFunction(token.TokenValue):
			projections.tokenIterator.Drop()
			fn, err := projections.function()
			if err != nil {
				return Expressions{}, err
			}
			columns = append(columns, expressionWithFunction(fn))
		}
	}
	return Expressions{expressions: columns}, nil
}

func (projections *Projections) function() (*Function, error) {
	var expectOpeningParentheses bool
	var expectClosingParentheses bool

	functionStack := linkedliststack.New()
	var column tokenizer.Token

	for projections.tokenIterator.HasNext() && !projections.tokenIterator.Peek().Equals(",") {
		token := projections.tokenIterator.Next()
		switch {
		case expectOpeningParentheses:
			if !token.Equals("(") {
				functionStack.Clear()
				return nil, errors.New("expected an opening parentheses in projection")
			}
			expectOpeningParentheses = false
		case expectClosingParentheses:
			if !token.Equals(")") {
				functionStack.Clear()
				return nil, errors.New("expected a closing parentheses in projection")
			}
			expectClosingParentheses = true
		case isASupportedFunction(token.TokenValue):
			functionStack.Push(token)
			expectOpeningParentheses = true
		case isASupportedColumn(token.TokenValue):
			expectOpeningParentheses = false
			expectClosingParentheses = true
			column = token
		}
	}
	functionToken, _ := functionStack.Pop()
	var rootFunction = &Function{
		id:   (functionToken.(tokenizer.Token)).TokenValue,
		left: &Expression{column: column.TokenValue},
	}
	for !functionStack.Empty() {
		functionToken, _ := functionStack.Pop()
		rootFunction = &Function{
			id:   (functionToken.(tokenizer.Token)).TokenValue,
			left: &Expression{function: rootFunction},
		}
	}
	return rootFunction, nil
}
