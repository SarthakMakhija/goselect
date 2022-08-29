package context

import (
	"testing"
	"time"
)

func TestLower1(t *testing.T) {
	value, _ := NewFunctions().Execute("lower", "ABC")
	expected := "abc"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestLower2(t *testing.T) {
	value, _ := NewFunctions().Execute("low", "ABC")
	expected := "abc"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestLowerWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("low")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestUpper1(t *testing.T) {
	value, _ := NewFunctions().Execute("upper", "abc")
	expected := "ABC"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestUpper2(t *testing.T) {
	value, _ := NewFunctions().Execute("up", "abc")
	expected := "ABC"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestUpperWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("up")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestTitle(t *testing.T) {
	value, _ := NewFunctions().Execute("title", "Sample content")
	expected := "Sample Content"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestTitleWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("title")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestBase641(t *testing.T) {
	value, _ := NewFunctions().Execute("base64", "a")
	expected := "YQ=="

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestBase642(t *testing.T) {
	value, _ := NewFunctions().Execute("b64", "a")
	expected := "YQ=="

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestBase64WithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("b64")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestLength1(t *testing.T) {
	value, _ := NewFunctions().Execute("length", "sample")
	expected := 6

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestLength2(t *testing.T) {
	value, _ := NewFunctions().Execute("len", "sample")
	expected := 6

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestLengthWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("len")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestLeftTrim1(t *testing.T) {
	value, _ := NewFunctions().Execute("ltrim", "  sample")
	expected := "sample"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestLeftTrim2(t *testing.T) {
	value, _ := NewFunctions().Execute("lTrim", "  sample")
	expected := "sample"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestLeftTrimWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("ltrim")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestRightTrim1(t *testing.T) {
	value, _ := NewFunctions().Execute("rtrim", "sample  ")
	expected := "sample"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestRightTrim2(t *testing.T) {
	value, _ := NewFunctions().Execute("rTrim", "sample  ")
	expected := "sample"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestRightTrimWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("rtrim")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestTrim1(t *testing.T) {
	value, _ := NewFunctions().Execute("trim", "  sample  ")
	expected := "sample"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestTrimWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("trim")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestDay(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("day")
	expected := 22

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestMonth1(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("month")
	expected := "August"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestMonth2(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("mon")
	expected := "August"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestYear1(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("year")
	expected := 2022

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestYear2(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("yr")
	expected := 2022

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestDayOfWeek1(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("dayOfWeek")
	expected := "Sunday"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestDayOfWeek2(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("dayofweek")
	expected := "Sunday"

	if value != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}
