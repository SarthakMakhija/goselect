package parser

func isATokenSeparator(ch rune) bool {
	if ch == ' ' || ch == ',' {
		return true
	}
	return false
}
