package parser

import "errors"

type Projections struct {
	tokenIterator *TokenIterator
}

func newProjections(tokenIterator *TokenIterator) *Projections {
	return &Projections{tokenIterator: tokenIterator}
}

/*
projections: fields Or functions Or expressions
fields: 	 name, size etc
functions: 	 min(size), lower(name), min(count(size)) etc
expressions: 2 + 3, 2 > 3 etc
*/
func (projections *Projections) all() ([]string, error) {
	var columns []string
	var expectComma bool

	for projections.tokenIterator.hasNext() && !projections.tokenIterator.peek().equals("from") {
		token := projections.tokenIterator.next()
		switch {
		case expectComma:
			if !token.equals(",") {
				return []string{}, errors.New("expected a comma in projection list")
			}
			expectComma = false
		case isAWildcard(token.tokenValue):
			columns = append(columns, columnsOnWildcard()...)
			expectComma = true
		case isASupportedColumn(token.tokenValue):
			columns = append(columns, token.tokenValue)
			expectComma = true
		}
	}
	return columns, nil
}
