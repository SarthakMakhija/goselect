package context

import (
	"errors"
	"fmt"
	"goselect/parser/error/messages"
	"time"
)

type valueType int

const (
	ValueTypeString    = 1
	ValueTypeInt       = 2
	ValueTypeInt64     = 3
	ValueTypeDateTime  = 4
	ValueTypeUndefined = 5
)

type Value struct {
	valueType   valueType
	stringValue string
	intValue    int
	int64Value  int64
	timeValue   time.Time
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

func (value Value) Get() interface{} {
	switch value.valueType {
	case ValueTypeString:
		return value.stringValue
	case ValueTypeInt:
		return value.intValue
	case ValueTypeInt64:
		return value.int64Value
	case ValueTypeDateTime:
		return value.timeValue
	case ValueTypeUndefined:
		return ""
	}
	return ""
}
