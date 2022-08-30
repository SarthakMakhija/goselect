package context

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"os"
	"strconv"
	"strings"
	"time"
)

type Function struct {
	aliases []string
	block   func(args ...Value) (Value, error)
}

type AllFunctions struct {
	supportedFunctions map[string]*Function
}

const (
	FunctionNameLower               = "lower"
	FunctionNameUpper               = "upper"
	FunctionNameTitle               = "title"
	FunctionNameBase64              = "base64"
	FunctionNameLength              = "length"
	FunctionNameTrim                = "trim"
	FunctionNameLeftTrim            = "ltrim"
	FunctionNameRightTrim           = "rtrim"
	FunctionNameNow                 = "now"
	FunctionNameCurrentDay          = "day"
	FunctionNameCurrentDate         = "date"
	FunctionNameCurrentMonth        = "month"
	FunctionNameCurrentYear         = "year"
	FunctionNameDayOfWeek           = "dayofweek"
	FunctionNameWorkingDirectory    = "cwd"
	FunctionNameConcat              = "concat"
	FunctionNameConcatWithSeparator = "concatws"
	FunctionNameContains            = "contains"
)

var functionDefinitions = map[string]*Function{
	FunctionNameLower: {aliases: []string{"lower", "low"}, block: func(args ...Value) (Value, error) {
		if err := ensureNParametersOrError(args, FunctionNameLower, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.ToLower(args[0].stringValue)), nil
	}},
	FunctionNameUpper: {aliases: []string{"upper", "up"}, block: func(args ...Value) (Value, error) {
		if err := ensureNParametersOrError(args, FunctionNameUpper, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.ToUpper(args[0].stringValue)), nil
	}},
	FunctionNameTitle: {aliases: []string{"title"}, block: func(args ...Value) (Value, error) {
		if err := ensureNParametersOrError(args, FunctionNameTitle, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.Title(args[0].stringValue)), nil
	}},
	FunctionNameBase64: {aliases: []string{"base64", "b64"}, block: func(args ...Value) (Value, error) {
		if err := ensureNParametersOrError(args, FunctionNameBase64, 1); err != nil {
			return EmptyValue(), err
		}
		d := []byte(args[0].stringValue)
		return StringValue(b64.StdEncoding.EncodeToString(d)), nil
	}},
	FunctionNameLength: {aliases: []string{"length", "len"}, block: func(args ...Value) (Value, error) {
		if err := ensureNParametersOrError(args, FunctionNameLength, 1); err != nil {
			return EmptyValue(), err
		}
		return IntValue(len(args[0].stringValue)), nil
	}},
	FunctionNameTrim: {aliases: []string{"trim"}, block: func(args ...Value) (Value, error) {
		if err := ensureNParametersOrError(args, FunctionNameTrim, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.TrimSpace(args[0].stringValue)), nil
	}},
	FunctionNameLeftTrim: {aliases: []string{"ltrim", "lefttrim"}, block: func(args ...Value) (Value, error) {
		if err := ensureNParametersOrError(args, FunctionNameLeftTrim, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.TrimLeft(args[0].stringValue, " ")), nil
	}},
	FunctionNameRightTrim: {aliases: []string{"rtrim", "righttrim"}, block: func(args ...Value) (Value, error) {
		if err := ensureNParametersOrError(args, FunctionNameRightTrim, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.TrimRight(args[0].stringValue, " ")), nil
	}},
	FunctionNameNow: {aliases: []string{"now"}, block: func(args ...Value) (Value, error) {
		return DateTimeValue(now()), nil
	}},
	FunctionNameCurrentDay: {aliases: []string{"day"}, block: func(args ...Value) (Value, error) {
		return IntValue(now().Day()), nil
	}},
	FunctionNameCurrentDate: {aliases: []string{"date"}, block: func(args ...Value) (Value, error) {
		year, month, day := now().Date()
		return StringValue(strconv.Itoa(year) + "-" + month.String() + "-" + strconv.Itoa(day)), nil
	}},
	FunctionNameCurrentMonth: {aliases: []string{"month", "mon"}, block: func(args ...Value) (Value, error) {
		return StringValue(now().Month().String()), nil
	}},
	FunctionNameCurrentYear: {aliases: []string{"year", "yr"}, block: func(args ...Value) (Value, error) {
		return IntValue(now().Year()), nil
	}},
	FunctionNameDayOfWeek: {aliases: []string{"dayofweek", "dow"}, block: func(args ...Value) (Value, error) {
		return StringValue(now().Weekday().String()), nil
	}},
	FunctionNameWorkingDirectory: {aliases: []string{"cwd", "wd"}, block: func(args ...Value) (Value, error) {
		if dir, err := os.Getwd(); err != nil {
			return EmptyValue(), err
		} else {
			return StringValue(dir), nil
		}
	}},
	FunctionNameConcat: {aliases: []string{"concat"}, block: func(args ...Value) (Value, error) {
		var values []string
		for _, value := range args {
			values = append(values, value.stringValue)
		}
		return StringValue(strings.Join(values, "")), nil
	}},
	FunctionNameConcatWithSeparator: {aliases: []string{"concatws", "concatwithseparator"}, block: func(args ...Value) (Value, error) {
		var values []string
		for index := 0; index < len(args)-1; index++ {
			values = append(values, args[index].stringValue)
		}
		return StringValue(strings.Join(values, args[len(args)-1].stringValue)), nil
	}},
	FunctionNameContains: {aliases: []string{"contains"}, block: func(args ...Value) (Value, error) {
		if err := ensureNParametersOrError(args, FunctionNameContains, 2); err != nil {
			return EmptyValue(), err
		}
		return BooleanValue(strings.Contains(args[0].stringValue, args[1].stringValue)), nil
	}},
}

func NewFunctions() *AllFunctions {
	supportedFunctions := make(map[string]*Function)
	for _, functionDefinition := range functionDefinitions {
		for _, alias := range functionDefinition.aliases {
			supportedFunctions[alias] = functionDefinition
		}
	}
	return &AllFunctions{
		supportedFunctions: supportedFunctions,
	}
}

func (functions *AllFunctions) IsASupportedFunction(function string) bool {
	_, ok := functions.supportedFunctions[strings.ToLower(function)]
	return ok
}

func (functions *AllFunctions) Execute(fn string, args ...Value) (Value, error) {
	return functions.supportedFunctions[strings.ToLower(fn)].block(args...)
}

func ensureNParametersOrError(parameters []Value, fn string, n int) error {
	nonNilParameterCount := func() int {
		count := 0
		for _, parameter := range parameters {
			if parameter.valueType != ValueTypeUndefined {
				count = count + 1
			}
		}
		return count
	}
	if len(parameters) < n || nonNilParameterCount() < n {
		return errors.New(fmt.Sprintf(messages.ErrorMessageMissingParameterInScalarFunctions, n, fn))
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
