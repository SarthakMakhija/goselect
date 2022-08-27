package limit

import (
	"errors"
	"goselect/parser/errors/messages"
	"goselect/parser/tokenizer"
	"strconv"
	"strings"
)

type Limit struct {
	Limit   uint32
	defined bool
}

func NewLimit(iterator *tokenizer.TokenIterator) (*Limit, error) {
	if !iterator.HasNext() {
		return nil, nil
	}
	if !iterator.Next().Equals("limit") {
		return nil, nil
	}
	if !iterator.HasNext() {
		return nil, errors.New(messages.ErrorMessageLimitValue)
	}
	if strings.Contains(iterator.Peek().TokenValue, ".") {
		return nil, errors.New(messages.ErrorMessageLimitValueInt)
	}
	token := iterator.Next()
	if value, err := strconv.ParseUint(token.TokenValue, 10, 32); err != nil {
		return nil, err
	} else {
		return &Limit{Limit: uint32(value), defined: true}, nil
	}
}
