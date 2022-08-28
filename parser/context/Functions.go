package context

import (
	b64 "encoding/base64"
	"strings"
	"time"
)

type AllFunctions struct {
	supportedFunctions map[string]bool
}

func NewFunctions() *AllFunctions {
	return &AllFunctions{
		supportedFunctions: map[string]bool{
			"lower":     true,
			"low":       true,
			"upper":     true,
			"up":        true,
			"title":     true,
			"base64":    true,
			"b64":       true,
			"length":    true,
			"len":       true,
			"trim":      true,
			"ltrim":     true,
			"lTrim":     true,
			"rtrim":     true,
			"rTrim":     true,
			"now":       true,
			"date":      true,
			"day":       true,
			"month":     true,
			"mon":       true,
			"year":      true,
			"yr":        true,
			"dayOfWeek": true,
			"dayofweek": true,
		},
	}
}

func (functions *AllFunctions) IsASupportedFunction(function string) bool {
	return functions.supportedFunctions[strings.ToLower(function)]
}

func (functions *AllFunctions) Execute(fn string, args ...interface{}) interface{} {
	switch strings.ToLower(fn) {
	case "lower", "low":
		return strings.ToLower(args[0].(string))
	case "upper", "up":
		return strings.ToUpper(args[0].(string))
	case "length", "len":
		return len(args[0].(string))
	case "title":
		return strings.Title(args[0].(string))
	case "trim":
		return strings.TrimSpace(args[0].(string))
	case "ltrim", "lTrim":
		return strings.TrimLeft(args[0].(string), " ")
	case "rtrim", "rTrim":
		return strings.TrimRight(args[0].(string), " ")
	case "base64", "b64":
		d := []byte(args[0].(string))
		return b64.StdEncoding.EncodeToString(d)
	case "now":
		return now().String()
	case "day":
		return now().Day()
	case "month", "mon":
		return now().Month().String()
	case "year", "yr":
		return now().Year()
	case "dayOfWeek", "dayofweek":
		return now().Weekday().String()
	}
	return ""
}

var nowFunc = func() time.Time {
	return time.Now()
}

func now() time.Time {
	return nowFunc().UTC()
}

func resetClock() {
	nowFunc = func() time.Time {
		return time.Now()
	}
}
