package projection

import (
	b64 "encoding/base64"
	"strings"
	"time"
)

var supportedFunctions = map[string]bool{
	"lower":  true,
	"upper":  true,
	"title":  true,
	"base64": true,
	"now":    true,
}

func isASupportedFunction(function string) bool {
	return supportedFunctions[function]
}

func ExecuteFn(fn string, args ...interface{}) interface{} {
	switch fn {
	case "lower":
		return strings.ToLower(args[0].(string))
	case "upper":
		return strings.ToUpper(args[0].(string))
	case "title":
		return strings.Title(args[0].(string))
	case "now":
		return time.Now().String()
	case "base64":
		d := []byte(args[0].(string))
		return b64.StdEncoding.EncodeToString(d)
	}
	return ""
}
