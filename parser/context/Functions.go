package context

import (
	"strings"
	"time"
)

type FunctionDefinition struct {
	aliases        []string
	tags           map[string]bool
	block          FunctionBlock
	aggregateBlock AggregationFunctionBlock
	isAggregate    bool
}

type FunctionBlock interface {
	run(args ...Value) (Value, error)
}

type AggregationFunctionBlock interface {
	initialState() *FunctionState
	run(initialState *FunctionState, args ...Value) (*FunctionState, error)
	finalValue(*FunctionState, []Value) (Value, error)
}

type AllFunctions struct {
	supportedFunctions map[string]*FunctionDefinition
}

type FunctionState struct {
	Initial   Value
	isUpdated bool
	extras    map[interface{}]Value
}

const (
	FunctionNameIdentity            = "identity"
	FunctionNameAdd                 = "add"
	FunctionNameSubtract            = "subtract"
	FunctionNameMultiply            = "multiply"
	FunctionNameDivide              = "divide"
	FunctionNameEqual               = "equal"
	FunctionNameNotEqual            = "notequal"
	FunctionNameLessThan            = "lessthan"
	FunctionNameGreaterThan         = "greaterthan"
	FunctionNameLessThanEqual       = "lessthanequal"
	FunctionNameGreaterThanEqual    = "greaterthanequal"
	FunctionNameOr                  = "or"
	FunctionNameAnd                 = "and"
	FunctionNameNot                 = "not"
	FunctionNameLike                = "like"
	FunctionNameLower               = "lower"
	FunctionNameUpper               = "upper"
	FunctionNameTitle               = "title"
	FunctionNameBase64              = "base64"
	FunctionNameLength              = "length"
	FunctionNameTrim                = "trim"
	FunctionNameLeftTrim            = "ltrim"
	FunctionNameRightTrim           = "rtrim"
	FunctionNameNow                 = "now"
	FunctionNameCurrentDay          = "cday"
	FunctionNameCurrentDate         = "cdate"
	FunctionNameCurrentMonth        = "cmonth"
	FunctionNameCurrentYear         = "cyear"
	FunctionNameDayOfWeek           = "dayofweek"
	FunctionNameExtract             = "extract"
	FunctionNameHoursDifference     = "hoursdifference"
	FunctionNameDaysDifference      = "daysdifference"
	FunctionNameWorkingDirectory    = "cwd"
	FunctionNameConcat              = "concat"
	FunctionNameConcatWithSeparator = "concatws"
	FunctionNameContains            = "contains"
	FunctionNameSubstring           = "substr"
	FunctionNameReplace             = "replace"
	FunctionNameReplaceAll          = "replaceall"
	FunctionNameCount               = "count"
	FunctionNameCountDistinct       = "countdistinct"
	FunctionNameSum                 = "sum"
	FunctionNameAverage             = "average"
	FunctionNameMin                 = "min"
)

var functionDefinitions = map[string]*FunctionDefinition{
	FunctionNameIdentity: {
		aliases: []string{"identity", "iden"},
		block:   IdentityFunctionBlock{},
	},
	FunctionNameAdd: {
		aliases: []string{"add", "addition"},
		block:   AddFunctionBlock{},
	},
	FunctionNameSubtract: {
		aliases: []string{"sub", "subtract"},
		block:   SubtractFunctionBlock{},
	},
	FunctionNameMultiply: {
		aliases: []string{"mul", "multiply"},
		block:   MultipleFunctionBlock{},
	},
	FunctionNameDivide: {
		aliases: []string{"div", "divide"},
		block:   DivideFunctionBlock{},
	},
	FunctionNameEqual: {
		aliases: []string{"equal", "eq", "equals"},
		block:   EqualFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
	FunctionNameNotEqual: {
		aliases: []string{"notequal", "ne", "notequals"},
		block:   NotEqualFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
	FunctionNameLessThan: {
		aliases: []string{"lt", "lessthan", "less"},
		block:   LessThanFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
	FunctionNameGreaterThan: {
		aliases: []string{"gt", "greater", "greaterthan"},
		block:   GreaterThanFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
	FunctionNameLessThanEqual: {
		aliases: []string{"lte", "lessthanequal", "lessequal", "le"},
		block:   LessThanEqualFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
	FunctionNameGreaterThanEqual: {
		aliases: []string{"gte", "greaterthanequal", "greaterequal", "ge"},
		block:   GreaterThanEqualFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
	FunctionNameOr: {
		aliases: []string{"or"},
		block:   OrFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
	FunctionNameAnd: {
		aliases: []string{"and"},
		block:   AndFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
	FunctionNameNot: {
		aliases: []string{"not"},
		block:   NotFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
	FunctionNameLike: {
		aliases: []string{"like"},
		block:   LikeFunctionBlock{},
		tags:    map[string]bool{"where": true},
	},
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
		aliases: []string{"cday", "currentday"},
		block:   CurrentDayFunctionBlock{},
	},
	FunctionNameCurrentDate: {
		aliases: []string{"cdate", "currentdate"},
		block:   CurrentDateFunctionBlock{},
	},
	FunctionNameCurrentMonth: {
		aliases: []string{"cmonth", "cmon", "currentmonth", "currentmon"},
		block:   CurrentMonthFunctionBlock{},
	},
	FunctionNameCurrentYear: {
		aliases: []string{"cyear", "cyr", "currentyear", "currentyr"},
		block:   CurrentYearFunctionBlock{},
	},
	FunctionNameDayOfWeek: {
		aliases: []string{"dayofweek", "dow"},
		block:   DayOfWeekFunctionBlock{},
	},
	FunctionNameExtract: {
		aliases: []string{"extract"},
		block:   ExtractFunctionBlock{},
	},
	FunctionNameHoursDifference: {
		aliases: []string{"hoursdifference", "hourdifference", "hoursdiff", "hourdiff"},
		block:   HoursDifferenceFunctionBlock{},
	},
	FunctionNameDaysDifference: {
		aliases: []string{"daysdifference", "daydifference", "daysdiff", "daydiff"},
		block:   DaysDifferenceFunctionBlock{},
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
		tags:    map[string]bool{"where": true},
	},
	FunctionNameSubstring: {
		aliases: []string{"substr", "str"},
		block:   SubstringFunctionBlock{},
	},
	FunctionNameReplace: {
		aliases: []string{"replace"},
		block:   ReplaceFunctionBlock{},
	},
	FunctionNameReplaceAll: {
		aliases: []string{"replaceall"},
		block:   ReplaceAllFunctionBlock{},
	},
	FunctionNameCount: {
		aliases:        []string{"count"},
		isAggregate:    true,
		aggregateBlock: &CountFunctionBlock{},
	},
	FunctionNameCountDistinct: {
		aliases:        []string{"countdistinct", "countd"},
		isAggregate:    true,
		aggregateBlock: &CountDistinctFunctionBlock{},
	},
	FunctionNameSum: {
		aliases:        []string{"summation", "sum"},
		isAggregate:    true,
		aggregateBlock: &SumFunctionBlock{},
	},
	FunctionNameAverage: {
		aliases:        []string{"average", "avg"},
		isAggregate:    true,
		aggregateBlock: &AverageFunctionBlock{},
	},
	FunctionNameMin: {
		aliases:        []string{"min"},
		isAggregate:    true,
		aggregateBlock: &MinFunctionBlock{},
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

func (functions *AllFunctions) ContainsATag(function string, tag string) bool {
	definition, ok := functions.supportedFunctions[strings.ToLower(function)]
	if !ok {
		return false
	}
	return definition.tags[strings.ToLower(tag)]
}

func (functions *AllFunctions) IsAnAggregateFunction(function string) bool {
	fn, ok := functions.supportedFunctions[strings.ToLower(function)]
	if ok {
		return fn.isAggregate
	}
	return false
}

func (functions *AllFunctions) AllFunctionsWithAliases() map[string][]string {
	aliasesByFunction := make(map[string][]string, len(functionDefinitions))
	for function, definition := range functionDefinitions {
		aliasesByFunction[function] = definition.aliases
	}
	return aliasesByFunction
}

func (functions *AllFunctions) AllFunctionsWithAliasesHavingTag(tag string) map[string][]string {
	aliasesByFunction := make(map[string][]string)
	for function, definition := range functionDefinitions {
		if definition.tags[tag] {
			aliasesByFunction[function] = definition.aliases
		}
	}
	return aliasesByFunction
}

func (functions *AllFunctions) Execute(fn string, args ...Value) (Value, error) {
	return functions.supportedFunctions[strings.ToLower(fn)].block.run(args...)
}

func (functions *AllFunctions) ExecuteAggregate(fn string, initialState *FunctionState, args ...Value) (*FunctionState, error) {
	return functions.supportedFunctions[strings.ToLower(fn)].aggregateBlock.run(initialState, args...)
}

func (functions *AllFunctions) InitialState(fn string) *FunctionState {
	function := functions.supportedFunctions[strings.ToLower(fn)]
	if function.isAggregate {
		return function.aggregateBlock.initialState()
	}
	return nil
}

func (functions *AllFunctions) FinalValue(fn string, state *FunctionState, values []Value) (Value, error) {
	function := functions.supportedFunctions[strings.ToLower(fn)]
	if function.isAggregate {
		return function.aggregateBlock.finalValue(state, values)
	}
	return EmptyValue, nil
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
