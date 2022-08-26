package parser

type Tokens struct {
	tokens []Token
}

func newEmptyTokens() *Tokens {
	return &Tokens{}
}

func (tokens *Tokens) add(token Token) {
	if !token.isEmpty() {
		tokens.tokens = append(tokens.tokens, token)
	}
}

func (tokens *Tokens) count() int {
	return len(tokens.tokens)
}

func (tokens *Tokens) iterator() *TokenIterator {
	return &TokenIterator{index: 0, tokens: tokens.tokens}
}

type TokenIterator struct {
	index  int
	tokens []Token
}

func (tokenIterator *TokenIterator) hasNext() bool {
	if tokenIterator.index < len(tokenIterator.tokens) {
		return true
	}
	return false
}

func (tokenIterator *TokenIterator) next() Token {
	token := tokenIterator.peek()
	tokenIterator.index = tokenIterator.index + 1
	return token
}

func (tokenIterator *TokenIterator) peek() Token {
	token := tokenIterator.tokens[tokenIterator.index]
	return token
}
