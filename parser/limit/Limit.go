package limit

import (
	"errors"
	"goselect/parser/tokenizer"
	"strconv"
	"strings"
)

type Limit struct {
	limit   uint32
	defined bool
}

func NewLimit(iterator *tokenizer.TokenIterator) (Limit, error) {
	if iterator.HasNext() && !iterator.Next().Equals("limit") {
		return Limit{defined: false}, nil
	}
	if !iterator.HasNext() {
		return Limit{}, errors.New("need to define limit value")
	}
	if strings.Contains(iterator.Peek().TokenValue, ".") {
		return Limit{}, errors.New("limit must be an integer")
	}
	token := iterator.Next()
	if value, err := strconv.ParseUint(token.TokenValue, 10, 32); err != nil {
		return Limit{}, err
	} else {
		return Limit{limit: uint32(value), defined: true}, nil
	}
}
