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
		if tokenizer.isTokenSeparator(ch) {
			tokens.add(token.String())
			tokens.add(string(ch))
			token.Reset()
		} else {
			token.WriteRune(ch)
		}
	}
	tokens.add(token.String())
	return tokens
}

func (tokenizer *Tokenizer) isTokenSeparator(ch rune) bool {
	if ch == ' ' || ch == ',' {
		return true
	}
	return false
}
