package context

import (
	"testing"
)

func TestLikeWithMissingParameterValue(t *testing.T) {
	_, err := NewFunctions().Execute("like")

	if err == nil {
		t.Fatalf("Expected an error while executing like with no parameter value")
	}
}

func TestLikeWithInvalidRegex(t *testing.T) {
	_, err := NewFunctions().Execute("like", StringValue("name"), StringValue("*"))

	if err == nil {
		t.Fatalf("Expected an error while executing like with invalid regex")
	}
}

func TestLike1(t *testing.T) {
	value, _ := NewFunctions().Execute("like", StringValue("sample.log"), StringValue(".*log"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected like to be %v, received %v", true, actualValue)
	}
}

func TestLike2(t *testing.T) {
	value, _ := NewFunctions().Execute("like", StringValue("sample.log"), StringValue("sam.*"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected like to be %v, received %v", true, actualValue)
	}
}

func TestLike3(t *testing.T) {
	value, _ := NewFunctions().Execute("like", StringValue("test123_Definitive"), StringValue("[a-z]{1}[0-9]"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected like to be %v, received %v", true, actualValue)
	}
}

func TestLike4(t *testing.T) {
	value, _ := NewFunctions().Execute("like", StringValue("test123_Definitive"), StringValue("[a-z]{1}[0-9]_"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected like to be %v, received %v", false, actualValue)
	}
}

func TestLike5(t *testing.T) {
	value, _ := NewFunctions().Execute("like", StringValue("test123_Definitive"), StringValue("[a-z]{4}[0-9]{3}_"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected like to be %v, received %v", true, actualValue)
	}
}

func TestLike6(t *testing.T) {
	value, _ := NewFunctions().Execute("like", StringValue("123_image_1"), StringValue("^[0-9]"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected like to be %v, received %v", true, actualValue)
	}
}

func TestLike7(t *testing.T) {
	value, _ := NewFunctions().Execute("like", StringValue("image_123"), StringValue("[0-9]$"))

	actualValue, _ := value.GetBoolean()
	if actualValue != true {
		t.Fatalf("Expected like to be %v, received %v", true, actualValue)
	}
}

func TestLike8(t *testing.T) {
	value, _ := NewFunctions().Execute("like", StringValue("anything"), StringValue("[0-9]$"))

	actualValue, _ := value.GetBoolean()
	if actualValue != false {
		t.Fatalf("Expected like to be %v, received %v", false, actualValue)
	}
}
