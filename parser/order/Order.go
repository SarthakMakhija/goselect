package order

import (
	"errors"
	"goselect/parser"
	"goselect/parser/projection"
	"goselect/parser/tokenizer"
	"strconv"
)

type Order struct {
	ascendingColumns  []ColumnRef
	descendingColumns []ColumnRef
}

type ColumnRef struct {
	name               string
	projectionPosition int
}

const (
	sortingDirectionAscending  int = iota
	sortingDirectionDescending     = 1
)

func NewOrder(iterator *tokenizer.TokenIterator, projectionCount int) (*Order, error) {
	if iterator.HasNext() && !iterator.Next().Equals("order") {
		return nil, nil
	}
	if iterator.HasNext() && !iterator.Next().Equals("by") {
		return nil, errors.New(parser.ErrorMessageMissingBy)
	}

	var ascendingColumns, descendingColumns []ColumnRef
	var expectComma bool
	for iterator.HasNext() && !iterator.Peek().Equals("limit") {
		token := iterator.Next()
		switch {
		case expectComma:
			if !token.Equals(",") {
				return nil, errors.New(parser.ErrorMessageMissingCommaOrderBy)
			}
			expectComma = false
		case projection.IsASupportedColumn(token.TokenValue):
			if sortingDirection(iterator) == sortingDirectionDescending {
				descendingColumns = append(descendingColumns, ColumnRef{name: token.TokenValue, projectionPosition: -1})
			} else {
				ascendingColumns = append(ascendingColumns, ColumnRef{name: token.TokenValue, projectionPosition: -1})
			}
			expectComma = true
		default:
			if projectionPosition, err := strconv.Atoi(token.TokenValue); err != nil {
				return nil, err
			} else {
				if projectionPosition <= projectionCount {
					if sortingDirection(iterator) == sortingDirectionDescending {
						descendingColumns = append(descendingColumns, ColumnRef{projectionPosition: projectionPosition})
					} else {
						ascendingColumns = append(ascendingColumns, ColumnRef{projectionPosition: projectionPosition})
					}
					expectComma = true
				}
			}
		}
	}
	if len(ascendingColumns) == 0 && len(descendingColumns) == 0 {
		return nil, errors.New(parser.ErrorMessageMissingOrderByColumns)
	}
	return &Order{ascendingColumns: ascendingColumns, descendingColumns: descendingColumns}, nil
}

func sortingDirection(iterator *tokenizer.TokenIterator) int {
	if iterator.HasNext() && iterator.Peek().Equals("desc") {
		iterator.Next()
		return sortingDirectionDescending
	} else {
		if iterator.HasNext() && iterator.Peek().Equals("asc") {
			iterator.Next()
		}
		return sortingDirectionAscending
	}
}
