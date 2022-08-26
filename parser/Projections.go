package parser

import (
	"errors"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

type Projections struct {
	tokenIterator *TokenIterator
}

func newProjections(tokenIterator *TokenIterator) *Projections {
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

	for projections.tokenIterator.hasNext() && !projections.tokenIterator.peek().equals("from") {
		token := projections.tokenIterator.next()
		switch {
		case expectComma:
			if !token.equals(",") {
				return Expressions{}, errors.New("expected a comma in projection list")
			}
			expectComma = false
		case isAWildcard(token.tokenValue):
			columns = append(columns, expressionsWithColumns(columnsOnWildcard())...)
			expectComma = true
		case isASupportedColumn(token.tokenValue):
			columns = append(columns, expressionWithColumn(token.tokenValue))
			expectComma = true
		case isASupportedFunction(token.tokenValue):
			projections.tokenIterator.drop()
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
	var column Token

	for projections.tokenIterator.hasNext() && !projections.tokenIterator.peek().equals(",") {
		token := projections.tokenIterator.next()
		switch {
		case expectOpeningParentheses:
			if !token.equals("(") {
				functionStack.Clear()
				return nil, errors.New("expected an opening parentheses in projection")
			}
			expectOpeningParentheses = false
		case expectClosingParentheses:
			if !token.equals(")") {
				functionStack.Clear()
				return nil, errors.New("expected a closing parentheses in projection")
			}
			expectClosingParentheses = true
		case isASupportedFunction(token.tokenValue):
			functionStack.Push(token)
			expectOpeningParentheses = true
		case isASupportedColumn(token.tokenValue):
			expectOpeningParentheses = false
			expectClosingParentheses = true
			column = token
		}
	}
	functionToken, _ := functionStack.Pop()
	var rootFunction = &Function{
		id:   (functionToken.(Token)).tokenValue,
		left: &Expression{column: column.tokenValue},
	}
	for !functionStack.Empty() {
		functionToken, _ := functionStack.Pop()
		rootFunction = &Function{
			id:   (functionToken.(Token)).tokenValue,
			left: &Expression{function: rootFunction},
		}
	}
	return rootFunction, nil
}
