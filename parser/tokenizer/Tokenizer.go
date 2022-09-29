package tokenizer

import (
	"strings"
)

type Tokenizer struct {
	query string
}

func NewTokenizer(query string) *Tokenizer {
	return &Tokenizer{query: query}
}

func (tokenizer *Tokenizer) Tokenize() *Tokens {
	tokens := NewEmptyTokens()
	var token strings.Builder
	for index := 0; index < len(tokenizer.query); index++ {
		ch := rune(tokenizer.query[index])
		switch {
		case isCharATokenSeparator(ch):
			tokens.Add(tokenFrom(token.String()))
			token.Reset()
		case ch == '\'':
			tokens.Add(tokenFrom(token.String()))
			literal, newIndex := tokenizer.readSingleQuotedLiteralFrom(index + 1)
			index = newIndex
			tokens.Add(literal)
			token.Reset()
		case ch == '"':
			tokens.Add(tokenFrom(token.String()))
			literal, newIndex := tokenizer.readDoubleQuotedLiteralFrom(index + 1)
			index = newIndex
			tokens.Add(literal)
			token.Reset()
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

func (tokenizer *Tokenizer) readSingleQuotedLiteralFrom(index int) (Token, int) {
	token, nextIndex := tokenizer.readQuotedLiteral(index, func(ch rune) bool {
		return isCharATokenSeparator(ch) || ch == '\''
	})
	return tokenFrom(token.String()), nextIndex
}

func (tokenizer *Tokenizer) readDoubleQuotedLiteralFrom(index int) (Token, int) {
	token, nextIndex := tokenizer.readQuotedLiteral(index, func(ch rune) bool {
		return isCharATokenSeparator(ch) || ch == '"'
	})
	return tokenFrom("\"" + token.String() + "\""), nextIndex
}

func (tokenizer *Tokenizer) readQuotedLiteral(index int, breakOn func(ch rune) bool) (strings.Builder, int) {
	var token strings.Builder

	runningIndex := index
	for ; runningIndex < len(tokenizer.query); runningIndex++ {
		ch := rune(tokenizer.query[runningIndex])
		if breakOn(ch) {
			break
		} else {
			token.WriteRune(ch)
		}
	}
	return token, runningIndex
}
