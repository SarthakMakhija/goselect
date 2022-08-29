package order

import (
	"errors"
	"goselect/parser/context"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
	"strconv"
)

type Order struct {
	Attributes []AttributeRef
	Directions []bool //true signifies ascending, false signified descending
}

type AttributeRef struct {
	Name               string
	ProjectionPosition int
}

const (
	sortingDirectionAscending  int = iota
	sortingDirectionDescending     = 1
)

func NewOrder(iterator *tokenizer.TokenIterator, context *context.ParsingApplicationContext, projectionCount int) (*Order, error) {
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
		case context.IsASupportedAttribute(token.TokenValue):
			attributes = append(attributes, AttributeRef{Name: token.TokenValue, ProjectionPosition: -1})
			if sortingDirection(iterator) == sortingDirectionDescending {
				directions = append(directions, false)
			} else {
				directions = append(directions, true)
			}
			expectComma = true
		default:
			if projectionPosition, err := strconv.Atoi(token.TokenValue); err != nil {
				return nil, err
			} else {
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
	return &Order{Attributes: attributes, Directions: directions}, nil
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
