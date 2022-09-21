//go:build unit
// +build unit

package context

import (
	"goselect/parser/tokenizer"
	"testing"
	"time"
)

func TestCompareIntEmptyValue(t *testing.T) {
	value := IntValue(10)
	other := EmptyValue

	if value.CompareTo(other) != CompareToNotPossible {
		t.Fatalf("Expected comparison between int and empty value to be not possible but was possible")
	}
}

func TestCompareIntInt(t *testing.T) {
	value := IntValue(10)
	other := IntValue(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int values to match but they did not")
	}
}

func TestCompareIntInt2(t *testing.T) {
	value := IntValue(10)
	other := IntValue(20)

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected int values to not match but they did")
	}
}

func TestCompareIntString(t *testing.T) {
	value := IntValue(10)
	other := StringValue("10")

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int and string values to match but they did not")
	}
}

func TestCompareInt64Int64(t *testing.T) {
	value := Int64Value(10)
	other := Int64Value(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int64 and int64 values to match but they did not")
	}
}

func TestCompareInt64Int64LessThan(t *testing.T) {
	value := Int64Value(5)
	other := Int64Value(10)

	if value.CompareTo(other) != CompareToLessThan {
		t.Fatalf("Expected int64 value to be less than the other but was not")
	}
}

func TestCompareInt64Int64GreaterThan(t *testing.T) {
	value := Int64Value(15)
	other := Int64Value(10)

	if value.CompareTo(other) != CompareToGreaterThan {
		t.Fatalf("Expected int64 value to be greater than the other but was not")
	}
}

func TestCompareIntInt64(t *testing.T) {
	value := IntValue(10)
	other := Int64Value(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int and int64 values to match but they did not")
	}
}

func TestCompareInt64Int(t *testing.T) {
	value := Int64Value(10)
	other := IntValue(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int and int64 values to match but they did not")
	}
}

func TestCompareToUin32Uint32(t *testing.T) {
	value := Uint32Value(10)
	other := Uint32Value(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected uint32 and uint32 values to match but they did not")
	}
}

func TestCompareToUin32Uint32LessThan(t *testing.T) {
	value := Uint32Value(5)
	other := Uint32Value(10)

	if value.CompareTo(other) != CompareToLessThan {
		t.Fatalf("Expected uint32 value to be less than the other but was not")
	}
}

func TestCompareToUin32Uint32GreaterThan(t *testing.T) {
	value := Uint32Value(15)
	other := Uint32Value(10)

	if value.CompareTo(other) != CompareToGreaterThan {
		t.Fatalf("Expected uint32 value to be greater than the other but was not")
	}
}

func TestCompareToIntUint32(t *testing.T) {
	value := IntValue(10)
	other := Uint32Value(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int and uint32 values to match but they did not")
	}
}

func TestCompareToUint32Int(t *testing.T) {
	value := Uint32Value(10)
	other := IntValue(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int and uint32 values to match but they did not")
	}
}

func TestCompareToIntFloat(t *testing.T) {
	value := IntValue(10)
	other := Float64Value(10.0)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int and float64 values to match but they did not")
	}
}

func TestCompareToFloatInt(t *testing.T) {
	value := Float64Value(10)
	other := IntValue(10.0)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int and float64 values to match but they did not")
	}
}

func TestCompareToUint32Float64(t *testing.T) {
	value := Uint32Value(10)
	other := Float64Value(10.0)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected uin32 and float64 values to match but they did not")
	}
}

func TestCompareToFloat64Uint32(t *testing.T) {
	value := Float64Value(10.0)
	other := Uint32Value(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected uin32 and float64 values to match but they did not")
	}
}

func TestCompareToInt64Float64(t *testing.T) {
	value := Int64Value(10)
	other := Float64Value(10.0)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int64 and float64 values to match but they did not")
	}
}

func TestCompareToUint32Int64(t *testing.T) {
	value := Uint32Value(10)
	other := Int64Value(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected uint32 and int64 values to match but they did not")
	}
}

func TestCompareToInt64Uint32(t *testing.T) {
	value := Int64Value(10)
	other := Uint32Value(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected uint32 and int64 values to match but they did not")
	}
}

func TestCompareToStringFloat64(t *testing.T) {
	value := StringValue("3.27")
	other := Float64Value(3.27)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected string and float64 values to match but they did not")
	}
}

func TestCompareToTrueBooleanString(t *testing.T) {
	value := trueBooleanValue
	other := StringValue("Y")

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected string and boolean values to match but they did not")
	}
}

func TestCompareToFalseBooleanString1(t *testing.T) {
	value := falseBooleanValue
	other := StringValue("Y")

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToFalseBooleanString2(t *testing.T) {
	value := falseBooleanValue
	other := StringValue("n")

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected string and boolean values match but they did not")
	}
}

func TestCompareToFalseBooleanInvalidString(t *testing.T) {
	value := falseBooleanValue
	other := StringValue("ok")

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToInvalidStringTrueBoolean(t *testing.T) {
	value := StringValue("ok")
	other := trueBooleanValue

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToStringTrueBoolean(t *testing.T) {
	value := StringValue("n")
	other := trueBooleanValue

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToStringFalseBoolean(t *testing.T) {
	value := StringValue("Y")
	other := falseBooleanValue

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToStringInt1(t *testing.T) {
	value := StringValue("10")
	other := IntValue(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected string and int values to match but they did not")
	}
}

func TestCompareToStringInt2(t *testing.T) {
	value := StringValue("a")
	other := IntValue(10)

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and int values to not match but they did")
	}
}

func TestCompareToStringInt641(t *testing.T) {
	value := StringValue("10")
	other := Int64Value(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected string and int64 values to match but they did not")
	}
}

func TestCompareToStringInt642(t *testing.T) {
	value := StringValue("a")
	other := Int64Value(10)

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and int64 values to not match but they did")
	}
}

func TestCompareToStringUint321(t *testing.T) {
	value := StringValue("10")
	other := Uint32Value(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected string and uint32 values to match but they did not")
	}
}

func TestCompareToStringUint322(t *testing.T) {
	value := StringValue("a")
	other := Uint32Value(10)

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and uint32 values to not match but they did")
	}
}

func TestCompareToIntFloat64(t *testing.T) {
	value := IntValue(10)
	other := Float64Value(10.1)

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected int and float64 values to not match but they did")
	}
}

func TestCompareToIntString(t *testing.T) {
	value := IntValue(10)
	other := StringValue("content")

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected int and string values to not match but they did")
	}
}

func TestCompareToFloat64String1(t *testing.T) {
	value := Float64Value(10)
	other := StringValue("content")

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected float64 and string values to not match but they did")
	}
}

func TestCompareToFloat64String2(t *testing.T) {
	value := Float64Value(10.2)
	other := StringValue("10.2")

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected float64 and string values to match but they did not")
	}
}

func TestCompareToFloat64Int1(t *testing.T) {
	value := Float64Value(10.1)
	other := IntValue(10)

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected int and float64 values to not match but they did")
	}
}

func TestCompareToFloat64Int2(t *testing.T) {
	value := Float64Value(10)
	other := IntValue(10)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected int and float64 values to match but they did not")
	}
}

func TestCompareToInt64String1(t *testing.T) {
	value := Int64Value(3)
	other := StringValue("3")

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected string and int64 values to match but they did not")
	}
}

func TestCompareToInt64String2(t *testing.T) {
	value := Int64Value(3)
	other := StringValue("a")

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and int64 values to not match but they did")
	}
}

func TestCompareToUint32String1(t *testing.T) {
	value := Uint32Value(3)
	other := StringValue("3")

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected string and uint32 values to match but they did not")
	}
}

func TestCompareToUint32String2(t *testing.T) {
	value := Uint32Value(3)
	other := StringValue("a")

	if value.CompareTo(other) == CompareToEqual {
		t.Fatalf("Expected string and uint32 values to not match but they did")
	}
}

func TestCompareToFloat64Int64(t *testing.T) {
	value := Float64Value(3.0)
	other := Int64Value(3)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected float64 and int64 values to match but they did not")
	}
}

func TestGetCommonType1(t *testing.T) {
	value := Float64Value(3.0)
	other := Int64Value(3)

	first, second, _ := getCommonType(value, other, TypePair{aType: ValueTypeFloat64, bType: ValueTypeInt64})
	if first.CompareTo(second) != CompareToEqual {
		t.Fatalf("Expected first and second values to match but they did not")
	}
}

func TestGetCommonType2(t *testing.T) {
	value := StringValue("a")
	other := Int64Value(3)

	_, _, err := getCommonType(value, other, TypePair{aType: ValueTypeString, bType: ValueTypeInt64})
	if err == nil {
		t.Fatalf("Expected an error while trying to convert invalid string to int64")
	}
}

func TestGetCommonType3(t *testing.T) {
	value := StringValue("a")
	other := Int64Value(3)

	_, _, err := getCommonType(value, other, TypePair{aType: ValueTypeString, bType: ValueTypeUndefined})
	if err == nil {
		t.Fatalf("Expected an error while trying to convert string to undefined type")
	}
}

func TestGetIntGivenValueWithInt(t *testing.T) {
	value := IntValue(10)
	expected := 10

	actual, _ := value.GetInt()
	if actual != expected {
		t.Fatalf("Expected GetInt to be %v, received %v", expected, actual)
	}
}

func TestGetIntGivenValueAsNonInt(t *testing.T) {
	value := Int64Value(10)

	_, err := value.GetInt()
	if err == nil {
		t.Fatalf("Expected error while trying to get int for an int64 value")
	}
}

func TestCompareToDateTime(t *testing.T) {
	now := time.Now()
	value := DateTimeValue(now)
	other := DateTimeValue(now)

	if value.CompareTo(other) != CompareToEqual {
		t.Fatalf("Expected date/time value to match with other but did not")
	}
}

func TestCompareToDateTimeLessThan(t *testing.T) {
	now := time.Now()
	value := DateTimeValue(time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC))
	other := DateTimeValue(now)

	if value.CompareTo(other) != CompareToLessThan {
		t.Fatalf("Expected date/time value to be less than the other but was not")
	}
}

func TestCompareToDateTimeGreaterThan(t *testing.T) {
	now := time.Now()
	value := DateTimeValue(now)
	other := DateTimeValue(time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC))

	if value.CompareTo(other) != CompareToGreaterThan {
		t.Fatalf("Expected date/time value to be greater than the other but was not")
	}
}

func TestCompareToUndefinedValues(t *testing.T) {
	value := EmptyValue
	other := EmptyValue

	if value.CompareTo(other) != CompareToNotPossible {
		t.Fatalf("Expected comparison not possible between empty values but was possible")
	}
}

func TestGetBooleanAsTrue1(t *testing.T) {
	value, _ := booleanValueUsing(true).GetBoolean()

	if value != true {
		t.Fatalf("Expected value to be true but was not")
	}
}

func TestGetBooleanAsTrue2(t *testing.T) {
	value, _ := StringValue("true").GetBoolean()

	if value != true {
		t.Fatalf("Expected value to be true but was not")
	}
}

func TestGetBooleanAsFalse1(t *testing.T) {
	value, _ := booleanValueUsing(false).GetBoolean()

	if value != false {
		t.Fatalf("Expected value to be false but was not")
	}
}

func TestGetBooleanAsFalse2(t *testing.T) {
	value, _ := StringValue("false").GetBoolean()

	if value != false {
		t.Fatalf("Expected value to be false but was not")
	}
}

func TestGetBoolean(t *testing.T) {
	_, err := StringValue("nothing").GetBoolean()

	if err == nil {
		t.Fatalf("Expected an error while trying to get boolean value for nothing but received none")
	}
}

func TestTokenToInt64Value(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.Numeric, "12")
	value, _ := ToValue(token)

	if value.CompareTo(Int64Value(12)) != CompareToEqual {
		t.Fatalf("Expected token %v to be converted to a value %v, but received %v", "12", Int64Value(12), value)
	}
}

func TestTokenToInt64ValueWithError(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.Numeric, "non-numeric")
	_, err := ToValue(token)

	if err == nil {
		t.Fatalf("Expected an error while converting %v to int64 but received none", "non-numeric")
	}
}

func TestTokenToFloat64Value(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.FloatingPoint, "12.78")
	value, _ := ToValue(token)

	if value.CompareTo(Float64Value(12.78)) != CompareToEqual {
		t.Fatalf("Expected token %v to be converted to a value %v, but received %v", "12.78", Float64Value(12.78), value)
	}
}

func TestTokenToFloat64ValueWithError(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.FloatingPoint, "non-numeric")
	_, err := ToValue(token)

	if err == nil {
		t.Fatalf("Expected an error while converting %v to float64 but received none", "non-numeric")
	}
}

func TestTokenToBooleanAsTrue1(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.Boolean, "true")
	value, _ := ToValue(token)

	if value.CompareTo(booleanValueUsing(true)) != CompareToEqual {
		t.Fatalf("Expected token %v to be converted to a value %v, but received %v", "true", booleanValueUsing(true), value)
	}
}

func TestTokenToBooleanAsTrue2(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.Boolean, "y")
	value, _ := ToValue(token)

	if value.CompareTo(booleanValueUsing(true)) != CompareToEqual {
		t.Fatalf("Expected token %v to be converted to a value %v, but received %v", "y", booleanValueUsing(true), value)
	}
}

func TestTokenToBooleanAsFalse1(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.Boolean, "false")
	value, _ := ToValue(token)

	if value.CompareTo(booleanValueUsing(false)) != CompareToEqual {
		t.Fatalf("Expected token %v to be converted to a value %v, but received %v", "false", booleanValueUsing(false), value)
	}
}

func TestTokenToBooleanAsFalse2(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.Boolean, "n")
	value, _ := ToValue(token)

	if value.CompareTo(booleanValueUsing(false)) != CompareToEqual {
		t.Fatalf("Expected token %v to be converted to a value %v, but received %v", "n", booleanValueUsing(false), value)
	}
}

func TestTokenToString1(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.Boolean, "non-boolean")
	value, _ := ToValue(token)

	if value.CompareTo(StringValue("non-boolean")) != CompareToEqual {
		t.Fatalf("Expected token %v to be converted to a value %v, but received %v", "non-boolean", StringValue("non-boolean"), value)
	}
}

func TestTokenToString2(t *testing.T) {
	token := tokenizer.NewToken(tokenizer.RawString, "string")
	value, _ := ToValue(token)

	if value.CompareTo(StringValue("string")) != CompareToEqual {
		t.Fatalf("Expected token %v to be converted to a value %v, but received %v", "string", StringValue("string"), value)
	}
}
