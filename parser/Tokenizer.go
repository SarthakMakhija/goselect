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
		if isATokenSeparator(ch) {
			if token := token.String(); len(token) != 0 {
				tokens.add(token)
			}
			tokens.add(string(ch))
			token.Reset()
		} else {
			token.WriteRune(ch)
		}
	}
	tokens.add(token.String())
	return tokens
}
