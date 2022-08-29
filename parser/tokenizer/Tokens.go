package tokenizer

type Tokens struct {
	tokens []Token
}

func NewEmptyTokens() *Tokens {
	return &Tokens{}
}

func (tokens *Tokens) Add(token Token) {
	if !token.isEmpty() {
		tokens.tokens = append(tokens.tokens, token)
	}
}

func (tokens *Tokens) count() int {
	return len(tokens.tokens)
}

func (tokens *Tokens) Iterator() *TokenIterator {
	return &TokenIterator{index: 0, tokens: tokens.tokens}
}

type TokenIterator struct {
	index  int
	tokens []Token
}

func (tokenIterator *TokenIterator) HasNext() bool {
	if tokenIterator.index < len(tokenIterator.tokens) {
		return true
	}
	return false
}

func (tokenIterator *TokenIterator) Next() Token {
	token := tokenIterator.Peek()
	tokenIterator.index = tokenIterator.index + 1
	return token
}

func (tokenIterator *TokenIterator) Peek() Token {
	token := tokenIterator.tokens[tokenIterator.index]
	return token
}
