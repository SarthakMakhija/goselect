package context

import (
	"errors"
	"goselect/parser/error/messages"
	"strings"
	"time"
)

type FunctionDefinition struct {
	aliases     []string
	block       FunctionBlock
	isAggregate bool
}

type FunctionBlock interface {
	run(args ...Value) (Value, error)
}

type AggregationFunctionBlock interface {
	finalState() Value
}

type AllFunctions struct {
	supportedFunctions map[string]*FunctionDefinition
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
	FunctionNameSubstring           = "substr"
	FunctionNameCount               = "count"
	FunctionNameCountDistinct       = "countdistinct"
)

var functionDefinitions = map[string]*FunctionDefinition{
	FunctionNameLower: {
		aliases: []string{"lower", "low"},
		block:   LowerFunctionBlock{},
	},
	FunctionNameUpper: {
		aliases: []string{"upper", "up"},
		block:   UpperFunctionBlock{},
	},
	FunctionNameTitle: {
		aliases: []string{"title"},
		block:   TitleFunctionBlock{},
	},
	FunctionNameBase64: {
		aliases: []string{"base64", "b64"},
		block:   Base64FunctionBlock{},
	},
	FunctionNameLength: {
		aliases: []string{"length", "len"},
		block:   LengthFunctionBlock{},
	},
	FunctionNameTrim: {
		aliases: []string{"trim"},
		block:   TrimFunctionBlock{},
	},
	FunctionNameLeftTrim: {
		aliases: []string{"ltrim", "lefttrim"},
		block:   LeftTrimFunctionBlock{},
	},
	FunctionNameRightTrim: {
		aliases: []string{"rtrim", "righttrim"},
		block:   RightTrimFunctionBlock{},
	},
	FunctionNameNow: {
		aliases: []string{"now"},
		block:   NowFunctionBlock{},
	},
	FunctionNameCurrentDay: {
		aliases: []string{"day"},
		block:   CurrentDayFunctionBlock{},
	},
	FunctionNameCurrentDate: {
		aliases: []string{"date"},
		block:   CurrentDateFunctionBlock{},
	},
	FunctionNameCurrentMonth: {
		aliases: []string{"month", "mon"},
		block:   CurrentMonthFunctionBlock{},
	},
	FunctionNameCurrentYear: {
		aliases: []string{"year", "yr"},
		block:   CurrentYearFunctionBlock{},
	},
	FunctionNameDayOfWeek: {
		aliases: []string{"dayofweek", "dow"},
		block:   DayOfWeekFunctionBlock{},
	},
	FunctionNameWorkingDirectory: {
		aliases: []string{"cwd", "wd"},
		block:   WorkingDirectoryFunctionBlock{},
	},
	FunctionNameConcat: {
		aliases: []string{"concat"},
		block:   ConcatFunctionBlock{},
	},
	FunctionNameConcatWithSeparator: {
		aliases: []string{"concatws", "concatwithseparator"},
		block:   ConcatWithSeparatorFunctionBlock{},
	},
	FunctionNameContains: {
		aliases: []string{"contains"},
		block:   ContainsFunctionBlock{},
	},
	FunctionNameSubstring: {
		aliases: []string{"substr", "str"},
		block:   SubstringFunctionBlock{},
	},
	FunctionNameCount: {
		aliases:     []string{"count"},
		isAggregate: true,
		block:       &CountFunctionBlock{},
	},
	FunctionNameCountDistinct: {
		aliases:     []string{"countdistinct"},
		isAggregate: true,
		block:       &CountDistinctFunctionBlock{values: make(map[Value]bool)},
	},
}

func NewFunctions() *AllFunctions {
	supportedFunctions := make(map[string]*FunctionDefinition)
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
	return functions.supportedFunctions[strings.ToLower(fn)].block.run(args...)
}

func (functions *AllFunctions) FinalState(fn string) (Value, error) {
	function := functions.supportedFunctions[strings.ToLower(fn)]
	if function.isAggregate {
		return function.block.(AggregationFunctionBlock).finalState(), nil
	}
	return EmptyValue(), errors.New(messages.ErrorMessageFinalStateCalledOnScalarFunction)
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
