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

func TestTokenTypeAsNumeric(t *testing.T) {
	numericTests := map[string]struct {
		input     string
		isNumeric bool
	}{
		"12 is numeric": {
			input:     "12",
			isNumeric: true,
		},
		"-12 is numeric": {
			input:     "-12",
			isNumeric: true,
		},
		"+12 is numeric": {
			input:     "+12",
			isNumeric: true,
		},
		"999999999 is numeric": {
			input:     "999999999",
			isNumeric: true,
		},
		"test is not numeric": {
			input:     "test",
			isNumeric: false,
		},
		"90a is not numeric": {
			input:     "90a",
			isNumeric: false,
		},
		"90.0 is not numeric": {
			input:     "90.0",
			isNumeric: false,
		},
	}

	for testName, input := range numericTests {
		t.Run(testName, func(t *testing.T) {
			token := tokenFrom(input.input)
			if input.isNumeric {
				if token.TokenType != Numeric {
					t.Fatalf("Expected %v to be numeric but was %v", input.input, token.TokenType)
				}
			}
		})
	}
}

func TestTokenTypeAsFloatingPoint(t *testing.T) {
	numericTests := map[string]struct {
		input           string
		isFloatingPoint bool
	}{
		"12 is not float": {
			input:           "12",
			isFloatingPoint: false,
		},
		"999999999 is float": {
			input:           "999999999",
			isFloatingPoint: false,
		},
		"test is not float": {
			input:           "test",
			isFloatingPoint: false,
		},
		"90a is not float": {
			input:           "90a",
			isFloatingPoint: false,
		},
		"90.0 is float": {
			input:           "90.0",
			isFloatingPoint: true,
		},
		"90.12 is float": {
			input:           "90.12",
			isFloatingPoint: true,
		},
		".123 is float": {
			input:           ".123",
			isFloatingPoint: true,
		},
		"0.123 is float": {
			input:           "0.123",
			isFloatingPoint: true,
		},
		"-0.123 is float": {
			input:           "-0.123",
			isFloatingPoint: true,
		},
		"-.123 is not float": {
			input:           "-.123",
			isFloatingPoint: false,
		},
		"+0.123 is float": {
			input:           "+0.123",
			isFloatingPoint: true,
		},
		"+11.123 is float": {
			input:           "+11.123",
			isFloatingPoint: true,
		},
	}

	for testName, input := range numericTests {
		t.Run(testName, func(t *testing.T) {
			token := tokenFrom(input.input)
			if input.isFloatingPoint {
				if token.TokenType != FloatingPoint {
					t.Fatalf("Expected %v to be floating point but was %v", input.input, token.TokenType)
				}
			}
		})
	}
}

func TestTokenTypeAsBoolean(t *testing.T) {
	numericTests := map[string]struct {
		input     string
		isBoolean bool
	}{
		"true is boolean": {
			input:     "true",
			isBoolean: true,
		},
		"false is boolean": {
			input:     "false",
			isBoolean: true,
		},
		"y is boolean": {
			input:     "y",
			isBoolean: true,
		},
		"n is boolean": {
			input:     "n",
			isBoolean: true,
		},
		"test is not boolean": {
			input:     "test",
			isBoolean: false,
		},
		"no is not boolean": {
			input:     "no",
			isBoolean: false,
		},
	}

	for testName, input := range numericTests {
		t.Run(testName, func(t *testing.T) {
			token := tokenFrom(input.input)
			if input.isBoolean {
				if token.TokenType != Boolean {
					t.Fatalf("Expected %v to be boolean but was %v", input.input, token.TokenType)
				}
			}
		})
	}
}

func TestTokenTypeIsNumeric(t *testing.T) {
	numericTests := map[string]struct {
		token     Token
		isNumeric bool
	}{
		"12 is numeric": {
			token:     NewToken(Numeric, "12"),
			isNumeric: true,
		},
		"-12 is numeric": {
			token:     NewToken(Numeric, "-12"),
			isNumeric: true,
		},
		"+12 is numeric": {
			token:     NewToken(Numeric, "+12"),
			isNumeric: true,
		},
		"999999999 is numeric": {
			token:     NewToken(Numeric, "999999999"),
			isNumeric: true,
		},
		"test is not numeric": {
			token:     NewToken(RawString, "test"),
			isNumeric: false,
		},
	}

	for testName, input := range numericTests {
		t.Run(testName, func(t *testing.T) {
			if input.token.isNumeric() != input.isNumeric {
				t.Fatalf("Expected isBoolean for input %v to be %v", input.token.TokenValue, input.isNumeric)
			}
		})
	}
}

func TestTokenTypeIsFloatingPoint(t *testing.T) {
	numericTests := map[string]struct {
		token           Token
		isFloatingPoint bool
	}{
		"12.12 is float": {
			token:           NewToken(FloatingPoint, "12.12"),
			isFloatingPoint: true,
		},
		"-12.90 is float": {
			token:           NewToken(FloatingPoint, "-12.90"),
			isFloatingPoint: true,
		},
		"+12.89 is float": {
			token:           NewToken(FloatingPoint, "+12.89"),
			isFloatingPoint: true,
		},
		"0.12 is float": {
			token:           NewToken(FloatingPoint, "0.12"),
			isFloatingPoint: true,
		},
		"test is not float": {
			token:           NewToken(RawString, "test"),
			isFloatingPoint: false,
		},
	}

	for testName, input := range numericTests {
		t.Run(testName, func(t *testing.T) {
			if input.token.isFloatingPoint() != input.isFloatingPoint {
				t.Fatalf("Expected isBoolean for input %v to be %v", input.token.TokenValue, input.isFloatingPoint)
			}
		})
	}
}

func TestTokenTypeIsBoolean(t *testing.T) {
	numericTests := map[string]struct {
		token     Token
		isBoolean bool
	}{
		"true is boolean": {
			token:     NewToken(Boolean, "true"),
			isBoolean: true,
		},
		"false is boolean": {
			token:     NewToken(Boolean, "false"),
			isBoolean: true,
		},
		"y is boolean": {
			token:     NewToken(Boolean, "y"),
			isBoolean: true,
		},
		"n is boolean": {
			token:     NewToken(Boolean, "n"),
			isBoolean: true,
		},
		"test is not boolean": {
			token:     NewToken(RawString, "test"),
			isBoolean: false,
		},
	}

	for testName, input := range numericTests {
		t.Run(testName, func(t *testing.T) {
			if input.token.isBoolean() != input.isBoolean {
				t.Fatalf("Expected isBoolean for input %v to be %v", input.token.TokenValue, input.isBoolean)
			}
		})
	}
}
