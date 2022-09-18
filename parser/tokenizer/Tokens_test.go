//go:build unit
// +build unit

package tokenizer

import "testing"

func TestTokenIteratorWithNoNextToken(t *testing.T) {
	tokens := NewEmptyTokens()
	iterator := tokens.Iterator()

	if iterator.HasNext() == true {
		t.Fatalf("Expected HasNext to return false but it returned true")
	}
}

func TestTokenIteratorWithANextToken(t *testing.T) {
	tokens := NewEmptyTokens()
	tokens.Add(NewToken(RawString, "select"))
	iterator := tokens.Iterator()

	if iterator.HasNext() == false {
		t.Fatalf("Expected HasNext to return true but it returned false")
	}
}

func TestTokenIteratorWithNextTokenValue(t *testing.T) {
	tokens := NewEmptyTokens()
	tokens.Add(NewToken(RawString, "select"))
	iterator := tokens.Iterator()

	expectedToken := "select"
	actualToken := iterator.Next()

	if expectedToken != actualToken.TokenValue {
		t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
	}
	if RawString != actualToken.TokenType {
		t.Fatalf("Expected token type to be %v, received %v", RawString, actualToken)
	}
}

func TestTokenIteratorWithMultipleTokens(t *testing.T) {
	tokens := NewEmptyTokens()
	tokens.Add(NewToken(RawString, "select"))
	tokens.Add(NewToken(RawString, "name"))
	tokens.Add(NewToken(RawString, "from"))
	tokens.Add(NewToken(RawString, "/home"))

	iterator := tokens.Iterator()

	expectedTokenPresence := []bool{true, true, true, true, false}
	expectedTokens := []string{"select", "name", "from", "/home"}

	for count := 1; count <= len(expectedTokenPresence); count++ {
		hasNext := iterator.HasNext()
		expectedHasNext := expectedTokenPresence[count-1]

		if expectedHasNext != hasNext {
			t.Fatalf("Expected HasNext to be %v, received %v", expectedHasNext, hasNext)
		}
		if hasNext {
			actualToken := iterator.Next()
			expectedToken := expectedTokens[count-1]

			if expectedToken != actualToken.TokenValue {
				t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken.TokenValue)
			}
		}
	}
}
