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

func (functions *AllFunctions) Execute(fn string, args ...Value) (Value, error) {
	switch strings.ToLower(fn) {
	case "lower", "low":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.ToLower(args[0].Get().(string))), nil
	case "upper", "up":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.ToUpper(args[0].Get().(string))), nil
	case "length", "len":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return EmptyValue(), err
		}
		return IntValue(len(args[0].Get().(string))), nil
	case "title":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.Title(args[0].Get().(string))), nil
	case "trim":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.TrimSpace(args[0].Get().(string))), nil
	case "ltrim", "lTrim":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.TrimLeft(args[0].Get().(string), " ")), nil
	case "rtrim", "rTrim":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.TrimRight(args[0].Get().(string), " ")), nil
	case "base64", "b64":
		if err := functions.ensureOneParameterOrError(args, fn); err != nil {
			return EmptyValue(), err
		}
		d := []byte(args[0].Get().(string))
		return StringValue(b64.StdEncoding.EncodeToString(d)), nil
	case "now":
		return DateTimeValue(now()), nil
	case "day":
		return IntValue(now().Day()), nil
	case "month", "mon":
		return StringValue(now().Month().String()), nil
	case "year", "yr":
		return IntValue(now().Year()), nil
	case "dayOfWeek", "dayofweek":
		return StringValue(now().Weekday().String()), nil
	}
	return EmptyValue(), nil
}

func (functions *AllFunctions) ensureOneParameterOrError(parameters []Value, fn string) error {
	nonNilParameterCount := func() int {
		count := 0
		for _, parameter := range parameters {
			if parameter.valueType != ValueTypeUndefined {
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