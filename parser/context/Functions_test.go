package context

import (
	"os"
	"testing"
	"time"
)

func TestLower1(t *testing.T) {
	value, _ := NewFunctions().Execute("lower", StringValue("ABC"))
	expected := "abc"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, value)
	}
}

func TestLower2(t *testing.T) {
	value, _ := NewFunctions().Execute("low", StringValue("ABC"))
	expected := "abc"

	actualValue, _ := value.GetString()
	if actualValue != expected {
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
	value, _ := NewFunctions().Execute("upper", StringValue("abc"))
	expected := "ABC"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected upper to be %v, received %v", expected, value)
	}
}

func TestUpper2(t *testing.T) {
	value, _ := NewFunctions().Execute("up", StringValue("abc"))
	expected := "ABC"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected upper to be %v, received %v", expected, value)
	}
}

func TestUpperWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("up")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestTitle(t *testing.T) {
	value, _ := NewFunctions().Execute("title", StringValue("Sample content"))
	expected := "Sample Content"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected title to be %v, received %v", expected, value)
	}
}

func TestTitleWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("title")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestBase641(t *testing.T) {
	value, _ := NewFunctions().Execute("base64", StringValue("a"))
	expected := "YQ=="

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected base64 to be %v, received %v", expected, value)
	}
}

func TestBase642(t *testing.T) {
	value, _ := NewFunctions().Execute("b64", StringValue("a"))
	expected := "YQ=="

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected base64 to be %v, received %v", expected, value)
	}
}

func TestBase64WithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("b64")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestLength1(t *testing.T) {
	value, _ := NewFunctions().Execute("length", StringValue("sample"))
	expected := 6

	actualValue, _ := value.GetInt()
	if actualValue != expected {
		t.Fatalf("Expected length to be %v, received %v", expected, value)
	}
}

func TestLength2(t *testing.T) {
	value, _ := NewFunctions().Execute("len", StringValue("sample"))
	expected := 6

	actualValue, _ := value.GetInt()
	if actualValue != expected {
		t.Fatalf("Expected length to be %v, received %v", expected, value)
	}
}

func TestLengthWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("len")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestLeftTrim1(t *testing.T) {
	value, _ := NewFunctions().Execute("ltrim", StringValue("  sample"))
	expected := "sample"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected leftTrim to be %v, received %v", expected, value)
	}
}

func TestLeftTrim2(t *testing.T) {
	value, _ := NewFunctions().Execute("lTrim", StringValue("  sample"))
	expected := "sample"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected leftTrim to be %v, received %v", expected, value)
	}
}

func TestLeftTrimWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("ltrim")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestRightTrim1(t *testing.T) {
	value, _ := NewFunctions().Execute("rtrim", StringValue("sample  "))
	expected := "sample"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected rightTrim to be %v, received %v", expected, value)
	}
}

func TestRightTrim2(t *testing.T) {
	value, _ := NewFunctions().Execute("rTrim", StringValue("sample  "))
	expected := "sample"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected rightTrim to be %v, received %v", expected, value)
	}
}

func TestRightTrimWithoutAnyParameter(t *testing.T) {
	_, err := NewFunctions().Execute("rtrim")

	if err == nil {
		t.Fatalf("Expected an error on executing low without any parameter")
	}
}

func TestTrim1(t *testing.T) {
	value, _ := NewFunctions().Execute("trim", StringValue("  sample  "))
	expected := "sample"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected trim to be %v, received %v", expected, value)
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

	actualValue, _ := value.GetInt()
	if actualValue != expected {
		t.Fatalf("Expected day to be %v, received %v", expected, value)
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

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected month to be %v, received %v", expected, value)
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

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected month to be %v, received %v", expected, value)
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

	actualValue, _ := value.GetInt()
	if actualValue != expected {
		t.Fatalf("Expected year to be %v, received %v", expected, value)
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

	actualValue, _ := value.GetInt()
	if actualValue != expected {
		t.Fatalf("Expected year to be %v, received %v", expected, value)
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

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected day of week to be %v, received %v", expected, value)
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

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected day of week to be %v, received %v", expected, value)
	}
}

func TestCurrentWorkingDirectory1(t *testing.T) {
	value, _ := NewFunctions().Execute("cwd")
	expected, _ := os.Getwd()

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected current working directory to be %v, received %v", expected, value)
	}
}

func TestCurrentWorkingDirectory2(t *testing.T) {
	value, _ := NewFunctions().Execute("wd")
	expected, _ := os.Getwd()

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected current working directory to be %v, received %v", expected, value)
	}
}
