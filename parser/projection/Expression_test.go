package projection

import (
	"reflect"
	"testing"
)

func TestExpressionsDisplayableAttributesWithAttributeName(t *testing.T) {
	expressions := Expressions{expressions: []*Expression{expressionWithAttribute("name")}}
	attributes := expressions.displayableAttributes()
	expected := []string{"name"}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestExpressionsDisplayableAttributesWithFunction(t *testing.T) {
	fun := &Function{
		name: "lower",
		left: &Expression{
			function: &Function{
				name: "upper",
				left: &Expression{attribute: "uid"},
			},
		},
	}
	expressions := Expressions{expressions: []*Expression{expressionWithFunction(fun)}}
	attributes := expressions.displayableAttributes()
	expected := []string{"lower(upper(uid))"}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestExpressionIsAFunction(t *testing.T) {
	expression := Expression{
		function: &Function{
			name: "upper",
			left: &Expression{attribute: "uid"},
		},
	}
	if expression.isAFunction() != true {
		t.Fatalf("Expected the expression to be a function")
	}
}

func TestExpressionIsNotFunction(t *testing.T) {
	expression := Expression{
		attribute: "uid",
	}
	if expression.isAFunction() != false {
		t.Fatalf("Expected the expression to not be a function")
	}
}