package context

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"os"
	"strconv"
	"strings"
)

type LowerFunctionBlock struct{}
type UpperFunctionBlock struct{}
type TitleFunctionBlock struct{}
type Base64FunctionBlock struct{}
type LengthFunctionBlock struct{}
type TrimFunctionBlock struct{}
type LeftTrimFunctionBlock struct{}
type RightTrimFunctionBlock struct{}
type NowFunctionBlock struct{}
type CurrentDayFunctionBlock struct{}
type CurrentDateFunctionBlock struct{}
type CurrentMonthFunctionBlock struct{}
type CurrentYearFunctionBlock struct{}
type DayOfWeekFunctionBlock struct{}
type WorkingDirectoryFunctionBlock struct{}
type ConcatFunctionBlock struct{}
type ConcatWithSeparatorFunctionBlock struct{}
type ContainsFunctionBlock struct{}
type SubstringFunctionBlock struct{}

func (l LowerFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLower, 1); err != nil {
		return EmptyValue(), err
	}
	return StringValue(strings.ToLower(args[0].stringValue)), nil
}

func (u UpperFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameUpper, 1); err != nil {
		return EmptyValue(), err
	}
	return StringValue(strings.ToUpper(args[0].stringValue)), nil
}

func (t TitleFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameTitle, 1); err != nil {
		return EmptyValue(), err
	}
	return StringValue(strings.Title(args[0].stringValue)), nil
}

func (b Base64FunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameBase64, 1); err != nil {
		return EmptyValue(), err
	}
	d := []byte(args[0].stringValue)
	return StringValue(b64.StdEncoding.EncodeToString(d)), nil
}

func (l LengthFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLength, 1); err != nil {
		return EmptyValue(), err
	}
	return IntValue(len(args[0].stringValue)), nil
}

func (t TrimFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameTrim, 1); err != nil {
		return EmptyValue(), err
	}
	return StringValue(strings.TrimSpace(args[0].stringValue)), nil
}

func (l LeftTrimFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLeftTrim, 1); err != nil {
		return EmptyValue(), err
	}
	return StringValue(strings.TrimLeft(args[0].stringValue, " ")), nil
}

func (r RightTrimFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameRightTrim, 1); err != nil {
		return EmptyValue(), err
	}
	return StringValue(strings.TrimRight(args[0].stringValue, " ")), nil
}

func (n NowFunctionBlock) run(_ ...Value) (Value, error) {
	return DateTimeValue(now()), nil
}

func (c CurrentDayFunctionBlock) run(_ ...Value) (Value, error) {
	return IntValue(now().Day()), nil
}

func (c CurrentDateFunctionBlock) run(_ ...Value) (Value, error) {
	year, month, day := now().Date()
	return StringValue(strconv.Itoa(year) + "-" + month.String() + "-" + strconv.Itoa(day)), nil
}

func (c CurrentMonthFunctionBlock) run(_ ...Value) (Value, error) {
	return StringValue(now().Month().String()), nil
}

func (c CurrentYearFunctionBlock) run(_ ...Value) (Value, error) {
	return IntValue(now().Year()), nil
}

func (d DayOfWeekFunctionBlock) run(_ ...Value) (Value, error) {
	return StringValue(now().Weekday().String()), nil
}

func (w WorkingDirectoryFunctionBlock) run(args ...Value) (Value, error) {
	if dir, err := os.Getwd(); err != nil {
		return EmptyValue(), err
	} else {
		return StringValue(dir), nil
	}
}

func (c ConcatFunctionBlock) run(args ...Value) (Value, error) {
	var values []string
	for _, value := range args {
		values = append(values, value.stringValue)
	}
	return StringValue(strings.Join(values, "")), nil
}

func (c ConcatWithSeparatorFunctionBlock) run(args ...Value) (Value, error) {
	var values []string
	for index := 0; index < len(args)-1; index++ {
		values = append(values, args[index].stringValue)
	}
	return StringValue(strings.Join(values, args[len(args)-1].stringValue)), nil
}

func (c ContainsFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameContains, 2); err != nil {
		return EmptyValue(), err
	}
	return BooleanValue(strings.Contains(args[0].stringValue, args[1].stringValue)), nil
}

func (s SubstringFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameSubstring, 2); err != nil {
		return EmptyValue(), err
	}
	str := args[0].stringValue
	length := len(str)
	from, err := strconv.Atoi(args[1].stringValue)
	if err != nil {
		return EmptyValue(), errors.New(messages.ErrorMessageIllegalFromToIndexInSubstring)
	}
	if from < 0 {
		return EmptyValue(), errors.New(messages.ErrorMessageIllegalFromToIndexInSubstring)
	}
	if from >= length {
		from = 0
	}
	to := length - 1
	if len(args) >= 3 {
		to, err = strconv.Atoi(args[2].stringValue)
		if err != nil {
			return EmptyValue(), errors.New(messages.ErrorMessageIllegalFromToIndexInSubstring)
		}
		if to < 0 {
			return EmptyValue(), errors.New(messages.ErrorMessageIllegalFromToIndexInSubstring)
		}
		if to < from {
			return EmptyValue(), errors.New(messages.ErrorMessageIncorrectEndIndexInSubstring)
		}
		if to >= length {
			to = length - 1
		}
	}
	return StringValue(str[from : to+1]), nil
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
