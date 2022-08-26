package parser

var supportedFunctions = map[string]bool{
	"lower": true,
	"upper": true,
	"min":   true,
	"max":   true,
	"avg":   true,
	"sum":   true,
}

func isASupportedFunction(function string) bool {
	return supportedFunctions[function]
}
