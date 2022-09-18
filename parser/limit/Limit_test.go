//go:build unit
// +build unit

package limit

import (
	"fmt"
	"goselect/parser/tokenizer"
	"testing"
)

func TestLimitWithoutLimitClause(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()

	limit, _ := NewLimit(tokens.Iterator())
	if limit != nil {
		t.Fatalf("Expected limit to be nil but was not")
	}
}

func TestLimitWithWordOtherThanLimit(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "unknown"))

	limit, _ := NewLimit(tokens.Iterator())
	if limit != nil {
		t.Fatalf("Expected limit to be nil but was not")
	}
}

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
	fmt.Println(err)
	if err == nil {
		t.Fatalf("Expected an error with floating point limit")
	}
}

func TestLimitWithIllegalLimit1(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Limit, "limit"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "10@123"))

	_, err := NewLimit(tokens.Iterator())
	if err == nil {
		t.Fatalf("Expected an error with floating point limit")
	}
}

func TestLimitWithIllegalLimit2(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Limit, "limit"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "-10"))

	_, err := NewLimit(tokens.Iterator())
	if err == nil {
		t.Fatalf("Expected an error negative limit")
	}
}

func TestLimitWithAValue(t *testing.T) {
	tokens := tokenizer.NewEmptyTokens()
	tokens.Add(tokenizer.NewToken(tokenizer.Limit, "limit"))
	tokens.Add(tokenizer.NewToken(tokenizer.RawString, "10"))

	limit, _ := NewLimit(tokens.Iterator())
	if limit.Limit != 10 {
		t.Fatalf("Expected limit to be %v, but received %v", 10, limit.Limit)
	}
}
