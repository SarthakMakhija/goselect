package context

import (
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"strconv"
	"time"
)

type valueType int

const (
	ValueTypeString    = 1
	ValueTypeInt       = 2
	ValueTypeInt64     = 3
	ValueTypeDateTime  = 4
	ValueTypeBoolean   = 5
	ValueTypeUint32    = 6
	ValueTypeFloat64   = 7
	ValueTypeUndefined = 8
)

type Value struct {
	valueType    valueType
	stringValue  string
	intValue     int
	int64Value   int64
	booleanValue bool
	uint32Value  uint32
	float64Value float64
	timeValue    time.Time
}

func StringValue(value string) Value {
	return Value{
		stringValue: value,
		valueType:   ValueTypeString,
	}
}

func IntValue(value int) Value {
	return Value{
		intValue:  value,
		valueType: ValueTypeInt,
	}
}

func Int64Value(value int64) Value {
	return Value{
		int64Value: value,
		valueType:  ValueTypeInt64,
	}
}

func Uint32Value(value uint32) Value {
	return Value{
		uint32Value: value,
		valueType:   ValueTypeUint32,
	}
}

func Float64Value(value float64) Value {
	return Value{
		float64Value: value,
		valueType:    ValueTypeFloat64,
	}
}

func BooleanValue(value bool) Value {
	return Value{
		booleanValue: value,
		valueType:    ValueTypeBoolean,
	}
}

func DateTimeValue(time time.Time) Value {
	return Value{
		timeValue: time,
		valueType: ValueTypeDateTime,
	}
}

func EmptyValue() Value {
	return Value{valueType: ValueTypeUndefined}
}

func (value Value) GetString() (string, error) {
	if value.valueType != ValueTypeString {
		return "", errors.New(fmt.Sprintf(messages.ErrorMessageIncorrectValueType, "string"))
	}
	return value.stringValue, nil
}

func (value Value) GetInt() (int, error) {
	if value.valueType != ValueTypeInt {
		return -1, errors.New(fmt.Sprintf(messages.ErrorMessageIncorrectValueType, "int"))
	}
	return value.intValue, nil
}

func (value Value) GetBoolean() (bool, error) {
	if value.valueType != ValueTypeBoolean {
		return false, errors.New(fmt.Sprintf(messages.ErrorMessageIncorrectValueType, "boolean"))
	}
	return value.booleanValue, nil
}

func (value Value) GetNumericAsFloat64() (float64, error) {
	switch value.valueType {
	case ValueTypeInt:
		return float64(value.intValue), nil
	case ValueTypeInt64:
		return float64(value.int64Value), nil
	case ValueTypeUint32:
		return float64(value.uint32Value), nil
	case ValueTypeFloat64:
		return value.float64Value, nil
	}
	return -1, errors.New(messages.ErrorMessageExpectedNumericArgument)
}

func (value Value) CompareTo(other Value) int {
	switch value.valueType {
	case ValueTypeString:
		first, second := value.stringValue, other.stringValue
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeInt:
		first, second := value.intValue, other.intValue
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeInt64:
		first, second := value.int64Value, other.int64Value
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeUint32:
		first, second := value.uint32Value, other.uint32Value
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeFloat64:
		first, second := value.float64Value, other.float64Value
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeBoolean:
		first, second := value.booleanValue, other.booleanValue
		if first == second {
			return 0
		}
		return 1
	case ValueTypeDateTime:
		first, second := value.timeValue, other.timeValue
		if first.Equal(second) {
			return 0
		}
		if first.Before(second) {
			return -1
		}
		return 1
	}
	return -1
}

func (value Value) GetAsString() string {
	switch value.valueType {
	case ValueTypeString:
		return value.stringValue
	case ValueTypeInt:
		return strconv.Itoa(value.intValue)
	case ValueTypeInt64:
		return strconv.FormatInt(value.int64Value, 10)
	case ValueTypeUint32:
		return strconv.Itoa(int(value.uint32Value))
	case ValueTypeFloat64:
		return strconv.FormatFloat(float64(value.float64Value), 'f', 2, 64)
	case ValueTypeBoolean:
		if value.booleanValue {
			return "Y"
		}
		return "N"
	case ValueTypeDateTime:
		return value.timeValue.String()
	}
	return ""
}
