package tokenizer

import (
	"testing"
)

func TestTokenizerWithTokenCount(t *testing.T) {
	tokenizer := NewTokenizer("select fName from /home/apps")
	tokens := tokenizer.Tokenize()

	expectedTokenCount := 4
	actualTokenCount := tokens.count()

	if expectedTokenCount != actualTokenCount {
		t.Fatalf("Expected token count %v, received %v", expectedTokenCount, actualTokenCount)
	}
}

func TestTokenizerWithAllTokens1(t *testing.T) {
	tokenizer := NewTokenizer("select fName from /home/apps")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "fName", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens2(t *testing.T) {
	tokenizer := NewTokenizer("select fName,fSize from /home/apps")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "fName", ",", "fSize", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens3(t *testing.T) {
	tokenizer := NewTokenizer("select fName,   fSize from    /home/apps")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "fName", ",", "fSize", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be: %v, received: %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens4(t *testing.T) {
	tokenizer := NewTokenizer("select * from /home/apps")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "*", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens5(t *testing.T) {
	tokenizer := NewTokenizer("select name, length(name),UPPER( name ) from /home/apps")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "name", ",", "length", "(", "name", ")", ",", "UPPER", "(", "name", ")", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens6(t *testing.T) {
	tokenizer := NewTokenizer("select name, rand() from /home/apps order by rand() limit 10")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "name", ",", "rand", "(", ")", "from", "/home/apps", "order", "by", "rand", "(", ")", "limit", "10"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens7(t *testing.T) {
	tokenizer := NewTokenizer("select COUNT(*), MIN(size) from /home/apps")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "COUNT", "(", "*", ")", ",", "MIN", "(", "size", ")", "from", "/home/apps"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]

		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens8(t *testing.T) {
	tokenizer := NewTokenizer("select size from /home/apps where name=*.txt order by modified")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "size", "from", "/home/apps", "where", "name=*.txt", "order", "by", "modified"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]
		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens9(t *testing.T) {
	tokenizer := NewTokenizer("select size from /home/apps where name='*.txt' order by modified")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "size", "from", "/home/apps", "where", "name=", "*.txt", "order", "by", "modified"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]
		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}

func TestTokenizerWithAllTokens10(t *testing.T) {
	tokenizer := NewTokenizer("select 1 * 2, name from /home/apps where size>1000")
	tokens := tokenizer.Tokenize()

	iterator := tokens.Iterator()
	expectedTokens := []string{"select", "1", "*", "2", ",", "name", "from", "/home/apps", "where", "size>1000"}

	for count := 1; count <= len(expectedTokens); count++ {
		actualToken := iterator.Next()
		expectedToken := expectedTokens[count-1]
		if expectedToken != actualToken.TokenValue {
			t.Fatalf("Expected token to be %v, received %v", expectedToken, actualToken)
		}
	}
}
