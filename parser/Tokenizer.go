package parser

import "strings"

type Tokenizer struct {
	query string
}

func newTokenizer(query string) *Tokenizer {
	return &Tokenizer{query: query}
}

func (tokenizer *Tokenizer) tokenize() *Tokens {
	tokens := newEmptyTokens()
	var token strings.Builder
	for _, ch := range tokenizer.query {
		switch {
		case isCharATokenSeparator(ch):
			tokens.add(tokenFrom(token.String()))
			token.Reset()
		case ch == '\'':
			tokens.add(tokenFrom(token.String()))
			token.Reset()
		case isCharAComparisonOperator(ch):
			if !isAComparisonOperator(token.String()) {
				tokens.add(tokenFrom(token.String()))
				token.Reset()
			}
			token.WriteRune(ch)
		case ch == ',':
			tokens.add(tokenFrom(token.String()))
			tokens.add(newToken(Comma, string(ch)))
			token.Reset()
		case ch == '(':
			tokens.add(tokenFrom(token.String()))
			tokens.add(newToken(OpeningParentheses, string(ch)))
			token.Reset()
		case ch == ')':
			tokens.add(tokenFrom(token.String()))
			tokens.add(newToken(ClosingParentheses, string(ch)))
			token.Reset()
		default:
			token.WriteRune(ch)
		}
	}
	tokens.add(tokenFrom(token.String()))
	return tokens
}
