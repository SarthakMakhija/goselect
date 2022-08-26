package limit

import (
	"goselect/parser/tokenizer"
	"testing"
)

func TestLimitWithoutAnyLimitDefined(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Limit, "limit"))

	_, err := NewLimit(tokens.Iterator())
	if err == nil {
		t.Fatalf("Expected an error with limit keyword without any limit")
	}
}

func TestLimitWithAFloatingValue(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Limit, "limit"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "10.12"))

	_, err := NewLimit(tokens.Iterator())
	if err == nil {
		t.Fatalf("Expected an error with floating point limit")
	}
}

func TestLimitWithIllegalLimit(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Limit, "limit"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "10@123"))

	_, err := NewLimit(tokens.Iterator())
	if err == nil {
		t.Fatalf("Expected an error with floating point limit")
	}
}

func TestLimitWithAValue(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Limit, "limit"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "10"))

	limit, _ := NewLimit(tokens.Iterator())
	if limit.limit != 10 {
		t.Fatalf("Expected limit to be %v, but received %v", 10, limit.limit)
	}
}
