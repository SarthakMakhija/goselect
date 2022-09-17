package limit

import (
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
	"strconv"
)

type Limit struct {
	Limit uint32
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
	token := iterator.Next()
	limitValue, err := strconv.ParseUint(token.TokenValue, 10, 32)
	if err != nil {
		return nil, fmt.Errorf(messages.ErrorMessageLimitValueIntWithExistingError, err)
	}
	return &Limit{Limit: uint32(limitValue)}, nil
}
