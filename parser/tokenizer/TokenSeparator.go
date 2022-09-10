package tokenizer

func isCharATokenSeparator(ch rune) bool {
	return ch == ' '
}

func isCharAComparisonOperator(ch rune) bool {
	if ch == '=' || ch == '!' || ch == '>' || ch == '<' {
		return true
	}
	return false
}

func isArithmeticOperator(token string) bool {
	if token == "+" || token == "-" || token == "*" || token == "/" || token == "%" {
		return true
	}
	return false
}

func isAComparisonOperator(token string) bool {
	if token == "=" || token == "!=" || token == ">" || token == "<" || token == ">=" ||
		token == "<=" || token == "like" || token == "regexp" {
		return true
	}
	return false
}
