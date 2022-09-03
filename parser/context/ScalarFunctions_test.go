package context

import (
	"os"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	value, _ := NewFunctions().Execute("add", StringValue("1"), IntValue(2), IntValue(4))
	var expected float64 = 7

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue != expected {
		t.Fatalf("Expected addition to be %v, received %v", expected, actualValue)
	}
}

func TestSubtract(t *testing.T) {
	value, _ := NewFunctions().Execute("sub", StringValue("1"), IntValue(2))
	var expected float64 = -1

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue != expected {
		t.Fatalf("Expected subtraction to be %v, received %v", expected, actualValue)
	}
}

func TestMultiply(t *testing.T) {
	value, _ := NewFunctions().Execute("mul", StringValue("6"), IntValue(8))
	var expected float64 = 48

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue != expected {
		t.Fatalf("Expected multiplication to be %v, received %v", expected, actualValue)
	}
}

func TestDivision(t *testing.T) {
	value, _ := NewFunctions().Execute("div", StringValue("9"), IntValue(2))
	var expected = 4.5

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue != expected {
		t.Fatalf("Expected division to be %v, received %v", expected, actualValue)
	}
}

func TestDivisionFailure(t *testing.T) {
	_, err := NewFunctions().Execute("div", StringValue("9"), IntValue(0))

	if err == nil {
		t.Fatalf("Expected an error while dividing with zero but received none")
	}
}

func TestDivisionWithNegative(t *testing.T) {
	value, _ := NewFunctions().Execute("div", StringValue("9"), IntValue(-2))
	var expected = -4.5

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue != expected {
		t.Fatalf("Expected division to be %v, received %v", expected, actualValue)
	}
}

func TestEqualsReturningTrue(t *testing.T) {
	value, _ := NewFunctions().Execute("eq", StringValue("one"), StringValue("one"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected equals to be %v, received %v", true, actualValue)
	}
}

func TestEqualsReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("eq", StringValue("one"), StringValue("another"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected equals to be %v, received %v", false, actualValue)
	}
}

func TestLessThanReturningTrue1(t *testing.T) {
	value, _ := NewFunctions().Execute("lt", StringValue("one"), StringValue("two"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected less than to be %v, received %v", true, actualValue)
	}
}

func TestLessThanReturningTrue2(t *testing.T) {
	value, _ := NewFunctions().Execute("lt", StringValue("1"), Float64Value(2))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected less than to be %v, received %v", true, actualValue)
	}
}

func TestLessThanReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("lt", StringValue("3"), Float64Value(2))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected less than to be %v, received %v", false, actualValue)
	}
}

func TestGreaterThanReturningTrue1(t *testing.T) {
	value, _ := NewFunctions().Execute("gt", StringValue("two"), StringValue("one"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected greater than to be %v, received %v", true, actualValue)
	}
}

func TestGreaterThanReturningTrue2(t *testing.T) {
	value, _ := NewFunctions().Execute("gt", StringValue("2"), Float64Value(1))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected greater than to be %v, received %v", true, actualValue)
	}
}

func TestGreaterThanReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("gt", StringValue("2"), Float64Value(3))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected greater than to be %v, received %v", false, actualValue)
	}
}

func TestLessThanEqualReturningTrue1(t *testing.T) {
	value, _ := NewFunctions().Execute("lte", StringValue("one"), StringValue("two"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected less than equal to be %v, received %v", true, actualValue)
	}
}

func TestLessThanEqualReturningTrue2(t *testing.T) {
	value, _ := NewFunctions().Execute("le", StringValue("1"), Float64Value(1))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected less than equal to be %v, received %v", true, actualValue)
	}
}

func TestLessThanEqualReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("le", StringValue("5"), Float64Value(4))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected less than equal to be %v, received %v", false, actualValue)
	}
}

func TestGreaterThanEqualReturningTrue1(t *testing.T) {
	value, _ := NewFunctions().Execute("gte", StringValue("two"), StringValue("one"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected greater than equal to be %v, received %v", true, actualValue)
	}
}

func TestGreaterThanEqualReturningTrue2(t *testing.T) {
	value, _ := NewFunctions().Execute("ge", StringValue("1"), Float64Value(1))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected greater than equal to be %v, received %v", true, actualValue)
	}
}

func TestGreaterThanEqualReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("ge", StringValue("4"), Float64Value(5))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected greater than equal to be %v, received %v", false, actualValue)
	}
}

func TestOrReturningTrue(t *testing.T) {
	value, _ := NewFunctions().Execute("or", BooleanValue(false), StringValue("n"), BooleanValue(true))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected or to be %v, received %v", true, actualValue)
	}
}

func TestOrReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("or", BooleanValue(false), StringValue("n"), BooleanValue(false))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected or to be %v, received %v", false, actualValue)
	}
}

func TestAndReturningTrue(t *testing.T) {
	value, _ := NewFunctions().Execute("and", BooleanValue(true), StringValue("y"), BooleanValue(true))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected and to be %v, received %v", true, actualValue)
	}
}

func TestAndReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("and", BooleanValue(true), StringValue("y"), BooleanValue(false))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected and to be %v, received %v", false, actualValue)
	}
}

func TestLower1(t *testing.T) {
	value, _ := NewFunctions().Execute("lower", StringValue("ABC"))
	expected := "abc"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, actualValue)
	}
}

func TestLower2(t *testing.T) {
	value, _ := NewFunctions().Execute("low", StringValue("ABC"))
	expected := "abc"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected upper to be %v, received %v", expected, actualValue)
	}
}

func TestUpper2(t *testing.T) {
	value, _ := NewFunctions().Execute("up", StringValue("abc"))
	expected := "ABC"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected upper to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected title to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected base64 to be %v, received %v", expected, actualValue)
	}
}

func TestBase642(t *testing.T) {
	value, _ := NewFunctions().Execute("b64", StringValue("a"))
	expected := "YQ=="

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected base64 to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected length to be %v, received %v", expected, actualValue)
	}
}

func TestLength2(t *testing.T) {
	value, _ := NewFunctions().Execute("len", StringValue("sample"))
	expected := 6

	actualValue, _ := value.GetInt()
	if actualValue != expected {
		t.Fatalf("Expected length to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected leftTrim to be %v, received %v", expected, actualValue)
	}
}

func TestLeftTrim2(t *testing.T) {
	value, _ := NewFunctions().Execute("lTrim", StringValue("  sample"))
	expected := "sample"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected leftTrim to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected rightTrim to be %v, received %v", expected, actualValue)
	}
}

func TestRightTrim2(t *testing.T) {
	value, _ := NewFunctions().Execute("rTrim", StringValue("sample  "))
	expected := "sample"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected rightTrim to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected trim to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected day to be %v, received %v", expected, actualValue)
	}
}

func TestDate(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("date")
	expected := "2022-August-22"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected date to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected month to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected month to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected year to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected year to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected day of week to be %v, received %v", expected, actualValue)
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
		t.Fatalf("Expected day of week to be %v, received %v", expected, actualValue)
	}
}

func TestCurrentWorkingDirectory1(t *testing.T) {
	value, _ := NewFunctions().Execute("cwd")
	expected, _ := os.Getwd()

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected current working directory to be %v, received %v", expected, actualValue)
	}
}

func TestCurrentWorkingDirectory2(t *testing.T) {
	value, _ := NewFunctions().Execute("wd")
	expected, _ := os.Getwd()

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected current working directory to be %v, received %v", expected, actualValue)
	}
}

func TestConcat(t *testing.T) {
	value, _ := NewFunctions().Execute("concat", StringValue("a"), StringValue("b"))
	expected := "ab"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected concat  to be %v, received %v", expected, actualValue)
	}
}

func TestConcatWithSeparator(t *testing.T) {
	value, _ := NewFunctions().Execute("concatws", StringValue("a"), StringValue("b"), StringValue("@"))
	expected := "a@b"

	actualValue, _ := value.GetString()
	if actualValue != expected {
		t.Fatalf("Expected concat  to be %v, received %v", expected, actualValue)
	}
}

func TestContains(t *testing.T) {
	value, _ := NewFunctions().Execute("contains", StringValue("String"), StringValue("ing"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected concat  to be %v, received %v", true, actualValue)
	}
}

func TestContainsWithInsufficientParameters(t *testing.T) {
	_, err := NewFunctions().Execute("contains", StringValue("String"))

	if err == nil {
		t.Fatalf("Expected an error on executing contains with insufficient parameters")
	}
}

func TestSubstringWithBeginIndexOnly(t *testing.T) {
	value, _ := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("2"))

	actualValue, _ := value.GetString()
	if actualValue != "cdef" {
		t.Fatalf("Expected substring to be %v, received %v", "cdef", actualValue)
	}
}

func TestSubstringWithBeginAndEndIndex(t *testing.T) {
	value, _ := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("2"), StringValue("4"))

	actualValue, _ := value.GetString()
	if actualValue != "cde" {
		t.Fatalf("Expected substring to be %v, received %v", "cde", actualValue)
	}
}

func TestSubstringWithEndIndexGreaterThanLength(t *testing.T) {
	value, _ := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("2"), StringValue("100"))

	actualValue, _ := value.GetString()
	if actualValue != "cdef" {
		t.Fatalf("Expected substring to be %v, received %v", "cdef", actualValue)
	}
}

func TestSubstringWithBeginIndexGreaterThanLength(t *testing.T) {
	value, _ := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("100"), StringValue("6"))

	actualValue, _ := value.GetString()
	if actualValue != "abcdef" {
		t.Fatalf("Expected substring to be %v, received %v", "abcdef", actualValue)
	}
}

func TestSubstringWithInsufficientParameters(t *testing.T) {
	_, err := NewFunctions().Execute("substr", StringValue("abcdef"))

	if err == nil {
		t.Fatalf("Expected an error on executing substring with insufficient parameters")
	}
}

func TestSubstringWithToLessThanFrom(t *testing.T) {
	_, err := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("3"), StringValue("0"))

	if err == nil {
		t.Fatalf("Expected an error on executing substring with to less than from")
	}
}

func TestSubstringWithNegativeFrom(t *testing.T) {
	_, err := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("-3"))

	if err == nil {
		t.Fatalf("Expected an error on executing substring with negative from")
	}
}

func TestSubstringWithNegativeTo(t *testing.T) {
	_, err := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("0"), StringValue("-5"))

	if err == nil {
		t.Fatalf("Expected an error on executing substring with negative to")
	}
}

func TestSubstringWithIllegalFrom(t *testing.T) {
	_, err := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("illegal"))

	if err == nil {
		t.Fatalf("Expected an error on executing substring with illegal from")
	}
}

func TestSubstringWithIllegalTo(t *testing.T) {
	_, err := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("0"), StringValue("illegal"))

	if err == nil {
		t.Fatalf("Expected an error on executing substring with illegal to")
	}
}
