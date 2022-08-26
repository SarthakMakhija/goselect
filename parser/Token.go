package parser

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
	tokenType  int
	tokenValue string
}

func (token Token) isEmpty() bool {
	return len(token.tokenValue) == 0
}

func newToken(tokenType int, tokenValue string) Token {
	return Token{tokenType: tokenType, tokenValue: tokenValue}
}

func tokenFrom(token string) Token {
	switch {
	case token == "from":
		return newToken(From, token)
	case token == "where":
		return newToken(Where, token)
	case token == "or":
		return newToken(Or, token)
	case token == "and":
		return newToken(And, token)
	case token == "not":
		return newToken(Not, token)
	case token == "order":
		return newToken(Order, token)
	case token == "by":
		return newToken(By, token)
	case token == "asc":
		return newToken(AscendingOrder, token)
	case token == "desc":
		return newToken(DescendingOrder, token)
	case token == "limit":
		return newToken(Limit, token)
	case isArithmeticOperator(token):
		return newToken(ArithmeticOperator, token)
	case isAComparisonOperator(token):
		return newToken(Operator, token)
	default:
		return newToken(RawString, token)
	}
}

func (token Token) equals(value string) bool {
	return token.tokenValue == value
}
