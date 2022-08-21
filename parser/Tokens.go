package parser

type Tokens struct {
	tokens []string
}

func newEmptyTokens() *Tokens {
	return &Tokens{}
}

func newTokens(tokens []string) *Tokens {
	return &Tokens{tokens: tokens}
}

func (tokens *Tokens) add(token string) {
	tokens.tokens = append(tokens.tokens, token)
}

func (tokens *Tokens) count() int {
	return len(tokens.tokens)
}

func (tokens *Tokens) iterator() *TokenIterator {
	return &TokenIterator{index: 0, tokens: tokens.tokens}
}

type TokenIterator struct {
	index  int
	tokens []string
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
