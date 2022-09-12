package context

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"goselect/parser/error/messages"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type IdentityFunctionBlock struct{}
type AddFunctionBlock struct{}
type SubtractFunctionBlock struct{}
type MultipleFunctionBlock struct{}
type DivideFunctionBlock struct{}
type EqualFunctionBlock struct{}
type NotEqualFunctionBlock struct{}
type LessThanFunctionBlock struct{}
type GreaterThanFunctionBlock struct{}
type LessThanEqualFunctionBlock struct{}
type GreaterThanEqualFunctionBlock struct{}
type OrFunctionBlock struct{}
type AndFunctionBlock struct{}
type NotFunctionBlock struct{}
type LikeFunctionBlock struct{}
type LowerFunctionBlock struct{}
type UpperFunctionBlock struct{}
type TitleFunctionBlock struct{ caser cases.Caser }
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
type ExtractFunctionBlock struct{}
type HoursDifferenceFunctionBlock struct{}
type DaysDifferenceFunctionBlock struct{}
type ParseDateTimeFunctionBlock struct{}
type WorkingDirectoryFunctionBlock struct{}
type ConcatFunctionBlock struct{}
type ConcatWithSeparatorFunctionBlock struct{}
type ContainsFunctionBlock struct{}
type SubstringFunctionBlock struct{}
type ReplaceFunctionBlock struct{}
type ReplaceAllFunctionBlock struct{}
type IsFileTypeTextFunctionBlock struct{}
type IsFileTypeImageFunctionBlock struct{}
type IsFileTypeAudioFunctionBlock struct{}
type IsFileTypeVideoFunctionBlock struct{}
type IsFileTypePdfFunctionBlock struct{}

func (receiver IdentityFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameIdentity, 1); err != nil {
		return EmptyValue, err
	}
	return args[0], nil
}

func (a AddFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameAdd, 2); err != nil {
		return EmptyValue, err
	}
	var result float64 = 0
	for _, arg := range args {
		asFloat64, err := arg.GetNumericAsFloat64()
		if err != nil {
			return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameAdd, err)
		}
		result = result + asFloat64
	}
	return Float64Value(result), nil
}

func (s SubtractFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameSubtract, 2); err != nil {
		return EmptyValue, err
	}
	oneFloat64, err := args[0].GetNumericAsFloat64()
	if err != nil {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameSubtract, err)
	}
	otherFloat64, err := args[1].GetNumericAsFloat64()
	if err != nil {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameSubtract, err)
	}
	return Float64Value(oneFloat64 - otherFloat64), nil
}

func (m MultipleFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameMultiply, 2); err != nil {
		return EmptyValue, err
	}
	var result float64 = 1
	for _, arg := range args {
		asFloat64, err := arg.GetNumericAsFloat64()
		if err != nil {
			return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameMultiply, err)
		}
		result = result * asFloat64
	}
	return Float64Value(result), nil
}

func (d DivideFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameDivide, 2); err != nil {
		return EmptyValue, err
	}
	oneFloat64, err := args[0].GetNumericAsFloat64()
	if err != nil {
		return EmptyValue, err
	}
	otherFloat64, err := args[1].GetNumericAsFloat64()
	if err != nil {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameDivide, err)
	}
	if otherFloat64 == float64(0) {
		return EmptyValue, errors.New(messages.ErrorMessageExpectedNonZeroInDivide)
	}
	return Float64Value(oneFloat64 / otherFloat64), nil
}

func (e EqualFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameEqual, 2); err != nil {
		return EmptyValue, err
	}
	if args[0].CompareTo(args[1]) == CompareToEqual {
		return trueBooleanValue, nil
	}
	return falseBooleanValue, nil
}

func (n NotEqualFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameEqual, 2); err != nil {
		return EmptyValue, err
	}
	if args[0].CompareTo(args[1]) == CompareToEqual {
		return falseBooleanValue, nil
	}
	return trueBooleanValue, nil
}

func (l LessThanFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLessThan, 2); err != nil {
		return EmptyValue, err
	}
	if args[0].CompareTo(args[1]) == CompareToLessThan {
		return trueBooleanValue, nil
	}
	return falseBooleanValue, nil
}

func (g GreaterThanFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLessThan, 2); err != nil {
		return EmptyValue, err
	}
	if args[0].CompareTo(args[1]) == CompareToGreaterThan {
		return trueBooleanValue, nil
	}
	return falseBooleanValue, nil
}

func (l LessThanEqualFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLessThanEqual, 2); err != nil {
		return EmptyValue, err
	}
	if args[0].CompareTo(args[1]) == CompareToLessThan || args[0].CompareTo(args[1]) == CompareToEqual {
		return trueBooleanValue, nil
	}
	return falseBooleanValue, nil
}

func (g GreaterThanEqualFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameGreaterThanEqual, 2); err != nil {
		return EmptyValue, err
	}
	if args[0].CompareTo(args[1]) == CompareToGreaterThan || args[0].CompareTo(args[1]) == CompareToEqual {
		return trueBooleanValue, nil
	}
	return falseBooleanValue, nil
}

func (o OrFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameOr, 1); err != nil {
		return EmptyValue, err
	}
	for _, arg := range args {
		result, err := arg.GetBoolean()
		if err != nil {
			return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameOr, err)
		}
		if result {
			return trueBooleanValue, nil
		}
	}
	return falseBooleanValue, nil
}

func (a AndFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameAnd, 1); err != nil {
		return EmptyValue, err
	}
	for _, arg := range args {
		result, err := arg.GetBoolean()
		if err != nil {
			return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameAnd, err)
		}
		if !result {
			return falseBooleanValue, nil
		}
	}
	return trueBooleanValue, nil
}

func (n NotFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameNot, 1); err != nil {
		return EmptyValue, err
	}
	result, err := args[0].GetBoolean()
	if err != nil {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameNot, err)
	}
	return booleanValueUsing(!result), nil
}

func (l LikeFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLike, 2); err != nil {
		return EmptyValue, err
	}
	toMatch := args[0].GetAsString()
	if compiled, err := regexp.Compile(args[1].GetAsString()); err != nil {
		return EmptyValue, err
	} else {
		return booleanValueUsing(compiled.MatchString(toMatch)), nil
	}
}

func (l LowerFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLower, 1); err != nil {
		return EmptyValue, err
	}
	return StringValue(strings.ToLower(args[0].GetAsString())), nil
}

func (u UpperFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameUpper, 1); err != nil {
		return EmptyValue, err
	}
	return StringValue(strings.ToUpper(args[0].GetAsString())), nil
}

func (t TitleFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameTitle, 1); err != nil {
		return EmptyValue, err
	}
	return StringValue(t.caser.String(args[0].GetAsString())), nil
}

func (b Base64FunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameBase64, 1); err != nil {
		return EmptyValue, err
	}
	d := []byte(args[0].GetAsString())
	return StringValue(b64.StdEncoding.EncodeToString(d)), nil
}

func (l LengthFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLength, 1); err != nil {
		return EmptyValue, err
	}
	return IntValue(len(args[0].GetAsString())), nil
}

func (t TrimFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameTrim, 1); err != nil {
		return EmptyValue, err
	}
	return StringValue(strings.TrimSpace(args[0].GetAsString())), nil
}

func (l LeftTrimFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameLeftTrim, 1); err != nil {
		return EmptyValue, err
	}
	return StringValue(strings.TrimLeft(args[0].GetAsString(), " ")), nil
}

func (r RightTrimFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameRightTrim, 1); err != nil {
		return EmptyValue, err
	}
	return StringValue(strings.TrimRight(args[0].GetAsString(), " ")), nil
}

func (n NowFunctionBlock) run(_ ...Value) (Value, error) {
	return DateTimeValue(now()), nil
}

func (c CurrentDayFunctionBlock) run(_ ...Value) (Value, error) {
	return IntValue(now().Day()), nil
}

func (c CurrentDateFunctionBlock) run(_ ...Value) (Value, error) {
	return formatDate(now()), nil
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

func (w WorkingDirectoryFunctionBlock) run(_ ...Value) (Value, error) {
	if dir, err := os.Getwd(); err != nil {
		return EmptyValue, err
	} else {
		return StringValue(dir), nil
	}
}

func (c ConcatFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameConcat, 2); err != nil {
		return EmptyValue, err
	}
	var values []string
	for _, value := range args {
		values = append(values, value.GetAsString())
	}
	return StringValue(strings.Join(values, "")), nil
}

func (c ConcatWithSeparatorFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameConcatWithSeparator, 3); err != nil {
		return EmptyValue, err
	}

	var values []string
	for index := 0; index < len(args)-1; index++ {
		values = append(values, args[index].GetAsString())
	}
	return StringValue(strings.Join(values, args[len(args)-1].GetAsString())), nil
}

func (c ContainsFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameContains, 2); err != nil {
		return EmptyValue, err
	}
	return booleanValueUsing(strings.Contains(args[0].stringValue, args[1].GetAsString())), nil
}

func (s SubstringFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameSubstring, 2); err != nil {
		return EmptyValue, err
	}
	str := args[0].GetAsString()
	length := len(str)
	from, err := strconv.Atoi(args[1].GetAsString())
	if err != nil {
		return EmptyValue, fmt.Errorf(
			messages.ErrorMessageFunctionNamePrefixWithExistingError,
			FunctionNameSubstring,
			messages.ErrorMessageIllegalFromToIndexInSubstring,
		)
	}
	if from < 0 {
		return EmptyValue, fmt.Errorf(
			messages.ErrorMessageFunctionNamePrefixWithExistingError,
			FunctionNameSubstring,
			messages.ErrorMessageIllegalFromToIndexInSubstring,
		)
	}
	if from >= length {
		from = 0
	}
	to := length - 1
	if len(args) >= 3 {
		to, err = strconv.Atoi(args[2].GetAsString())
		if err != nil {
			return EmptyValue, fmt.Errorf(
				messages.ErrorMessageFunctionNamePrefixWithExistingError,
				FunctionNameSubstring,
				messages.ErrorMessageIllegalFromToIndexInSubstring,
			)
		}
		if to < 0 {
			return EmptyValue, fmt.Errorf(
				messages.ErrorMessageFunctionNamePrefixWithExistingError,
				FunctionNameSubstring,
				messages.ErrorMessageIllegalFromToIndexInSubstring,
			)
		}
		if to < from {
			return EmptyValue, fmt.Errorf(
				messages.ErrorMessageFunctionNamePrefixWithExistingError,
				FunctionNameSubstring,
				messages.ErrorMessageIncorrectEndIndexInSubstring,
			)
		}
		if to >= length {
			to = length - 1
		}
	}
	return StringValue(str[from : to+1]), nil
}

func (r ReplaceFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameReplace, 3); err != nil {
		return EmptyValue, err
	}
	return StringValue(strings.Replace(args[0].GetAsString(), args[1].GetAsString(), args[2].GetAsString(), 1)), nil
}

func (r ReplaceAllFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameReplaceAll, 3); err != nil {
		return EmptyValue, err
	}
	return StringValue(strings.ReplaceAll(args[0].GetAsString(), args[1].GetAsString(), args[2].GetAsString())), nil
}

func (i IsFileTypeTextFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameIsFileTypeText, 1); err != nil {
		return EmptyValue, err
	}
	return booleanValueUsing(mimeTypeMatches("text/plain", args[0])), nil
}

func (i IsFileTypeImageFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameIsFileTypeImage, 1); err != nil {
		return EmptyValue, err
	}
	return booleanValueUsing(mimeTypeMatches("image/", args[0])), nil
}

func (i IsFileTypeAudioFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameIsFileTypeAudio, 1); err != nil {
		return EmptyValue, err
	}
	return booleanValueUsing(mimeTypeMatches("audio/", args[0])), nil
}

func (i IsFileTypeVideoFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameIsFileTypeVideo, 1); err != nil {
		return EmptyValue, err
	}
	return booleanValueUsing(mimeTypeMatches("video/", args[0])), nil
}

func (i IsFileTypePdfFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameIsFileTypePdf, 1); err != nil {
		return EmptyValue, err
	}
	return booleanValueUsing(
		mimeTypeMatches("application/pdf", args[0]) || mimeTypeMatches("application/x-pdf", args[0]),
	), nil
}

func mimeTypeMatches(expectedMimeType string, arg Value) bool {
	mimeType := arg.GetAsString()
	return strings.Contains(mimeType, expectedMimeType)
}

func (e ExtractFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameExtract, 2); err != nil {
		return EmptyValue, err
	}
	aTime, err := args[0].GetDateTime()
	if err != nil {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameExtract, err)
	}
	extractionKey := args[1].GetAsString()
	switch strings.ToLower(extractionKey) {
	case "date":
		return formatDate(aTime), nil
	case "day":
		return IntValue(aTime.Day()), nil
	case "year":
		return IntValue(aTime.Year()), nil
	case "month":
		return StringValue(aTime.Month().String()), nil
	case "weekday":
		return StringValue(aTime.Weekday().String()), nil
	default:
		return EmptyValue, fmt.Errorf(messages.ErrorMessageIncorrectExtractionKey, "date, day, year, month, weekday")
	}
}

func (h HoursDifferenceFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameHoursDifference, 1); err != nil {
		return EmptyValue, err
	}
	aTime, err := args[0].GetDateTime()
	if err != nil {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameHoursDifference, err)
	}
	bTime := now()
	if len(args) > 1 {
		bTime, err = args[1].GetDateTime()
		if err != nil {
			return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameHoursDifference, err)
		}
	}
	duration := bTime.Sub(aTime) //bTime - aTime
	return Float64Value(duration.Hours()), nil
}

func (d DaysDifferenceFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameDaysDifference, 1); err != nil {
		return EmptyValue, err
	}
	aTime, err := args[0].GetDateTime()
	if err != nil {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameDaysDifference, err)
	}
	bTime := now()
	if len(args) > 1 {
		bTime, err = args[1].GetDateTime()
		if err != nil {
			return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameDaysDifference, err)
		}
	}
	duration := bTime.Sub(aTime) //bTime - aTime
	days := duration.Hours() / float64(24)
	return Float64Value(days), nil
}

func (p ParseDateTimeFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameDateTimeParse, 2); err != nil {
		return EmptyValue, err
	}

	timeAsStr := args[0].GetAsString()
	formatId := args[1].GetAsString()
	parsed, err := parse(timeAsStr, formatId)
	if err != nil {
		return EmptyValue, err
	}
	return DateTimeValue(parsed), nil
}

func formatDate(time time.Time) Value {
	return StringValue(strconv.Itoa(time.Year()) + "-" + time.Month().String() + "-" + fmt.Sprintf("%02v", time.Day()))
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
		return fmt.Errorf(messages.ErrorMessageMissingParameterInScalarFunctions, n, fn)
	}
	return nil
}
