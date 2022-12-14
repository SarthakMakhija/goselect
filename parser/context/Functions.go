package context

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"time"
)

type FunctionDefinition struct {
	aliases        []string
	tags           map[string]bool
	block          FunctionBlock
	description    string
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
	FunctionNameIfBlank             = "ifblank"
	FunctionNameStartsWith          = "startswith"
	FunctionNameEndsWith            = "endswith"
	FunctionNameNow                 = "now"
	FunctionNameCurrentDay          = "cday"
	FunctionNameCurrentDate         = "cdate"
	FunctionNameCurrentMonth        = "cmonth"
	FunctionNameCurrentYear         = "cyear"
	FunctionNameDayOfWeek           = "dayofweek"
	FunctionNameExtract             = "extract"
	FunctionNameHoursDifference     = "hoursdifference"
	FunctionNameDaysDifference      = "daysdifference"
	FunctionNameDateTimeParse       = "parsedatetime"
	FunctionNameWorkingDirectory    = "cwd"
	FunctionNameConcat              = "concat"
	FunctionNameConcatWithSeparator = "concatws"
	FunctionNameContains            = "contains"
	FunctionNameSubstring           = "substr"
	FunctionNameReplace             = "replace"
	FunctionNameReplaceAll          = "replaceall"
	FunctionNameIsFileTypeText      = "istext"
	FunctionNameIsFileTypeImage     = "isimage"
	FunctionNameIsFileTypeAudio     = "isaudio"
	FunctionNameIsFileTypeVideo     = "isvideo"
	FunctionNameIsFileTypePdf       = "ispdf"
	FunctionNameIsFileTypeArchive   = "isarchive"
	FunctionNameFormatSize          = "formatsize"
	FunctionNameParseSize           = "parsesize"
	FunctionNameCount               = "count"
	FunctionNameCountDistinct       = "countdistinct"
	FunctionNameSum                 = "sum"
	FunctionNameAverage             = "average"
	FunctionNameMin                 = "min"
	FunctionNameMax                 = "max"
)

var executionCache = NewFunctionExecutionCache()

var functionDefinitions = map[string]*FunctionDefinition{
	FunctionNameIdentity: {
		aliases:     []string{"identity", "iden"},
		description: "Returns the provided parameter value as it is, if the parameter value is not an attribute. \nFor example, identity(demo) will return the string demo, identity(name) will return the file name.",
		block:       IdentityFunctionBlock{},
	},
	FunctionNameAdd: {
		aliases:     []string{"add", "addition"},
		description: "Takes variable number of numeric type parameter values and returns the addition of all the values. \nFor example, add(1, 2) will return 3.00.",
		block:       AddFunctionBlock{},
	},
	FunctionNameSubtract: {
		aliases:     []string{"sub", "subtract"},
		description: "Takes 2 numeric type parameter values A and B and returns the result of A-B. \nFor example, sub(4, 5) will return -1.00.",
		block:       SubtractFunctionBlock{},
	},
	FunctionNameMultiply: {
		aliases:     []string{"mul", "multiply"},
		description: "Takes variable number of numeric type parameter values and returns the product of all the values. \nFor example, mul(3, 2) will return 6.00.",
		block:       MultiplyFunctionBlock{},
	},
	FunctionNameDivide: {
		aliases:     []string{"div", "divide"},
		description: "Takes 2 numeric type parameter values A and B and returns the result of A/B. \nFor example, div(4, 5) will return 0.80.",
		block:       DivideFunctionBlock{},
	},
	FunctionNameEqual: {
		aliases:     []string{"equal", "eq", "equals"},
		description: "Takes 2 parameter values A and B and returns true if A is equal to B, false otherwise.",
		block:       EqualFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameNotEqual: {
		aliases:     []string{"notequal", "ne", "notequals"},
		description: "Takes 2 parameter values A and B and returns true if A is not equal to B, false otherwise.",
		block:       NotEqualFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameLessThan: {
		aliases:     []string{"lt", "lessthan", "less"},
		description: "Takes 2 parameter values A and B and returns true if A is less than B, false otherwise.",
		block:       LessThanFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameGreaterThan: {
		aliases:     []string{"gt", "greater", "greaterthan"},
		description: "Takes 2 parameter values A and B and returns true if A is greater than B, false otherwise.",
		block:       GreaterThanFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameLessThanEqual: {
		aliases:     []string{"lte", "lessthanequal", "lessequal", "le"},
		description: "Takes 2 parameter values A and B and returns true if A is less than or equal to B, false otherwise.",
		block:       LessThanEqualFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameGreaterThanEqual: {
		aliases:     []string{"gte", "greaterthanequal", "greaterequal", "ge"},
		description: "Takes 2 parameter values A and B and returns true if A is greater than or equal to B, false otherwise.",
		block:       GreaterThanEqualFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameOr: {
		aliases:     []string{"or"},
		description: "Takes variable number of boolean parameter values and returns true if any of them evaluates to true, false otherwise. \nFor example, or(eq(add(1, 2), 3), false) will return true.",
		block:       OrFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameAnd: {
		aliases:     []string{"and"},
		description: "Takes variable number of boolean parameter values and returns true if all of them evaluate to true, false otherwise. \nFor example, or(eq(add(1, 2), 3), false) will return false.",
		block:       AndFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameNot: {
		aliases:     []string{"not"},
		description: "Takes a single boolean parameter value and returns its negation.",
		block:       NotFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameLike: {
		aliases:     []string{"like"},
		description: "Takes 2 parameter values and returns true if the first parameter value matches the regular expression represented by the second parameter value, false otherwise.",
		block:       LikeFunctionBlock{executionCache: executionCache},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameLower: {
		aliases:     []string{"lower", "low"},
		description: "Takes a single parameter value and returns the value in lower case.",
		block:       LowerFunctionBlock{},
	},
	FunctionNameUpper: {
		aliases:     []string{"upper", "up"},
		description: "Takes a single parameter value and returns the value in upper case.",
		block:       UpperFunctionBlock{},
	},
	FunctionNameTitle: {
		aliases:     []string{"title"},
		description: "Takes a single parameter value and returns the value in title case.",
		block: TitleFunctionBlock{
			caser: cases.Title(language.English),
		},
	},
	FunctionNameBase64: {
		aliases:     []string{"base64", "b64"},
		description: "Takes a single parameter value and returns the base64 encoding of the value.",
		block:       Base64FunctionBlock{},
	},
	FunctionNameLength: {
		aliases:     []string{"length", "len"},
		description: "Takes a single parameter value and returns its length.",
		block:       LengthFunctionBlock{},
	},
	FunctionNameTrim: {
		aliases:     []string{"trim"},
		description: "Takes a single parameter value and returns its value after removing leading and trailing space character(s).",
		block:       TrimFunctionBlock{},
	},
	FunctionNameLeftTrim: {
		aliases:     []string{"ltrim", "lefttrim"},
		description: "Takes a single parameter value and returns its value after removing leading space character(s).",
		block:       LeftTrimFunctionBlock{},
	},
	FunctionNameRightTrim: {
		aliases:     []string{"rtrim", "righttrim"},
		description: "Takes a single parameter value and returns its value after removing trailing space character(s).",
		block:       RightTrimFunctionBlock{},
	},
	FunctionNameIfBlank: {
		aliases:     []string{"ifblank"},
		description: "Takes two parameter values and returns the first one if it is not empty \nand doesn't consist solely of whitespace characters, \nelse returns the second parameter value.",
		block:       IfBlankFunctionBlock{},
	},
	FunctionNameStartsWith: {
		aliases:     []string{"startswith"},
		description: "Takes two parameter values and returns true if the first parameter value starts with the second one.",
		block:       StartsWithFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameEndsWith: {
		aliases:     []string{"endswith"},
		description: "Takes two parameter values and returns true if the first parameter value ends with the second one.",
		block:       EndsWithFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameNow: {
		aliases:     []string{"now"},
		description: "Returns the current date/time.",
		block:       NowFunctionBlock{},
	},
	FunctionNameCurrentDay: {
		aliases:     []string{"cday", "currentday"},
		description: "Returns the current day. If today is 9th September 2022, cday() will return 9.",
		block:       CurrentDayFunctionBlock{},
	},
	FunctionNameCurrentDate: {
		aliases:     []string{"cdate", "currentdate"},
		description: "Returns the current date formatted as year-month-day. \nIf today is 9th September 2022, cdate() will return 2022-September-09.",
		block:       CurrentDateFunctionBlock{},
	},
	FunctionNameCurrentMonth: {
		aliases:     []string{"cmonth", "cmon", "currentmonth", "currentmon"},
		description: "Returns the current month. \nIf today is 9th September 2022, cmonth() will return September.",
		block:       CurrentMonthFunctionBlock{},
	},
	FunctionNameCurrentYear: {
		aliases:     []string{"cyear", "cyr", "currentyear", "currentyr"},
		description: "Returns the current year. \nIf today is 9th September 2022, cyr() will return 2022.",
		block:       CurrentYearFunctionBlock{},
	},
	FunctionNameDayOfWeek: {
		aliases:     []string{"dayofweek", "dow"},
		description: "Returns the day of the week. \nIf today is a Friday, dow() will return Friday.",
		block:       DayOfWeekFunctionBlock{},
	},
	FunctionNameExtract: {
		aliases:     []string{"extract"},
		description: "Returns the extracted component from date/time. extract allows the extraction of date, day, year, month and weekday from date/time. \nFor example, extract(atime, month) will extract 'month' from the access time of a file.",
		block:       ExtractFunctionBlock{},
	},
	FunctionNameHoursDifference: {
		aliases:     []string{"hoursdifference", "hourdifference", "hoursdiff", "hourdiff"},
		description: "Returns the difference between 2 date/times in hours.",
		block:       HoursDifferenceFunctionBlock{},
	},
	FunctionNameDaysDifference: {
		aliases:     []string{"daysdifference", "daydifference", "daysdiff", "daydiff"},
		description: "Returns the difference between 2 date/times in days.",
		block:       DaysDifferenceFunctionBlock{},
	},
	FunctionNameDateTimeParse: {
		aliases:     []string{"parsedatetime", "parsedttime", "parsedttm", "parsedatetm"},
		description: "Returns the time representation after parsing the input string. \nIt takes 2 parameters, the first parameter is a string to be parsed and the second is the format identifier. Example, parsedatetime(2022-09-09, dt) \nreturns the date/time represented by the given input.",
		block:       ParseDateTimeFunctionBlock{},
	},
	FunctionNameWorkingDirectory: {
		aliases:     []string{"cwd", "wd"},
		description: "Returns working directory.",
		block:       WorkingDirectoryFunctionBlock{},
	},
	FunctionNameConcat: {
		aliases:     []string{"concat"},
		description: "Takes variable number of parameter values and returns a string concatenated of all these values.",
		block:       ConcatFunctionBlock{},
	},
	FunctionNameConcatWithSeparator: {
		aliases:     []string{"concatws", "concatwithseparator"},
		description: "Takes variable number of parameter values and returns a string concatenated of all these values. \nThis function uses the last parameter value as a separator.",
		block:       ConcatWithSeparatorFunctionBlock{},
	},
	FunctionNameContains: {
		aliases:     []string{"contains"},
		description: "Returns true, if the second parameter value is present within the first. \nFor example, contains(hello, lo) will return true.",
		block:       ContainsFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameSubstring: {
		aliases:     []string{"substr", "str"},
		description: "Returns a substring from the main string. \nsubstr() takes 3 parameter values, first parameter value is the main string, second is the starting index (starting from 0) and the optional third \nparameter value is the end index(inclusive).",
		block:       SubstringFunctionBlock{},
	},
	FunctionNameReplace: {
		aliases:     []string{"replace"},
		description: "Replaces the first occurrence of an old string with the new string. \nFor example, replace(name, test, best) will replace the first occurrence of the string 'test' with 'best' in the file name.",
		block:       ReplaceFunctionBlock{},
	},
	FunctionNameReplaceAll: {
		aliases:     []string{"replaceall"},
		description: "Replaces all the occurrences of an old string with the new string. \nFor example, replaceall(name, test, best) will replace all the occurrences of the string 'test' with 'best' in the file name.",
		block:       ReplaceAllFunctionBlock{},
	},
	FunctionNameIsFileTypeText: {
		aliases:     []string{"istext", "istxt"},
		description: "Returns true if the mime type of a file is text/plain, false otherwise.  \nFor example, the common use of this function is with mime attribute, istext(mime).",
		block:       IsFileTypeTextFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameIsFileTypeImage: {
		aliases:     []string{"isimage", "isimg"},
		description: "Returns true if the mime type of a file is an image, false otherwise.  \nFor example, the common use of this function is with mime attribute, isimage(mime).",
		block:       IsFileTypeImageFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameIsFileTypeAudio: {
		aliases:     []string{"isaudio"},
		description: "Returns true if the mime type of a file is an audio, false otherwise.  \nFor example, the common use of this function is with mime attribute, isaudio(mime).",
		block:       IsFileTypeAudioFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameIsFileTypeVideo: {
		aliases:     []string{"isvideo"},
		description: "Returns true if the mime type of a file is video, false otherwise.  \nFor example, the common use of this function is with mime attribute, isvideo(mime).",
		block:       IsFileTypeVideoFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameIsFileTypePdf: {
		aliases:     []string{"ispdf"},
		description: "Returns true if the mime type of a file is pdf, false otherwise.  \nFor example, the common use of this function is with mime attribute, ispdf(mime).",
		block:       IsFileTypePdfFunctionBlock{},
		tags:        map[string]bool{"where": true},
	},
	FunctionNameIsFileTypeArchive: {
		aliases:     []string{"isarchive"},
		description: "Returns true if the mime type of a file is an archive, false otherwise.  \nFor example, the common use of this function is with mime attribute, isarchive(mime).",
		block: IsFileTypeArchiveFunctionBlock{
			matchingMimeTypes: map[string]bool{
				"application/x-7z-compressed":   true,
				"application/zip":               true,
				"application/x-zip":             true,
				"application/x-zip-compressed":  true,
				"application/epub+zip":          true,
				"application/jar":               true,
				"application/x-archive":         true,
				"application/x-unix-archive":    true,
				"application/x-tar":             true,
				"application/x-xar":             true,
				"application/x-bzip2":           true,
				"application/gzip":              true,
				"application/x-gzip":            true,
				"application/x-gunzip":          true,
				"application/gzipped":           true,
				"application/gzip-compressed":   true,
				"application/x-gzip-compressed": true,
				"gzip/document":                 true,
				"application/x-rar-compressed":  true,
				"application/x-rar":             true,
				"application/lzip":              true,
				"application/x-lzip":            true,
			},
		},
		tags: map[string]bool{"where": true},
	},
	FunctionNameFormatSize: {
		aliases:     []string{"formatsize", "fmtsize"},
		description: "Returns a human readable file size in IEC units.  \nThese include B, KiB, MiB, GiB, TiB, PiB, EiB. This function takes a single parameter.",
		block:       FormatSizeFunctionBlock{},
	},
	FunctionNameParseSize: {
		aliases:     []string{"parsesize", "psize"},
		description: "Parses the input string into the number of bytes it represents.  \nFor example, parsesize(42 MB) returns 42000000, parsesize(42 mib) returns 44040192, parsesize(10.23 Mib) returns 10726932.  \nSize unit must be one of the following: B, KiB, MiB, GiB, TiB, PiB, EiB, kB, MB, GB, TB, PB, EB.",
		block:       ParseSizeFunctionBlock{},
	},
	FunctionNameCount: {
		aliases:        []string{"count"},
		description:    "count is an aggregate function that returns the total number of entries in the source directory. It does not take any parameter.",
		isAggregate:    true,
		aggregateBlock: &CountFunctionBlock{},
	},
	FunctionNameCountDistinct: {
		aliases:        []string{"countdistinct", "countd"},
		description:    "countdistinct is an aggregate function that returns the distinct number of entries based on the parameter type. \nFor example, countdistinct(ext) will return the count of the distinct file extensions in the source directory.",
		isAggregate:    true,
		aggregateBlock: &CountDistinctFunctionBlock{},
	},
	FunctionNameSum: {
		aliases:        []string{"summation", "sum"},
		description:    "sum is an aggregate function that returns the sum of all the values corresponding to the provided parameter. \nFor example, sum(size) will return the sum of size of all the files in the source directory.",
		isAggregate:    true,
		aggregateBlock: &SumFunctionBlock{},
	},
	FunctionNameAverage: {
		aliases:        []string{"average", "avg"},
		description:    "average is an aggregate function that returns the average of all the values corresponding to the provided parameter. \nFor example, avg(size) will return the average file size in the source directory.",
		isAggregate:    true,
		aggregateBlock: &AverageFunctionBlock{},
	},
	FunctionNameMin: {
		aliases:        []string{"min"},
		description:    "min is an aggregate function that returns the minimum of all the values corresponding to the provided parameter. \nFor example, min(size) will return the minimum file size in the source directory.",
		isAggregate:    true,
		aggregateBlock: &MinFunctionBlock{},
	},
	FunctionNameMax: {
		aliases:        []string{"max"},
		description:    "max is an aggregate function that returns the maximum of all the values corresponding to the provided parameter. \nFor example, max(size) will return the maximum file size in the source directory.",
		isAggregate:    true,
		aggregateBlock: &MaxFunctionBlock{},
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

func (functions *AllFunctions) DescriptionOf(fn string) string {
	definition, ok := functions.supportedFunctions[strings.ToLower(fn)]
	if ok {
		return definition.description
	}
	return ""
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
