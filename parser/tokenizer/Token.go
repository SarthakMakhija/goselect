package tokenizer

import (
	"regexp"
	"strings"
)

const (
	RawString          int = iota
	Comma                  = 1
	From                   = 2
	Where                  = 3
	OpeningParentheses     = 4
	ClosingParentheses     = 5
	Order                  = 6
	By                     = 7
	AscendingOrder         = 8
	DescendingOrder        = 9
	Limit                  = 10
	Numeric                = 11
	FloatingPoint          = 12
	Boolean                = 13
)

var numericRegexp, _ = regexp.Compile("^[-+]?(?:0|[1-9][0-9]*)$")
var floatingNumbersRegexp, _ = regexp.Compile("^(?:[-+]?[0-9]+)?(?:\\.[0-9]+)?(?:[eE][+\\-]?[0-9]+)?$")
var booleanRegexp, _ = regexp.Compile("^true|false|y|n$")

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
	casedToken := strings.ToLower(token)
	switch {
	case casedToken == "from":
		return NewToken(From, token)
	case casedToken == "where":
		return NewToken(Where, token)
	case casedToken == "order":
		return NewToken(Order, token)
	case casedToken == "by":
		return NewToken(By, token)
	case casedToken == "asc":
		return NewToken(AscendingOrder, token)
	case casedToken == "desc":
		return NewToken(DescendingOrder, token)
	case casedToken == "limit":
		return NewToken(Limit, token)
	default:
		return NewToken(determineTokenType(casedToken), token)
	}
}

func determineTokenType(token string) int {
	if numericRegexp.MatchString(token) {
		return Numeric
	}
	if floatingNumbersRegexp.MatchString(token) {
		return FloatingPoint
	}
	if booleanRegexp.MatchString(token) {
		return Boolean
	}
	return RawString
}

func (token Token) Equals(value string) bool {
	return strings.EqualFold(strings.ToLower(token.TokenValue), strings.ToLower(value))
}

func (token Token) isNumeric() bool {
	return token.TokenType == Numeric
}

func (token Token) isFloatingPoint() bool {
	return token.TokenType == FloatingPoint
}

func (token Token) isBoolean() bool {
	return token.TokenType == Boolean
}
