package order

import (
	"errors"
	"goselect/parser/context"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
	"strconv"
)

type Order struct {
	AscendingColumns  []ColumnRef
	DescendingColumns []ColumnRef
}

type ColumnRef struct {
	Name               string
	ProjectionPosition int
}

const (
	sortingDirectionAscending  int = iota
	sortingDirectionDescending     = 1
)

func NewOrder(iterator *tokenizer.TokenIterator, context *context.Context, projectionCount int) (*Order, error) {
	if !iterator.HasNext() {
		return nil, nil
	}
	if iterator.HasNext() && !iterator.Peek().Equals("order") {
		return nil, nil
	}
	iterator.Next()

	if iterator.HasNext() && !iterator.Peek().Equals("by") {
		return nil, errors.New(messages.ErrorMessageMissingBy)
	}

	var ascendingColumns, descendingColumns []ColumnRef
	var expectComma bool
	for iterator.HasNext() && !iterator.Peek().Equals("limit") {
		token := iterator.Next()
		if token.Equals("by") {
			continue
		}
		switch {
		case expectComma:
			if !token.Equals(",") {
				return nil, errors.New(messages.ErrorMessageMissingCommaOrderBy)
			}
			expectComma = false
		case context.IsASupportedAttribute(token.TokenValue):
			if sortingDirection(iterator) == sortingDirectionDescending {
				descendingColumns = append(descendingColumns, ColumnRef{Name: token.TokenValue, ProjectionPosition: -1})
			} else {
				ascendingColumns = append(ascendingColumns, ColumnRef{Name: token.TokenValue, ProjectionPosition: -1})
			}
			expectComma = true
		default:
			if projectionPosition, err := strconv.Atoi(token.TokenValue); err != nil {
				return nil, err
			} else {
				if projectionPosition <= projectionCount {
					if sortingDirection(iterator) == sortingDirectionDescending {
						descendingColumns = append(descendingColumns, ColumnRef{ProjectionPosition: projectionPosition})
					} else {
						ascendingColumns = append(ascendingColumns, ColumnRef{ProjectionPosition: projectionPosition})
					}
					expectComma = true
				}
			}
		}
	}
	if len(ascendingColumns) == 0 && len(descendingColumns) == 0 {
		return nil, errors.New(messages.ErrorMessageMissingOrderByColumns)
	}
	return &Order{AscendingColumns: ascendingColumns, DescendingColumns: descendingColumns}, nil
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
