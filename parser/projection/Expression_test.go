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
	fun := &FunctionInstance{
		name: "lower",
		args: []*Expression{
			{
				function: &FunctionInstance{
					name: "upper",
					args: []*Expression{expressionWithAttribute("uid")},
				},
			},
		},
	}
	expressions := Expressions{expressions: []*Expression{expressionWithFunctionInstance(fun)}}
	attributes := expressions.displayableAttributes()
	expected := []string{"lower(upper(uid))"}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestExpressionIsAFunction(t *testing.T) {
	expression := Expression{
		function: &FunctionInstance{
			name: "upper",
			args: []*Expression{
				{attribute: "uid"},
			},
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
