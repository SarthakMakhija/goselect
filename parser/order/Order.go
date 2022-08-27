package order

import (
	"errors"
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
		return Order{}, nil
	}

	var ascendingColumns, descendingColumns []string
	var expectComma bool
	for iterator.HasNext() {
		token := iterator.Next()
		switch {
		case expectComma:
			if !token.Equals(",") {
				return Order{}, errors.New("expected a comma after 'order by' in column separator")
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
	return Order{ascendingColumns: ascendingColumns, descendingColumns: descendingColumns}, nil
}
