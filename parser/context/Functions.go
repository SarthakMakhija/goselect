package context

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"os"
	"strings"
	"time"
)

type AllFunctions struct {
	supportedFunctions map[string]bool
}

type FunctionDefinition struct {
	aliases []string
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
	FunctionNameCurrentDate         = "date"
	FunctionNameCurrentDay          = "day"
	FunctionNameCurrentMonth        = "month"
	FunctionNameCurrentYear         = "year"
	FunctionNameDayOfWeek           = "dayofweek"
	FunctionNameWorkingDirectory    = "cwd"
	FunctionNameConcat              = "concat"
	FunctionNameConcatWithSeparator = "concatws"
	FunctionNameContains            = "contains"
)

var functionDefinitions = map[string]*FunctionDefinition{
	FunctionNameLower:               {aliases: []string{"lower", "low"}},
	FunctionNameUpper:               {aliases: []string{"upper", "up"}},
	FunctionNameTitle:               {aliases: []string{"title"}},
	FunctionNameBase64:              {aliases: []string{"base64", "b64"}},
	FunctionNameLength:              {aliases: []string{"length", "len"}},
	FunctionNameTrim:                {aliases: []string{"trim"}},
	FunctionNameLeftTrim:            {aliases: []string{"ltrim", "lefttrim"}},
	FunctionNameRightTrim:           {aliases: []string{"rtrim", "righttrim"}},
	FunctionNameNow:                 {aliases: []string{"now"}},
	FunctionNameCurrentDate:         {aliases: []string{"date"}},
	FunctionNameCurrentDay:          {aliases: []string{"day"}},
	FunctionNameCurrentMonth:        {aliases: []string{"month", "mon"}},
	FunctionNameCurrentYear:         {aliases: []string{"year", "yr"}},
	FunctionNameDayOfWeek:           {aliases: []string{"dayofweek", "dow"}},
	FunctionNameWorkingDirectory:    {aliases: []string{"cwd", "wd"}},
	FunctionNameConcat:              {aliases: []string{"concat"}},
	FunctionNameConcatWithSeparator: {aliases: []string{"concatws", "concatwithseparator"}},
	FunctionNameContains:            {aliases: []string{"contains"}},
}

func NewFunctions() *AllFunctions {
	supportedFunctions := make(map[string]bool)
	for _, functionDefinition := range functionDefinitions {
		for _, alias := range functionDefinition.aliases {
			supportedFunctions[alias] = true
		}
	}
	return &AllFunctions{
		supportedFunctions: supportedFunctions,
	}
}

func (functions *AllFunctions) IsASupportedFunction(function string) bool {
	return functions.supportedFunctions[strings.ToLower(function)]
}

func (functions *AllFunctions) Execute(fn string, args ...Value) (Value, error) {
	switch strings.ToLower(fn) {
	case "lower", "low":
		if err := functions.ensureNParametersOrError(args, fn, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.ToLower(args[0].stringValue)), nil
	case "upper", "up":
		if err := functions.ensureNParametersOrError(args, fn, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.ToUpper(args[0].stringValue)), nil
	case "length", "len":
		if err := functions.ensureNParametersOrError(args, fn, 1); err != nil {
			return EmptyValue(), err
		}
		return IntValue(len(args[0].stringValue)), nil
	case "title":
		if err := functions.ensureNParametersOrError(args, fn, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.Title(args[0].stringValue)), nil
	case "trim":
		if err := functions.ensureNParametersOrError(args, fn, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.TrimSpace(args[0].stringValue)), nil
	case "ltrim", "lTrim":
		if err := functions.ensureNParametersOrError(args, fn, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.TrimLeft(args[0].stringValue, " ")), nil
	case "rtrim", "rTrim":
		if err := functions.ensureNParametersOrError(args, fn, 1); err != nil {
			return EmptyValue(), err
		}
		return StringValue(strings.TrimRight(args[0].stringValue, " ")), nil
	case "base64", "b64":
		if err := functions.ensureNParametersOrError(args, fn, 1); err != nil {
			return EmptyValue(), err
		}
		d := []byte(args[0].stringValue)
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
	case "wd", "cwd":
		if dir, err := os.Getwd(); err != nil {
			return EmptyValue(), err
		} else {
			return StringValue(dir), nil
		}
	case "concat":
		var values []string
		for _, value := range args {
			values = append(values, value.stringValue)
		}
		return StringValue(strings.Join(values, "")), nil
	case "concatws", "concatWs":
		var values []string
		for index := 0; index < len(args)-1; index++ {
			values = append(values, args[index].stringValue)
		}
		return StringValue(strings.Join(values, args[len(args)-1].stringValue)), nil
	case "contains":
		if err := functions.ensureNParametersOrError(args, fn, 2); err != nil {
			return EmptyValue(), err
		}
		return BooleanValue(strings.Contains(args[0].stringValue, args[1].stringValue)), nil
	}
	return EmptyValue(), nil
}

func (functions *AllFunctions) ensureNParametersOrError(parameters []Value, fn string, n int) error {
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
