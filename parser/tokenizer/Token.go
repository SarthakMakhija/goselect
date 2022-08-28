package tokenizer

import "strings"

const (
	RawString          int = iota
	Comma                  = 1
	From                   = 2
	Where                  = 3
	Operator               = 4
	OpeningParentheses     = 5
	ClosingParentheses     = 6
	ArithmeticOperator     = 7
	And                    = 8
	Or                     = 9
	Not                    = 10
	Order                  = 11
	By                     = 12
	AscendingOrder         = 13
	DescendingOrder        = 14
	Limit                  = 15
)

type Token struct {
	TokenType  int
	TokenValue string
}

func (token Token) isEmpty() bool {
	return len(token.TokenValue) == 0
}

func NewToken(tokenType int, tokenValue string) Token {
	return Token{TokenType: tokenType, TokenValue: tokenValue}
}

func tokenFrom(token string) Token {
	switch {
	case token == "from":
		return NewToken(From, token)
	case token == "where":
		return NewToken(Where, token)
	case token == "or":
		return NewToken(Or, token)
	case token == "and":
		return NewToken(And, token)
	case token == "not":
		return NewToken(Not, token)
	case token == "order":
		return NewToken(Order, token)
	case token == "by":
		return NewToken(By, token)
	case token == "asc":
		return NewToken(AscendingOrder, token)
	case token == "desc":
		return NewToken(DescendingOrder, token)
	case token == "limit":
		return NewToken(Limit, token)
	case isArithmeticOperator(token):
		return NewToken(ArithmeticOperator, token)
	case isAComparisonOperator(token):
		return NewToken(Operator, token)
	default:
		return NewToken(RawString, token)
	}
}

func (token Token) Equals(value string) bool {
	return strings.ToLower(token.TokenValue) == strings.ToLower(value)
}
