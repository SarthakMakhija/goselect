package parser

import (
	"testing"
)

func TestTokenizerWithTokenCount(t *testing.T) {
	tokenizer := newTokenizer("select fName from /home/apps")
	tokens := tokenizer.tokenize()

	expectedTokenCount := 7
	actualTokenCount := tokens.count()

	if expectedTokenCount != actualTokenCount {
		t.Fatalf("Expected token count %v, received %v", expectedTokenCount, actualTokenCount)
	}
}

func TestTokenizerWithAllTokens1(t *testing.T) {
	tokenizer := newTokenizer("select fName from /home/apps")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", " ", "fName", " ", "from", " ", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens2(t *testing.T) {
	tokenizer := newTokenizer("select fName,fSize from /home/apps")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", " ", "fName", ",", "fSize", " ", "from", " ", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens3(t *testing.T) {
	tokenizer := newTokenizer("select fName,   fSize from    /home/apps")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", " ", "fName", ",", " ", " ", " ", "fSize", " ", "from", " ", " ", " ", " ", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken {
			t.Fatalf("Expected token to be: %v, received: %v", expectedToken, actualToken)
		}
	}
}
