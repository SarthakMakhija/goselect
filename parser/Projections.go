package parser

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
func (projections *Projections) all() []string {
	var columns []string
	for projections.tokenIterator.hasNext() && !projections.tokenIterator.peek().equals("from") {
		token := projections.tokenIterator.next()
		if isAWildcard(token.tokenValue) {
			columns = append(columns, columnsOnWildcard()...)
		} else if isASupportedColumn(token.tokenValue) {
			columns = append(columns, token.tokenValue)
		}
	}
	return columns
}
