package parser

import "testing"

func TestTokenIteratorWithNoNextToken(t *testing.T) {
	tokens := newEmptyTokens()
	iterator := tokens.iterator()

	if iterator.hasNext() == true {
		t.Fatalf("Expected hasNext to return false but it returned true")
	}
}

func TestTokenIteratorWithANextToken(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "select"))
	iterator := tokens.iterator()

	if iterator.hasNext() == false {
		t.Fatalf("Expected hasNext to return true but it returned false")
	}
}

func TestTokenIteratorWithNextTokenValue(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "select"))
	iterator := tokens.iterator()

	expectedToken := "select"
	actualToken := iterator.next()

	if expectedToken != actualToken.tokenValue {
		t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
	}
	if RawString != actualToken.tokenType {
		t.Fatalf("Expected token type to be %v, received %v", RawString, actualToken)
	}
}

func TestTokenIteratorWithMultipleTokens(t *testing.T) {
	tokens := newEmptyTokens()
	tokens.add(newToken(RawString, "select"))
	tokens.add(newToken(RawString, "name"))
	tokens.add(newToken(RawString, "from"))
	tokens.add(newToken(RawString, "/home"))

	iterator := tokens.iterator()

	expectedTokenPresence := []bool{true, true, true, true, false}
	expectedTokens := []string{"select", "name", "from", "/home"}

	for count := 1; count <= len(expectedTokenPresence); count++ {
		hasNext := iterator.hasNext()
		expectedHasNext := expectedTokenPresence[count-1]

		if expectedHasNext != hasNext {
			t.Fatalf("Expected hasNext to be %v, received %v", expectedHasNext, hasNext)
		}
		if hasNext {
			actualToken := iterator.next()
			expectedToken := expectedTokens[count-1]

			if expectedToken != actualToken.tokenValue {
				t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken.tokenValue)
			}
		}
	}
}
