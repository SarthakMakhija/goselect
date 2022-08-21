package parser

type Tokens struct {
	tokens []string
}

type TokenIterator struct {
	index  int
	tokens []string
}

func newTokens(tokens []string) Tokens {
	return Tokens{tokens: tokens}
}

func (tokens Tokens) iterator() *TokenIterator {
	return &TokenIterator{index: 0, tokens: tokens.tokens}
}

func (tokenIterator *TokenIterator) hasNext() bool {
	if tokenIterator.index < len(tokenIterator.tokens) {
		return true
	}
	return false
}

func (tokenIterator *TokenIterator) next() string {
	token := tokenIterator.tokens[tokenIterator.index]
	tokenIterator.index = tokenIterator.index + 1
	return token
}
