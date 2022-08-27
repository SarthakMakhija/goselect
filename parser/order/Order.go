package order

import (
	"errors"
	"goselect"
	"goselect/parser/projection"
	"goselect/parser/tokenizer"
)

type Order struct {
	ascendingColumns  []string
	descendingColumns []string
}

func NewOrder(iterator *tokenizer.TokenIterator) (Order, error) {
	if iterator.HasNext() && !iterator.Next().Equals("order") {
		return Order{}, nil
	}
	if iterator.HasNext() && !iterator.Next().Equals("by") {
		return Order{}, errors.New(goselect.ErrorMessageMissingBy)
	}

	var ascendingColumns, descendingColumns []string
	var expectComma bool
	for iterator.HasNext() && !iterator.Peek().Equals("limit") {
		token := iterator.Next()
		switch {
		case expectComma:
			if !token.Equals(",") {
				return Order{}, errors.New(goselect.ErrorMessageMissingCommaOrderBy)
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
		return Order{}, errors.New(goselect.ErrorMessageMissingOrderByColumns)
	}
	return Order{ascendingColumns: ascendingColumns, descendingColumns: descendingColumns}, nil
}
