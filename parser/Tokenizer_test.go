package parser

import (
	"testing"
)

func TestTokenizerWithTokenCount(t *testing.T) {
	tokenizer := newTokenizer("select fName from /home/apps")
	tokens := tokenizer.tokenize()

	expectedTokenCount := 4
	actualTokenCount := tokens.count()

	if expectedTokenCount != actualTokenCount {
		t.Fatalf("Expected token count %v, received %v", expectedTokenCount, actualTokenCount)
	}
}

func TestTokenizerWithAllTokens1(t *testing.T) {
	tokenizer := newTokenizer("select fName from /home/apps")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", "fName", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.tokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens2(t *testing.T) {
	tokenizer := newTokenizer("select fName,fSize from /home/apps")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", "fName", ",", "fSize", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.tokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens3(t *testing.T) {
	tokenizer := newTokenizer("select fName,   fSize from    /home/apps")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", "fName", ",", "fSize", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.tokenValue {
			t.Fatalf("Expected token to be: %v, received: %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens4(t *testing.T) {
	tokenizer := newTokenizer("select * from /home/apps")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", "*", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.tokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens5(t *testing.T) {
	tokenizer := newTokenizer("select name, length(name),UPPER( name ) from /home/apps")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", "name", ",", "length", "(", "name", ")", ",", "UPPER", "(", "name", ")", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.tokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens6(t *testing.T) {
	tokenizer := newTokenizer("select name, rand() from /home/apps order by rand() limit 10")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", "name", ",", "rand", "(", ")", "from", "/home/apps", "order", "by", "rand", "(", ")", "limit", "10"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.tokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens7(t *testing.T) {
	tokenizer := newTokenizer("select COUNT(*), MIN(size) from /home/apps")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", "COUNT", "(", "*", ")", ",", "MIN", "(", "size", ")", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.tokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens8(t *testing.T) {
	tokenizer := newTokenizer("select size from /home/apps where name='*.txt' order by modified")
	tokens := tokenizer.tokenize()

	iterator := tokens.iterator()
	expectedTokens := []string{"select", "size", "from", "/home/apps", "where", "name", "=", "*.txt", "order", "by", "modified"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.next()
		expectedToken := expectedTokens[count-1]
		if expectedToken != actualToken.tokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}
