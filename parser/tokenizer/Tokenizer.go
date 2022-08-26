package tokenizer

import "strings"

type Tokenizer struct {
	query string
}

func newTokenizer(query string) *Tokenizer {
	return &Tokenizer{query: query}
}

func (tokenizer *Tokenizer) tokenize() *Tokens {
	tokens := NewEmptyTokens()
	var token strings.Builder
	for _, ch := range tokenizer.query {
		switch {
		case isCharATokenSeparator(ch):
			tokens.Add(tokenFrom(token.String()))
			token.Reset()
		case ch == '\'':
			tokens.Add(tokenFrom(token.String()))
			token.Reset()
		case isCharAComparisonOperator(ch):
			if !isAComparisonOperator(token.String()) {
				tokens.Add(tokenFrom(token.String()))
				token.Reset()
			}
			token.WriteRune(ch)
		case ch == ',':
			tokens.Add(tokenFrom(token.String()))
			tokens.Add(NewToken(Comma, string(ch)))
			token.Reset()
		case ch == '(':
			tokens.Add(tokenFrom(token.String()))
			tokens.Add(NewToken(OpeningParentheses, string(ch)))
			token.Reset()
		case ch == ')':
			tokens.Add(tokenFrom(token.String()))
			tokens.Add(NewToken(ClosingParentheses, string(ch)))
			token.Reset()
		default:
			token.WriteRune(ch)
		}
	}
	tokens.Add(tokenFrom(token.String()))
	return tokens
}
