package context

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"goselect/parser/error/messages"
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

func (functions *AllFunctions) Execute(fn string, args ...interface{}) (interface{}, error) {
	switch strings.ToLower(fn) {
	case "lower", "low":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return nil, err
		}
		return strings.ToLower(args[0].(string)), nil
	case "upper", "up":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return nil, err
		}
		return strings.ToUpper(args[0].(string)), nil
	case "length", "len":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return nil, err
		}
		return len(args[0].(string)), nil
	case "title":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return nil, err
		}
		return strings.Title(args[0].(string)), nil
	case "trim":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return nil, err
		}
		return strings.TrimSpace(args[0].(string)), nil
	case "ltrim", "lTrim":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return nil, err
		}
		return strings.TrimLeft(args[0].(string), " "), nil
	case "rtrim", "rTrim":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return nil, err
		}
		return strings.TrimRight(args[0].(string), " "), nil
	case "base64", "b64":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return nil, err
		}
		d := []byte(args[0].(string))
		return b64.StdEncoding.EncodeToString(d), nil
	case "now":
		return now().String(), nil
	case "day":
		return now().Day(), nil
	case "month", "mon":
		return now().Month().String(), nil
	case "year", "yr":
		return now().Year(), nil
	case "dayOfWeek", "dayofweek":
		return now().Weekday().String(), nil
	}
	return "", nil
}

func (functions *AllFunctions) ensureOneParameterOrError(parameters []interface{}, fn string) error {
	nonNilParameterCount := func() int {
		count := 0
		for _, parameter := range parameters {
			if parameter != nil {
				count = count + 1
			}
		}
		return count
	}
	if len(parameters) < 1 || nonNilParameterCount() < 1 {
		return errors.New(fmt.Sprintf(messages.ErrorMessageMissingParameterInScalarFunctions, fn))
	}
	return nil
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
