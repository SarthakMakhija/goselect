//go:build unit
// +build unit

package context

import (
	"math"
	"os"
	"testing"
	"time"
)

func TestIdentity(t *testing.T) {
	value, _ := NewFunctions().Execute("identity", StringValue("100"))
	var expected = "100"

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected identity to be %v, received %v", expected, actualValue)
	}
}

func TestIdentityWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("identity")

	if err == nil {
		t.Fatalf("Expected an error while executing identify with no parameter value")
	}
}

func TestAddWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("add")

	if err == nil {
		t.Fatalf("Expected an error while executing add with no parameter value")
	}
}

func TestAddWithNonNumericParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("add", StringValue("1"), StringValue("a"))

	if err == nil {
		t.Fatalf("Expected an error while executing add with non-numeric parameter value")
	}
}

func TestAdd(t *testing.T) {
	value, _ := NewFunctions().Execute("add", StringValue("1"), IntValue(2), IntValue(4))
	var expected float64 = 7

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue != expected {
		t.Fatalf("Expected addition to be %v, received %v", expected, actualValue)
	}
}

func TestSubtractWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("sub")

	if err == nil {
		t.Fatalf("Expected an error while executing sub with no parameter value")
	}
}

func TestSubtractWithNonNumericParameterValue1(t *testing.T) {
	_, err := NewFunctions().Execute("sub", StringValue("a"), StringValue("1"))

	if err == nil {
		t.Fatalf("Expected an error while executing sub with non-numeric parameter value")
	}
}

func TestSubtractWithNonNumericParameterValue2(t *testing.T) {
	_, err := NewFunctions().Execute("sub", StringValue("1"), StringValue("a"))

	if err == nil {
		t.Fatalf("Expected an error while executing sub with non-numeric parameter value")
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

func TestMultiplyWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("mul")

	if err == nil {
		t.Fatalf("Expected an error while executing mul with no parameter value")
	}
}

func TestMultiplyWithNonNumericParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("mul", StringValue(""), StringValue("a"))

	if err == nil {
		t.Fatalf("Expected an error while executing mul with non-numeric parameter value")
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

func TestDivisionWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("div")

	if err == nil {
		t.Fatalf("Expected an error while executing div with no parameter value")
	}
}

func TestDivisionWithNonNumericParameterValue1(t *testing.T) {
	_, err := NewFunctions().Execute("div", StringValue("a"), StringValue("1"))

	if err == nil {
		t.Fatalf("Expected an error while executing div with non-numeric parameter value")
	}
}

func TestDivisionWithNonNumericParameterValue2(t *testing.T) {
	_, err := NewFunctions().Execute("div", StringValue("1"), StringValue("a"))

	if err == nil {
		t.Fatalf("Expected an error while executing div with non-numeric parameter value")
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

func TestEqualsWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("eq")

	if err == nil {
		t.Fatalf("Expected an error while executing eq with no parameter value")
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

func TestNotEqualWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("ne")

	if err == nil {
		t.Fatalf("Expected an error while executing ne with no parameter value")
	}
}

func TestNotEqualReturningTrue(t *testing.T) {
	value, _ := NewFunctions().Execute("ne", StringValue("one"), StringValue("two"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected not equal to be %v, received %v", true, actualValue)
	}
}

func TestNotEqualReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("ne", StringValue("one"), StringValue("one"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected not equal to be %v, received %v", false, actualValue)
	}
}

func TestLessThanWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("lt")

	if err == nil {
		t.Fatalf("Expected an error while executing less than with no parameter value")
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

func TestGreaterThanWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("gt")

	if err == nil {
		t.Fatalf("Expected an error while executing gt with no parameter value")
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

func TestLessThanEqualWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("lte")

	if err == nil {
		t.Fatalf("Expected an error while executing lte with no parameter value")
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

func TestGreaterThanEqualWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("gte")

	if err == nil {
		t.Fatalf("Expected an error while executing gte with no parameter value")
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

func TestOrWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("or")

	if err == nil {
		t.Fatalf("Expected an error while executing or with no parameter value")
	}
}

func TestOrWithIllegalParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("or", BooleanValue(false), IntValue(5))

	if err == nil {
		t.Fatalf("Expected an error while executing or with illegal parameter value")
	}
}

func TestOrReturningTrue(t *testing.T) {
	value, _ := NewFunctions().Execute("or", falseBooleanValue, StringValue("n"), trueBooleanValue)

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected or to be %v, received %v", true, actualValue)
	}
}

func TestOrReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("or", falseBooleanValue, StringValue("n"), falseBooleanValue)

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected or to be %v, received %v", false, actualValue)
	}
}

func TestAndWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("and")

	if err == nil {
		t.Fatalf("Expected an error while executing and with no parameter value")
	}
}

func TestAndWithIllegalParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("and", BooleanValue(true), IntValue(5))

	if err == nil {
		t.Fatalf("Expected an error while executing and with illegal parameter value")
	}
}

func TestAndReturningTrue(t *testing.T) {
	value, _ := NewFunctions().Execute("and", trueBooleanValue, StringValue("y"), trueBooleanValue)

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected and to be %v, received %v", true, actualValue)
	}
}

func TestAndReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("and", StringValue("y"), falseBooleanValue)

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected and to be %v, received %v", false, actualValue)
	}
}

func TestNotWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("not")

	if err == nil {
		t.Fatalf("Expected an error while executing not with no parameter value")
	}
}

func TestNotWithIllegalParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("not", IntValue(5))

	if err == nil {
		t.Fatalf("Expected an error while executing not with illegal parameter value")
	}
}

func TestNotReturningTrue(t *testing.T) {
	value, _ := NewFunctions().Execute("not", falseBooleanValue)

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected not to be %v, received %v", true, actualValue)
	}
}

func TestNotReturningFalse(t *testing.T) {
	value, _ := NewFunctions().Execute("not", trueBooleanValue)

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected not to be %v, received %v", false, actualValue)
	}
}

func TestLower1(t *testing.T) {
	value, _ := NewFunctions().Execute("lower", StringValue("ABC"))
	expected := "abc"

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected lower to be %v, received %v", expected, actualValue)
	}
}

func TestLower2(t *testing.T) {
	value, _ := NewFunctions().Execute("low", StringValue("ABC"))
	expected := "abc"

	actualValue := value.GetAsString()
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

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected upper to be %v, received %v", expected, actualValue)
	}
}

func TestUpper2(t *testing.T) {
	value, _ := NewFunctions().Execute("up", StringValue("abc"))
	expected := "ABC"

	actualValue := value.GetAsString()
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

	actualValue := value.GetAsString()
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

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected base64 to be %v, received %v", expected, actualValue)
	}
}

func TestBase642(t *testing.T) {
	value, _ := NewFunctions().Execute("b64", StringValue("a"))
	expected := "YQ=="

	actualValue := value.GetAsString()
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

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected leftTrim to be %v, received %v", expected, actualValue)
	}
}

func TestLeftTrim2(t *testing.T) {
	value, _ := NewFunctions().Execute("lTrim", StringValue("  sample"))
	expected := "sample"

	actualValue := value.GetAsString()
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

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected rightTrim to be %v, received %v", expected, actualValue)
	}
}

func TestRightTrim2(t *testing.T) {
	value, _ := NewFunctions().Execute("rTrim", StringValue("sample  "))
	expected := "sample"

	actualValue := value.GetAsString()
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

	actualValue := value.GetAsString()
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

func TestNow(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()
	value, _ := NewFunctions().Execute("now")
	expected := nowFunc()

	if !value.timeValue.Equal(expected) {
		t.Fatalf("Expected now to return %v, received %v", expected, value.timeValue)
	}
}

func TestNowAsString(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()
	value, _ := NewFunctions().Execute("now")
	expected := "2022-08-22 15:08:00 +0000 UTC"

	if value.GetAsString() != expected {
		t.Fatalf("Expected now to return %v, received %v", expected, value.GetAsString())
	}
}

func TestDay(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 22, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("cday")
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

	value, _ := NewFunctions().Execute("cdate")
	expected := "2022-August-22"

	actualValue := value.GetAsString()
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

	value, _ := NewFunctions().Execute("cmonth")
	expected := "August"

	actualValue := value.GetAsString()
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

	value, _ := NewFunctions().Execute("cmon")
	expected := "August"

	actualValue := value.GetAsString()
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

	value, _ := NewFunctions().Execute("cyear")
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

	value, _ := NewFunctions().Execute("cyr")
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

	actualValue := value.GetAsString()
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

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected day of week to be %v, received %v", expected, actualValue)
	}
}

func TestExtractWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("extract")

	if err == nil {
		t.Fatalf("Expected an error while executing extract with no parameter value")
	}
}

func TestExtractWithIllegalParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("extract", BooleanValue(false), IntValue(5))

	if err == nil {
		t.Fatalf("Expected an error while executing extract with illegal parameter value")
	}
}

func TestExtractWithInvalidExtractionKey(t *testing.T) {
	_, err := NewFunctions().Execute("extract", DateTimeValue(now()), StringValue("unknown"))

	if err == nil {
		t.Fatalf("Expected an error while executing extract with invalid extraction key")
	}
}

func TestExtractDay(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("extract", DateTimeValue(now()), StringValue("day"))
	expected := "28"

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected extract with day to be %v, received %v", expected, actualValue)
	}
}

func TestExtractYear(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("extract", DateTimeValue(now()), StringValue("year"))
	expected := "2022"

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected extract with year to be %v, received %v", expected, actualValue)
	}
}

func TestExtractMonth(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("extract", DateTimeValue(now()), StringValue("month"))
	expected := "August"

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected extract with month to be %v, received %v", expected, actualValue)
	}
}

func TestExtractWeekDay(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("extract", DateTimeValue(now()), StringValue("weekday"))
	expected := "Sunday"

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected extract with week day to be %v, received %v", expected, actualValue)
	}
}

func TestExtractDate(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("extract", DateTimeValue(now()), StringValue("date"))
	expected := "2022-August-28"

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected extract with date to be %v, received %v", expected, actualValue)
	}
}

func TestFormatDate1(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 5, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value := formatDate(now())
	expected := "2022-August-05"

	if value.GetAsString() != expected {
		t.Fatalf("Expected date to be %v, received %v", expected, value.GetAsString())
	}
}

func TestFormatDate2(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 26, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value := formatDate(now())
	expected := "2022-August-26"

	if value.GetAsString() != expected {
		t.Fatalf("Expected date to be %v, received %v", expected, value.GetAsString())
	}
}

func TestHoursDifferenceWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("hoursdiff")

	if err == nil {
		t.Fatalf("Expected an error while executing hoursdiff with no parameter value")
	}
}

func TestHoursDifferenceWithIllegalParameterValue1(t *testing.T) {
	_, err := NewFunctions().Execute("hoursdiff", BooleanValue(false), IntValue(5))

	if err == nil {
		t.Fatalf("Expected an error while executing hoursdiff with illegal parameter value")
	}
}

func TestHoursDifferenceWithIllegalParameterValue2(t *testing.T) {
	_, err := NewFunctions().Execute("hoursdiff", DateTimeValue(now()), IntValue(5))

	if err == nil {
		t.Fatalf("Expected an error while executing hoursdiff with illegal parameter value")
	}
}

func TestHoursDifference1(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("hoursdiff",
		DateTimeValue(
			time.Date(2022, 8, 28, 14, 8, 00, 0, time.UTC),
		),
		DateTimeValue(
			now(),
		),
	)
	expected := 1.00

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue != expected {
		t.Fatalf("Expected hours difference to be %v, received %v", expected, actualValue)
	}
}

func TestHoursDifference2(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 15, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("hoursdiff",
		DateTimeValue(
			time.Date(2022, 8, 28, 12, 45, 00, 0, time.UTC),
		),
		DateTimeValue(
			now(),
		),
	)
	expected := 2.5

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue != expected {
		t.Fatalf("Expected hours difference to be %v, received %v", expected, actualValue)
	}
}

func TestDaysDifferenceWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("daysdiff")

	if err == nil {
		t.Fatalf("Expected an error while executing daysdiff with no parameter value")
	}
}

func TestDaysDifferenceWithIllegalParameterValue1(t *testing.T) {
	_, err := NewFunctions().Execute("daysdiff", BooleanValue(false), IntValue(5))

	if err == nil {
		t.Fatalf("Expected an error while executing daysdiff with illegal parameter value")
	}
}

func TestDaysDifferenceWithIllegalParameterValue2(t *testing.T) {
	_, err := NewFunctions().Execute("daysdiff", DateTimeValue(now()), IntValue(5))

	if err == nil {
		t.Fatalf("Expected an error while executing daysdiff with illegal parameter value")
	}
}

func TestDaysDifference1(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 28, 15, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("daysdiff",
		DateTimeValue(
			time.Date(2022, 8, 24, 15, 8, 00, 0, time.UTC),
		),
		DateTimeValue(
			now(),
		),
	)
	expected := 4.00

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue != expected {
		t.Fatalf("Expected days difference to be %v, received %v", expected, actualValue)
	}
}

func TestDaysDifference2(t *testing.T) {
	nowFunc = func() time.Time {
		return time.Date(2022, 8, 24, 18, 8, 00, 0, time.UTC)
	}
	// after finish with the test, reset the time implementation
	defer resetClock()

	value, _ := NewFunctions().Execute("daysdiff",
		DateTimeValue(
			time.Date(2022, 8, 24, 15, 8, 00, 0, time.UTC),
		),
		DateTimeValue(
			now(),
		),
	)

	actualValue, _ := value.GetNumericAsFloat64()
	if actualValue > 1 {
		t.Fatalf("Expected days difference to be less than 1 but received %v", actualValue)
	}
}

func TestDaysDifference3(t *testing.T) {
	value, _ := NewFunctions().Execute("daysdiff",
		DateTimeValue(now()),
		DateTimeValue(now()),
	)

	actualValue, _ := value.GetNumericAsFloat64()
	if math.Round(actualValue) != 0 {
		t.Fatalf("Expected days difference to be 0 but received %v, and round resulted in %v", actualValue, math.Round(actualValue))
	}
}

func TestDateTimeParseWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("parsedttm")

	if err == nil {
		t.Fatalf("Expected an error while executing parsedttm with missing parameter value")
	}
}

func TestDateTimeParse1(t *testing.T) {
	value, _ := NewFunctions().Execute("parsedatetime", StringValue("2022-09-28"), StringValue("dt"))
	expected := "2022-September-28"
	actual := formatDate(value.timeValue)

	if expected != actual.GetAsString() {
		t.Fatalf("Expected parsedatetime to return %v, received %v", expected, actual.GetAsString())
	}
}

func TestDateTimeParse2(t *testing.T) {
	_, err := NewFunctions().Execute("parsedatetime", StringValue("2022-28-12"), StringValue("dt"))

	if err == nil {
		t.Fatalf("Expected an error while parsing a date with invalid month but received none")
	}
}

func TestCurrentWorkingDirectory1(t *testing.T) {
	value, _ := NewFunctions().Execute("cwd")
	expected, _ := os.Getwd()

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected current working directory to be %v, received %v", expected, actualValue)
	}
}

func TestCurrentWorkingDirectory2(t *testing.T) {
	value, _ := NewFunctions().Execute("wd")
	expected, _ := os.Getwd()

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected current working directory to be %v, received %v", expected, actualValue)
	}
}

func TestConcatWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("concat")

	if err == nil {
		t.Fatalf("Expected an error while executing concat with no parameter value")
	}
}

func TestConcat(t *testing.T) {
	value, _ := NewFunctions().Execute("concat", StringValue("a"), StringValue("b"))
	expected := "ab"

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected concat  to be %v, received %v", expected, actualValue)
	}
}

func TestConcatWithSeparatorWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("concatws")

	if err == nil {
		t.Fatalf("Expected an error while executing concatws with no parameter value")
	}
}

func TestConcatWithSeparator(t *testing.T) {
	value, _ := NewFunctions().Execute("concatws", StringValue("a"), StringValue("b"), StringValue("@"))
	expected := "a@b"

	actualValue := value.GetAsString()
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

func TestReplace1(t *testing.T) {
	value, _ := NewFunctions().Execute("replace", StringValue("sample.log"), StringValue("log"), StringValue("txt"))

	actualValue := value.GetAsString()
	if actualValue != "sample.txt" {
		t.Fatalf("Expected replace to be %v, received %v", "sample.txt", actualValue)
	}
}

func TestReplace2(t *testing.T) {
	value, _ := NewFunctions().Execute("replace", StringValue("sample.log.log"), StringValue("log"), StringValue("txt"))

	actualValue := value.GetAsString()
	if actualValue != "sample.txt.log" {
		t.Fatalf("Expected replace to be %v, received %v", "sample.txt.log", actualValue)
	}
}

func TestReplace3(t *testing.T) {
	value, _ := NewFunctions().Execute("replace", StringValue("sample.log.log"), StringValue("log"), StringValue("12.39"))

	actualValue := value.GetAsString()
	if actualValue != "sample.12.39.log" {
		t.Fatalf("Expected replace to be %v, received %v", "sample.12.39.log", actualValue)
	}
}

func TestReplace4(t *testing.T) {
	_, err := NewFunctions().Execute("replace", StringValue("sample.log.log"), StringValue("log"))

	if err == nil {
		t.Fatalf("Expected an error with replace given insufficient parameters")
	}
}

func TestReplaceAll1(t *testing.T) {
	value, _ := NewFunctions().Execute("replaceall", StringValue("sample.log.log"), StringValue("log"), StringValue("txt"))

	actualValue := value.GetAsString()
	if actualValue != "sample.txt.txt" {
		t.Fatalf("Expected replace all to be %v, received %v", "sample.txt.txt", actualValue)
	}
}

func TestReplaceAll2(t *testing.T) {
	value, _ := NewFunctions().Execute("replaceall", StringValue("sample.log.log"), StringValue("log"), StringValue("12.39"))

	actualValue := value.GetAsString()
	if actualValue != "sample.12.39.12.39" {
		t.Fatalf("Expected replace all to be %v, received %v", "sample.12.39.12.39", actualValue)
	}
}

func TestReplaceAll3(t *testing.T) {
	_, err := NewFunctions().Execute("replaceall", StringValue("sample.log.log"), StringValue("log"))

	if err == nil {
		t.Fatalf("Expected an error with replace all given insufficient parameters")
	}
}

func TestSubstringWithBeginIndexOnly(t *testing.T) {
	value, _ := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("2"))

	actualValue := value.GetAsString()
	if actualValue != "cdef" {
		t.Fatalf("Expected substring to be %v, received %v", "cdef", actualValue)
	}
}

func TestSubstringWithBeginAndEndIndex(t *testing.T) {
	value, _ := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("2"), StringValue("4"))

	actualValue := value.GetAsString()
	if actualValue != "cde" {
		t.Fatalf("Expected substring to be %v, received %v", "cde", actualValue)
	}
}

func TestSubstringWithEndIndexGreaterThanLength(t *testing.T) {
	value, _ := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("2"), StringValue("100"))

	actualValue := value.GetAsString()
	if actualValue != "cdef" {
		t.Fatalf("Expected substring to be %v, received %v", "cdef", actualValue)
	}
}

func TestSubstringWithBeginIndexGreaterThanLength(t *testing.T) {
	value, _ := NewFunctions().Execute("substr", StringValue("abcdef"), StringValue("100"), StringValue("6"))

	actualValue := value.GetAsString()
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

func TestIsText1(t *testing.T) {
	value, _ := NewFunctions().Execute("istext", StringValue("text/plain"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected istext to be %v, received %v", true, actualValue)
	}
}

func TestIsText2(t *testing.T) {
	value, _ := NewFunctions().Execute("istext", StringValue("text/plain; charset=utf-8"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected istext to be %v, received %v", true, actualValue)
	}
}

func TestIsText3(t *testing.T) {
	value, _ := NewFunctions().Execute("istext", StringValue("image/png"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected istext to be %v, received %v", false, actualValue)
	}
}

func TestIsText4(t *testing.T) {
	value, _ := NewFunctions().Execute("istext")

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected istext to be %v, received %v", false, actualValue)
	}
}

func TestIsImage1(t *testing.T) {
	value, _ := NewFunctions().Execute("isimage", StringValue("image/png"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected isimage to be %v, received %v", true, actualValue)
	}
}

func TestIsImage2(t *testing.T) {
	value, _ := NewFunctions().Execute("isimage", StringValue("image/jpeg"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected isimage to be %v, received %v", true, actualValue)
	}
}

func TestIsImage3(t *testing.T) {
	value, _ := NewFunctions().Execute("isimage", StringValue("text/plain"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected isimage to be %v, received %v", false, actualValue)
	}
}

func TestIsImage4(t *testing.T) {
	value, _ := NewFunctions().Execute("isimage")

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected isimage to be %v, received %v", false, actualValue)
	}
}

func TestIsAudio1(t *testing.T) {
	value, _ := NewFunctions().Execute("isaudio", StringValue("audio/webm"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected isaudio to be %v, received %v", true, actualValue)
	}
}

func TestIsAudio2(t *testing.T) {
	value, _ := NewFunctions().Execute("isaudio", StringValue("audio/amr"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected isaudio to be %v, received %v", true, actualValue)
	}
}

func TestIsAudio3(t *testing.T) {
	value, _ := NewFunctions().Execute("isaudio", StringValue("text/plain"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected isaudio to be %v, received %v", false, actualValue)
	}
}

func TestIsAudio4(t *testing.T) {
	value, _ := NewFunctions().Execute("isaudio")

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected isaudio to be %v, received %v", false, actualValue)
	}
}

func TestIsVideo1(t *testing.T) {
	value, _ := NewFunctions().Execute("isvideo", StringValue("video/mpeg"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected isvideo to be %v, received %v", true, actualValue)
	}
}

func TestIsVideo2(t *testing.T) {
	value, _ := NewFunctions().Execute("isvideo", StringValue("video/3gpp"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected isvideo to be %v, received %v", true, actualValue)
	}
}

func TestIsVideo3(t *testing.T) {
	value, _ := NewFunctions().Execute("isvideo", StringValue("text/plain"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected isvideo to be %v, received %v", false, actualValue)
	}
}

func TestIsVideo4(t *testing.T) {
	value, _ := NewFunctions().Execute("isvideo")

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected isvideo to be %v, received %v", false, actualValue)
	}
}

func TestIsPdf1(t *testing.T) {
	value, _ := NewFunctions().Execute("ispdf", StringValue("application/pdf"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected ispdf to be %v, received %v", true, actualValue)
	}
}

func TestIsPdf2(t *testing.T) {
	value, _ := NewFunctions().Execute("ispdf", StringValue("application/x-pdf"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected ispdf to be %v, received %v", true, actualValue)
	}
}

func TestIsPdf3(t *testing.T) {
	value, _ := NewFunctions().Execute("ispdf", StringValue("text/plain"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected ispdf to be %v, received %v", false, actualValue)
	}
}

func TestIsPdf4(t *testing.T) {
	value, _ := NewFunctions().Execute("ispdf")

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected ispdf to be %v, received %v", false, actualValue)
	}
}

func TestFormatSizeWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("fmtsize")

	if err == nil {
		t.Fatalf("Expected an error while executing fmtsize with no parameter value")
	}
}

func TestFormatSizeWithIllegalValue(t *testing.T) {
	_, err := NewFunctions().Execute("fmtsize", BooleanValue(false))

	if err == nil {
		t.Fatalf("Expected an error while executing fmtsize with illegal parameter value")
	}
}

func TestFormatSize1(t *testing.T) {
	value, _ := NewFunctions().Execute("fmtsize", IntValue(4096))

	actualValue := value.GetAsString()
	if actualValue != "4.0 KiB" {
		t.Fatalf("Expected fmtsize to be %v, received %v", "4 Kib", actualValue)
	}
}

func TestFormatSize2(t *testing.T) {
	value, _ := NewFunctions().Execute("fmtsize", IntValue(98969))

	actualValue := value.GetAsString()
	if actualValue != "97 KiB" {
		t.Fatalf("Expected fmtsize to be %v, received %v", "4 Kib", actualValue)
	}
}

func TestIfBlank1(t *testing.T) {
	value, _ := NewFunctions().Execute("ifBlank", StringValue("   "), StringValue("NA"))
	expected := "NA"

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected ifBlank to be %v, received %v", expected, actualValue)
	}
}

func TestIfBlank2(t *testing.T) {
	value, _ := NewFunctions().Execute("ifBlank", StringValue(" content "), StringValue("NA"))
	expected := " content "

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected ifBlank to be %v, received %v", expected, actualValue)
	}
}

func TestIfBlank3(t *testing.T) {
	value, _ := NewFunctions().Execute("ifBlank", StringValue("  "), StringValue("   "))
	expected := "   "

	actualValue := value.GetAsString()
	if actualValue != expected {
		t.Fatalf("Expected ifBlank to be %v, received %v", expected, actualValue)
	}
}

func TestIfBlankWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("ifBlank")

	if err == nil {
		t.Fatalf("Expected an error while executing ifBlank with no parameter value")
	}
}

func TestStartsWith1(t *testing.T) {
	value, _ := NewFunctions().Execute("startsWith", StringValue("TestFile.log"), StringValue("Test"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected startsWith to be %v, received %v", true, actualValue)
	}
}

func TestStartsWith2(t *testing.T) {
	value, _ := NewFunctions().Execute("startsWith", StringValue("TestFile.log"), StringValue("File"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected startsWith to be %v, received %v", false, actualValue)
	}
}

func TestStartsWith3(t *testing.T) {
	value, _ := NewFunctions().Execute("startsWith", StringValue(" TestFile.log"), StringValue(" "))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected startsWith to be %v, received %v", true, actualValue)
	}
}

func TestStartsWith4(t *testing.T) {
	value, _ := NewFunctions().Execute("startsWith", StringValue(" TestFile.log"), StringValue("file"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected startsWith to be %v, received %v", false, actualValue)
	}
}

func TestStartsWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("startsWith")

	if err == nil {
		t.Fatalf("Expected an error while executing startsWith with no parameter value")
	}
}

func TestEndsWith1(t *testing.T) {
	value, _ := NewFunctions().Execute("endsWith", StringValue("TestFile.log"), StringValue("log"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected endsWith to be %v, received %v", true, actualValue)
	}
}

func TestEndsWith2(t *testing.T) {
	value, _ := NewFunctions().Execute("endsWith", StringValue("TestFile.log"), StringValue(".log"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected endsWith to be %v, received %v", true, actualValue)
	}
}

func TestEndsWith3(t *testing.T) {
	value, _ := NewFunctions().Execute("endsWith", StringValue("TestFile.log"), StringValue("File"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected endsWith to be %v, received %v", false, actualValue)
	}
}

func TestEndsWith4(t *testing.T) {
	value, _ := NewFunctions().Execute("endsWith", StringValue("TestFile.log"), StringValue("e.log"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected endsWith to be %v, received %v", true, actualValue)
	}
}

func TestEndsWith5(t *testing.T) {
	value, _ := NewFunctions().Execute("endsWith", StringValue("TestFile.log "), StringValue(" "))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected endsWith to be %v, received %v", true, actualValue)
	}
}
