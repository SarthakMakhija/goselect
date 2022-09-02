package context

import "testing"

func TestCompareTo1(t *testing.T) {
	value := IntValue(10)
	other := IntValue(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int values to match but they did not")
	}
}

func TestCompareTo2(t *testing.T) {
	value := IntValue(10)
	other := IntValue(20)

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected int values to not match but they did")
	}
}

func TestCompareTo3(t *testing.T) {
	value := IntValue(10)
	other := StringValue("10")

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and string values to match but they did not")
	}
}

func TestCompareTo4(t *testing.T) {
	value := IntValue(10)
	other := Int64Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and int64 values to match but they did not")
	}
}

func TestCompareTo5(t *testing.T) {
	value := IntValue(10)
	other := Uint32Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and uint32 values to match but they did not")
	}
}

func TestCompareTo6(t *testing.T) {
	value := IntValue(10)
	other := Float64Value(10.0)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected int and float64 values to match but they did not")
	}
}

func TestCompareTo7(t *testing.T) {
	value := Uint32Value(10)
	other := Int64Value(10)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected uint32 and int64 values to match but they did not")
	}
}

func TestCompareTo8(t *testing.T) {
	value := StringValue("3.27")
	other := Float64Value(3.27)

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and float64 values to match but they did not")
	}
}

func TestCompareTo9(t *testing.T) {
	value := BooleanValue(true)
	other := StringValue("Y")

	if value.CompareTo(other) != 0 {
		t.Fatalf("Expected string and boolean values to match but they did not")
	}
}

func TestCompareTo10(t *testing.T) {
	value := BooleanValue(false)
	other := StringValue("Y")

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected string and boolean values to not match but they did")
	}
}

func TestCompareTo11(t *testing.T) {
	value := IntValue(10)
	other := Float64Value(10.1)

	if value.CompareTo(other) == 0 {
		t.Fatalf("Expected int and float64 values to not match but they did")
	}
}
