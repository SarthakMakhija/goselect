package context

import (
	"fmt"
	"goselect/parser/error/messages"
	"goselect/parser/tokenizer"
	"strconv"
	"strings"
)

type TypePair struct {
	aType valueType
	bType valueType
}

type toCommonTypeValueFunction = func(aValue Value, bValue Value) (Value, Value, error)

var toTargetConversions = map[TypePair]toCommonTypeValueFunction{
	TypePair{aType: ValueTypeInt, bType: ValueTypeInt64}: func(aValue Value, bValue Value) (Value, Value, error) {
		return Int64Value(int64(aValue.intValue)), bValue, nil
	},
	TypePair{aType: ValueTypeInt64, bType: ValueTypeInt}: func(aValue Value, bValue Value) (Value, Value, error) {
		return aValue, Int64Value(int64(bValue.intValue)), nil
	},
	TypePair{aType: ValueTypeInt, bType: ValueTypeUint32}: func(aValue Value, bValue Value) (Value, Value, error) {
		return Uint32Value(uint32(aValue.intValue)), bValue, nil
	},
	TypePair{aType: ValueTypeUint32, bType: ValueTypeInt}: func(aValue Value, bValue Value) (Value, Value, error) {
		return aValue, Uint32Value(uint32(bValue.intValue)), nil
	},
	TypePair{aType: ValueTypeInt, bType: ValueTypeFloat64}: func(aValue Value, bValue Value) (Value, Value, error) {
		return Float64Value(float64(aValue.intValue)), bValue, nil
	},
	TypePair{aType: ValueTypeFloat64, bType: ValueTypeInt}: func(aValue Value, bValue Value) (Value, Value, error) {
		return aValue, Float64Value(float64(bValue.intValue)), nil
	},
	TypePair{aType: ValueTypeUint32, bType: ValueTypeFloat64}: func(aValue Value, bValue Value) (Value, Value, error) {
		return Float64Value(float64(aValue.uint32Value)), bValue, nil
	},
	TypePair{aType: ValueTypeFloat64, bType: ValueTypeUint32}: func(aValue Value, bValue Value) (Value, Value, error) {
		return aValue, Float64Value(float64(bValue.uint32Value)), nil
	},
	TypePair{aType: ValueTypeInt64, bType: ValueTypeFloat64}: func(aValue Value, bValue Value) (Value, Value, error) {
		return Float64Value(float64(aValue.int64Value)), bValue, nil
	},
	TypePair{aType: ValueTypeFloat64, bType: ValueTypeInt64}: func(aValue Value, bValue Value) (Value, Value, error) {
		return aValue, Float64Value(float64(bValue.int64Value)), nil
	},
	TypePair{aType: ValueTypeFloat64, bType: ValueTypeFloat64}: func(aValue Value, bValue Value) (Value, Value, error) {
		return aValue, bValue, nil
	},
	TypePair{aType: ValueTypeUint32, bType: ValueTypeInt64}: func(aValue Value, bValue Value) (Value, Value, error) {
		return Float64Value(float64(aValue.uint32Value)), Float64Value(float64(bValue.int64Value)), nil
	},
	TypePair{aType: ValueTypeInt64, bType: ValueTypeUint32}: func(aValue Value, bValue Value) (Value, Value, error) {
		return Float64Value(float64(aValue.int64Value)), Float64Value(float64(bValue.uint32Value)), nil
	},
	TypePair{aType: ValueTypeString, bType: ValueTypeBoolean}: func(aValue Value, bValue Value) (Value, Value, error) {
		if strings.ToLower(aValue.stringValue) == "true" || strings.ToLower(aValue.stringValue) == "y" {
			return trueBooleanValue, bValue, nil
		}
		if strings.ToLower(aValue.stringValue) == "false" || strings.ToLower(aValue.stringValue) == "n" {
			return falseBooleanValue, bValue, nil
		}
		return aValue, bValue, fmt.Errorf(messages.ErrorMessageCannotConvertToBoolean, aValue.stringValue)
	},
	TypePair{aType: ValueTypeBoolean, bType: ValueTypeString}: func(aValue Value, bValue Value) (Value, Value, error) {
		if strings.ToLower(bValue.stringValue) == "true" || strings.ToLower(bValue.stringValue) == "y" {
			return aValue, trueBooleanValue, nil
		}
		if strings.ToLower(bValue.stringValue) == "false" || strings.ToLower(bValue.stringValue) == "n" {
			return aValue, falseBooleanValue, nil
		}
		return aValue, bValue, fmt.Errorf(messages.ErrorMessageCannotConvertToBoolean, bValue.stringValue)
	},
	TypePair{aType: ValueTypeString, bType: ValueTypeInt}: func(aValue Value, bValue Value) (Value, Value, error) {
		v, err := strconv.Atoi(aValue.stringValue)
		if err != nil {
			return aValue, bValue, err
		}
		return IntValue(v), bValue, nil
	},
	TypePair{aType: ValueTypeInt, bType: ValueTypeString}: func(aValue Value, bValue Value) (Value, Value, error) {
		v, err := strconv.Atoi(bValue.stringValue)
		if err != nil {
			return aValue, bValue, err
		}
		return aValue, IntValue(v), nil
	},
	TypePair{aType: ValueTypeString, bType: ValueTypeInt64}: func(aValue Value, bValue Value) (Value, Value, error) {
		v, err := strconv.ParseInt(aValue.stringValue, 10, 64)
		if err != nil {
			return aValue, bValue, err
		}
		return Int64Value(v), bValue, nil
	},
	TypePair{aType: ValueTypeInt64, bType: ValueTypeString}: func(aValue Value, bValue Value) (Value, Value, error) {
		v, err := strconv.ParseInt(bValue.stringValue, 10, 64)
		if err != nil {
			return aValue, bValue, err
		}
		return aValue, Int64Value(v), nil
	},
	TypePair{aType: ValueTypeString, bType: ValueTypeUint32}: func(aValue Value, bValue Value) (Value, Value, error) {
		v, err := strconv.ParseUint(aValue.stringValue, 10, 32)
		if err != nil {
			return aValue, bValue, err
		}
		return Uint32Value(uint32(v)), bValue, nil
	},
	TypePair{aType: ValueTypeUint32, bType: ValueTypeString}: func(aValue Value, bValue Value) (Value, Value, error) {
		v, err := strconv.ParseUint(bValue.stringValue, 10, 32)
		if err != nil {
			return aValue, bValue, err
		}
		return aValue, Uint32Value(uint32(v)), nil
	},
	TypePair{aType: ValueTypeString, bType: ValueTypeFloat64}: func(aValue Value, bValue Value) (Value, Value, error) {
		v, err := strconv.ParseFloat(aValue.stringValue, 64)
		if err != nil {
			return aValue, bValue, err
		}
		return Float64Value(v), bValue, nil
	},
	TypePair{aType: ValueTypeFloat64, bType: ValueTypeString}: func(aValue Value, bValue Value) (Value, Value, error) {
		v, err := strconv.ParseFloat(bValue.stringValue, 64)
		if err != nil {
			return aValue, bValue, err
		}
		return aValue, Float64Value(v), nil
	},
}

func ToValue(token tokenizer.Token) (Value, error) {
	switch token.TokenType {
	case tokenizer.Numeric:
		v, err := strconv.ParseInt(token.TokenValue, 10, 64)
		if err != nil {
			return EmptyValue, err
		}
		return Int64Value(v), nil
	case tokenizer.FloatingPoint:
		v, err := strconv.ParseFloat(token.TokenValue, 64)
		if err != nil {
			return EmptyValue, err
		}
		return Float64Value(v), nil
	case tokenizer.Boolean:
		if strings.ToLower(token.TokenValue) == "true" || strings.ToLower(token.TokenValue) == "y" {
			return trueBooleanValue, nil
		}
		if strings.ToLower(token.TokenValue) == "false" || strings.ToLower(token.TokenValue) == "n" {
			return falseBooleanValue, nil
		}
		return StringValue(token.TokenValue), nil
	default:
		return StringValue(token.TokenValue), nil
	}
}

func getCommonType(value Value, other Value, typePair TypePair) (Value, Value, error) {
	fn, ok := toTargetConversions[typePair]
	if !ok {
		return value, other, fmt.Errorf(messages.ErrorMessageUndefinedConversionFunction, typePair.aType, typePair.bType)
	}
	return fn(value, other)
}

func toFloat64(value Value) (Value, error) {
	typePair := TypePair{aType: value.valueType, bType: ValueTypeFloat64}
	fn, ok := toTargetConversions[typePair]
	if !ok {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageUndefinedConversionFunction, typePair.aType, typePair.bType)
	}
	asFloat64Value, _, err := fn(value, EmptyValue)
	if err != nil {
		return EmptyValue, err
	}
	return asFloat64Value, nil
}
