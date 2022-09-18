//go:build unit
// +build unit

package context

import (
	"testing"
)

func TestParseDateTime1(t *testing.T) {
	time, _ := parse("2009-08-28", "dt")
	dateAsStr := formatDate(time).GetAsString()
	expected := "2009-August-28"

	if expected != dateAsStr {
		t.Fatalf("Expected parsing of date/time to return %v, received %v", expected, dateAsStr)
	}
}

func TestParseDateTime2(t *testing.T) {
	time, _ := parse("2009-08-28T10:14:28", "ts")
	expected := "2009-08-28 10:14:28 +0000 UTC"

	if expected != time.String() {
		t.Fatalf("Expected parsing of date/time to return %v, received %v", expected, time.String())
	}
}

func TestParseDateTime3(t *testing.T) {
	time, _ := parse("2009-08-28T10:14:29.009Z", "tsfull")
	expected := "2009-08-28 10:14:29.009 +0000 UTC"

	if expected != time.String() {
		t.Fatalf("Expected parsing of date/time to return %v, received %v", expected, time.String())
	}
}

func TestParseDateTime4(t *testing.T) {
	_, err := parse("2009-08-28T10:14:29.009Z", "unknown")

	if err == nil {
		t.Fatalf("Expected an error while parsing a date/time with an unknown Id")
	}
}

func TestParseDateTime5(t *testing.T) {
	_, err := parse("2009-August-28", "dt")

	if err == nil {
		t.Fatalf("Expected an error while parsing a date/time in an unsupported Format")
	}
}

func TestDateTimeSupportedFormats(t *testing.T) {
	supportedFormats := SupportedFormats()
	if len(supportedFormats) == 0 {
		t.Fatalf("Expected length of supported formats to be greater than zero but was zer")
	}
}
