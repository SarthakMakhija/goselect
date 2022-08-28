package projection

import (
	"reflect"
	"testing"
)

func TestExpressionsDisplayableColumnsWithColumnName(t *testing.T) {
	expressions := Expressions{expressions: []*Expression{expressionWithColumn("name")}}
	columns := expressions.displayableAttributes()
	expected := []string{"name"}

	if !reflect.DeepEqual(expected, columns) {
		t.Fatalf("Expected columns to be %v, received %v", expected, columns)
	}
}

func TestExpressionsDisplayableColumnsWithFunction(t *testing.T) {
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
	columns := expressions.displayableAttributes()
	expected := []string{"lower(upper(uid))"}

	if !reflect.DeepEqual(expected, columns) {
		t.Fatalf("Expected columns to be %v, received %v", expected, columns)
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
