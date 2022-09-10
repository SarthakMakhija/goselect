package order

import (
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
	"strconv"
)

type Order struct {
	Attributes []AttributeRef
	directions []bool //true signifies ascending, false signified descending
}

type AttributeRef struct {
	ProjectionPosition int
}

const (
	sortingDirectionAscending  int = iota
	sortingDirectionDescending     = 1
)

func NewOrder(iterator *tokenizer.TokenIterator, projectionCount int) (*Order, error) {
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

	var attributes []AttributeRef
	var directions []bool
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
		default:
			if projectionPosition, err := strconv.Atoi(token.TokenValue); err != nil {
				return nil, fmt.Errorf(messages.ErrorMessageNonZeroPositivePositionsWithExistingError, err)
			} else {
				if projectionPosition <= 0 {
					return nil, errors.New(messages.ErrorMessageNonZeroPositivePositions)
				}
				if projectionPosition > projectionCount {
					return nil, fmt.Errorf(messages.ErrorMessageOrderByPositionOutOfRange, 1, projectionCount)
				}
				if projectionPosition <= projectionCount {
					attributes = append(attributes, AttributeRef{ProjectionPosition: projectionPosition})
					if sortingDirection(iterator) == sortingDirectionDescending {
						directions = append(directions, false)
					} else {
						directions = append(directions, true)
					}
					expectComma = true
				}
			}
		}
	}
	if len(attributes) == 0 {
		return nil, errors.New(messages.ErrorMessageMissingOrderByAttributes)
	}
	return &Order{Attributes: attributes, directions: directions}, nil
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

func (order Order) IsAscendingAt(index int) bool {
	if index < len(order.directions) {
		return order.directions[index]
	}
	return false
}
