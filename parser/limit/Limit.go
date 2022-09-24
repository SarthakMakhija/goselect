package limit

import (
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
	"strconv"
	"strings"
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
	tokenValue := token.TokenValue
	if strings.HasPrefix(tokenValue, "+") {
		tokenValue = tokenValue[1:]
	}
	if len(tokenValue) == 0 {
		return nil, fmt.Errorf(messages.ErrorMessageLimitValueInt, token.TokenValue)
	}
	limitValue, err := strconv.ParseUint(tokenValue, 10, 32)
	if err != nil {
		return nil, fmt.Errorf(messages.ErrorMessageLimitValueInt, token.TokenValue)
	}
	return &Limit{Limit: uint32(limitValue)}, nil
}
