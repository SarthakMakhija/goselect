package context

import (
	"testing"
)

func TestCompareIntInt(t *testing.T) {
	value := IntValue(10)
	other := IntValue(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int values to match but they did not")
	}
}

func TestCompareIntInt2(t *testing.T) {
	value := IntValue(10)
	other := IntValue(20)

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected int values to not match but they did")
	}
}

func TestCompareIntString(t *testing.T) {
	value := IntValue(10)
	other := StringValue("10")

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and string values to match but they did not")
	}
}

func TestCompareIntInt64(t *testing.T) {
	value := IntValue(10)
	other := Int64Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and int64 values to match but they did not")
	}
}

func TestCompareInt64Int(t *testing.T) {
	value := Int64Value(10)
	other := IntValue(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and int64 values to match but they did not")
	}
}

func TestCompareToIntUint32(t *testing.T) {
	value := IntValue(10)
	other := Uint32Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and uint32 values to match but they did not")
	}
}

func TestCompareToUint32Int(t *testing.T) {
	value := Uint32Value(10)
	other := IntValue(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and uint32 values to match but they did not")
	}
}

func TestCompareToIntFloat(t *testing.T) {
	value := IntValue(10)
	other := Float64Value(10.0)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and float64 values to match but they did not")
	}
}

func TestCompareToFloatInt(t *testing.T) {
	value := Float64Value(10)
	other := IntValue(10.0)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and float64 values to match but they did not")
	}
}

func TestCompareToUint32Float64(t *testing.T) {
	value := Uint32Value(10)
	other := Float64Value(10.0)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected uin32 and float64 values to match but they did not")
	}
}

func TestCompareToFloat64Uint32(t *testing.T) {
	value := Float64Value(10.0)
	other := Uint32Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected uin32 and float64 values to match but they did not")
	}
}

func TestCompareToInt64Float64(t *testing.T) {
	value := Int64Value(10)
	other := Float64Value(10.0)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int64 and float64 values to match but they did not")
	}
}

func TestCompareToUint32Int64(t *testing.T) {
	value := Uint32Value(10)
	other := Int64Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected uint32 and int64 values to match but they did not")
	}
}

func TestCompareToInt64Uint32(t *testing.T) {
	value := Int64Value(10)
	other := Uint32Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected uint32 and int64 values to match but they did not")
	}
}

func TestCompareToStringFloat64(t *testing.T) {
	value := StringValue("3.27")
	other := Float64Value(3.27)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and float64 values to match but they did not")
	}
}

func TestCompareToTrueBooleanString(t *testing.T) {
	value := trueBooleanValue
	other := StringValue("Y")

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and boolean values to match but they did not")
	}
}

func TestCompareToFalseBooleanString1(t *testing.T) {
	value := falseBooleanValue
	other := StringValue("Y")

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToFalseBooleanString2(t *testing.T) {
	value := falseBooleanValue
	other := StringValue("n")

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and boolean values match but they did not")
	}
}

func TestCompareToFalseBooleanInvalidString(t *testing.T) {
	value := falseBooleanValue
	other := StringValue("ok")

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToInvalidStringTrueBoolean(t *testing.T) {
	value := StringValue("ok")
	other := trueBooleanValue

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToStringTrueBoolean(t *testing.T) {
	value := StringValue("n")
	other := trueBooleanValue

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToStringFalseBoolean(t *testing.T) {
	value := StringValue("Y")
	other := falseBooleanValue

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareToStringInt1(t *testing.T) {
	value := StringValue("10")
	other := IntValue(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and int values to match but they did not")
	}
}

func TestCompareToStringInt2(t *testing.T) {
	value := StringValue("a")
	other := IntValue(10)

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and int values to not match but they did")
	}
}

func TestCompareToStringInt641(t *testing.T) {
	value := StringValue("10")
	other := Int64Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and int64 values to match but they did not")
	}
}

func TestCompareToStringInt642(t *testing.T) {
	value := StringValue("a")
	other := Int64Value(10)

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and int64 values to not match but they did")
	}
}

func TestCompareToStringUint321(t *testing.T) {
	value := StringValue("10")
	other := Uint32Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and uint32 values to match but they did not")
	}
}

func TestCompareToStringUint322(t *testing.T) {
	value := StringValue("a")
	other := Uint32Value(10)

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and uint32 values to not match but they did")
	}
}

func TestCompareToIntFloat64(t *testing.T) {
	value := IntValue(10)
	other := Float64Value(10.1)

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected int and float64 values to not match but they did")
	}
}

func TestCompareToIntString(t *testing.T) {
	value := IntValue(10)
	other := StringValue("content")

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected int and string values to not match but they did")
	}
}

func TestCompareToFloat64String1(t *testing.T) {
	value := Float64Value(10)
	other := StringValue("content")

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected float64 and string values to not match but they did")
	}
}

func TestCompareToFloat64String2(t *testing.T) {
	value := Float64Value(10.2)
	other := StringValue("10.2")

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected float64 and string values to match but they did not")
	}
}

func TestCompareToFloat64Int1(t *testing.T) {
	value := Float64Value(10.1)
	other := IntValue(10)

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected int and float64 values to not match but they did")
	}
}

func TestCompareToFloat64Int2(t *testing.T) {
	value := Float64Value(10)
	other := IntValue(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and float64 values to match but they did not")
	}
}

func TestCompareToInt64String1(t *testing.T) {
	value := Int64Value(3)
	other := StringValue("3")

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and int64 values to match but they did not")
	}
}

func TestCompareToInt64String2(t *testing.T) {
	value := Int64Value(3)
	other := StringValue("a")

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and int64 values to not match but they did")
	}
}

func TestCompareToUint32String1(t *testing.T) {
	value := Uint32Value(3)
	other := StringValue("3")

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and uint32 values to match but they did not")
	}
}

func TestCompareToUint32String2(t *testing.T) {
	value := Uint32Value(3)
	other := StringValue("a")

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and uint32 values to not match but they did")
	}
}

func TestCompareToFloat64Int64(t *testing.T) {
	value := Float64Value(3.0)
	other := Int64Value(3)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected float64 and int64 values to match but they did not")
	}
}

func TestGetCommonType1(t *testing.T) {
	value := Float64Value(3.0)
	other := Int64Value(3)

	first, second, _ := getCommonType(value, other, TypePair{aType: ValueTypeFloat64, bType: ValueTypeInt64})
	if first.CompareTo(second) != 0 {
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
