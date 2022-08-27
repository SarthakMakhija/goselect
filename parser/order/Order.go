package order

import (
	"errors"
	"goselect/parser"
	"goselect/parser/projection"
	"goselect/parser/tokenizer"
)

type Order struct {
	ascendingColumns  []string
	descendingColumns []string
}

func NewOrder(iterator *tokenizer.TokenIterator) (*Order, error) {
	if iterator.HasNext() && !iterator.Next().Equals("order") {
		return nil, nil
	}
	if iterator.HasNext() && !iterator.Next().Equals("by") {
		return nil, errors.New(parser.ErrorMessageMissingBy)
	}

	var ascendingColumns, descendingColumns []string
	var expectComma bool
	for iterator.HasNext() && !iterator.Peek().Equals("limit") {
		token := iterator.Next()
		switch {
		case expectComma:
			if !token.Equals(",") {
				return nil, errors.New(parser.ErrorMessageMissingCommaOrderBy)
			}
			expectComma = false
		default:
			if projection.IsASupportedColumn(token.TokenValue) {
				if iterator.HasNext() && iterator.Peek().Equals("desc") {
					descendingColumns = append(descendingColumns, token.TokenValue)
					iterator.Next()
				} else {
					ascendingColumns = append(ascendingColumns, token.TokenValue)
					if iterator.HasNext() && iterator.Peek().Equals("asc") {
						iterator.Next()
					}
				}
				expectComma = true
			}
		}
	}
	if len(ascendingColumns) == 0 && len(descendingColumns) == 0 {
		return nil, errors.New(parser.ErrorMessageMissingOrderByColumns)
	}
	return &Order{ascendingColumns: ascendingColumns, descendingColumns: descendingColumns}, nil
}
