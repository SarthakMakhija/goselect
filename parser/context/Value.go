package context

import (
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"strconv"
	"strings"
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
	case ValueTypeString:
		return strconv.ParseFloat(value.stringValue, 64)
	}
	return -1, errors.New(messages.ErrorMessageExpectedNumericArgument)
}

func (value Value) CompareTo(other Value) int {
	receiver, arg := value, other
	if value.valueType != other.valueType {
		if rec, ar, possible, err := value.attemptCommonType(other); err != nil {
			return -1
		} else if possible {
			receiver, arg = rec, ar
		} else {
			return -1
		}
	}
	switch receiver.valueType {
	case ValueTypeString:
		first, second := receiver.stringValue, arg.stringValue
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeInt:
		first, second := receiver.intValue, arg.intValue
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeInt64:
		first, second := receiver.int64Value, arg.int64Value
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeUint32:
		first, second := receiver.uint32Value, arg.uint32Value
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeFloat64:
		first, second := receiver.float64Value, arg.float64Value
		if first == second {
			return 0
		}
		if first < second {
			return -1
		}
		return 1
	case ValueTypeBoolean:
		first, second := receiver.booleanValue, arg.booleanValue
		if first == second {
			return 0
		}
		return 1
	case ValueTypeDateTime:
		first, second := receiver.timeValue, arg.timeValue
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

func (value Value) attemptCommonType(other Value) (Value, Value, bool, error) {
	switch {
	case (value.valueType == ValueTypeInt ||
		value.valueType == ValueTypeInt64 ||
		value.valueType == ValueTypeUint32) && other.valueType == ValueTypeFloat64:
		if v, err := value.GetNumericAsFloat64(); err != nil {
			return value, other, false, err
		} else {
			return Float64Value(v), other, true, nil
		}
	case value.valueType == ValueTypeFloat64 &&
		(other.valueType == ValueTypeInt ||
			other.valueType == ValueTypeInt64 ||
			other.valueType == ValueTypeUint32):
		if v, err := other.GetNumericAsFloat64(); err != nil {
			return value, other, false, err
		} else {
			return value, Float64Value(v), true, nil
		}
	case value.valueType == ValueTypeInt &&
		other.valueType == ValueTypeInt64:
		return Int64Value(int64(value.intValue)), other, true, nil
	case value.valueType == ValueTypeInt64 &&
		other.valueType == ValueTypeInt:
		return value, Int64Value(int64(other.intValue)), true, nil
	case value.valueType == ValueTypeInt &&
		other.valueType == ValueTypeUint32:
		return Uint32Value(uint32(value.intValue)), other, true, nil
	case value.valueType == ValueTypeUint32 &&
		other.valueType == ValueTypeInt:
		return value, Uint32Value(uint32(other.intValue)), true, nil
	case value.valueType == ValueTypeUint32 &&
		other.valueType == ValueTypeInt64:
		if v, err := value.GetNumericAsFloat64(); err != nil {
			return value, other, false, err
		} else {
			if o, err := other.GetNumericAsFloat64(); err != nil {
				return value, other, false, err
			} else {
				return Float64Value(v), Float64Value(o), true, nil
			}
		}
	case value.valueType == ValueTypeString &&
		other.valueType == ValueTypeBoolean:
		if strings.ToLower(value.stringValue) == "true" || strings.ToLower(value.stringValue) == "y" {
			return BooleanValue(true), other, true, nil
		}
		if strings.ToLower(value.stringValue) == "false" || strings.ToLower(value.stringValue) == "n" {
			return BooleanValue(false), other, true, nil
		}
		return value, other, false, nil
	case value.valueType == ValueTypeBoolean &&
		other.valueType == ValueTypeString:
		if strings.ToLower(other.stringValue) == "true" || strings.ToLower(other.stringValue) == "y" {
			return value, BooleanValue(true), true, nil
		}
		if strings.ToLower(other.stringValue) == "false" || strings.ToLower(other.stringValue) == "n" {
			return value, BooleanValue(false), true, nil
		}
		return value, other, false, nil
	case value.valueType == ValueTypeString && other.isNumericType():
		if v, err := value.GetNumericAsFloat64(); err != nil {
			return value, other, false, err
		} else {
			if o, err := other.GetNumericAsFloat64(); err != nil {
				return value, other, false, err
			} else {
				return Float64Value(v), Float64Value(o), true, nil
			}
		}
	case value.isNumericType() && other.valueType == ValueTypeString:
		if v, err := value.GetNumericAsFloat64(); err != nil {
			return value, other, false, err
		} else {
			if o, err := other.GetNumericAsFloat64(); err != nil {
				return value, other, false, err
			} else {
				return Float64Value(v), Float64Value(o), true, nil
			}
		}
	}
	//Handle time
	return value, other, false, nil
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
		return strconv.FormatFloat(value.float64Value, 'f', 2, 64)
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

func (value Value) isNumericType() bool {
	if value.valueType == ValueTypeInt ||
		value.valueType == ValueTypeInt64 ||
		value.valueType == ValueTypeUint32 ||
		value.valueType == ValueTypeFloat64 {
		return true
	}
	return false
}
