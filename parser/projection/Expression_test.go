package projection

import (
	"goselect/parser/context"
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

func TestExpressionEvaluate1(t *testing.T) {
	expression := expressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{
			expressionWithValue("CONTENT"),
		},
	})
	expressions := Expressions{expressions: []*Expression{expression}}
	functions := context.NewFunctions()

	values, _, _, _ := expressions.evaluateWith(nil, functions)

	expected := "content"
	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
}

func TestExpressionEvaluate2(t *testing.T) {
	expression := expressionWithFunctionInstance(&FunctionInstance{
		name: "concat",
		args: []*Expression{
			expressionWithFunctionInstance(&FunctionInstance{
				name: "lower",
				args: []*Expression{
					expressionWithValue("CONTENT"),
				},
			}),
			expressionWithValue("##"),
			expressionWithFunctionInstance(&FunctionInstance{
				name: "upper",
				args: []*Expression{
					expressionWithValue("value"),
				},
			}),
		},
	})
	expressions := Expressions{expressions: []*Expression{expression}}
	functions := context.NewFunctions()

	values, _, _, _ := expressions.evaluateWith(nil, functions)

	expected := "content##VALUE"
	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
}

func TestExpressionEvaluate3(t *testing.T) {
	functions := context.NewFunctions()
	expression := expressionWithFunctionInstance(&FunctionInstance{
		name:  "count",
		args:  []*Expression{},
		state: functions.InitialState("count"),
	})
	expressions := Expressions{expressions: []*Expression{expression}}

	values, fullyEvaluated, _, _ := expressions.evaluateWith(nil, functions)

	expected := "1"
	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
	if fullyEvaluated[0] != false {
		t.Fatalf("Expected count to be partially evaluated but was not")
	}
}

func TestExpressionEvaluate4(t *testing.T) {
	functions := context.NewFunctions()
	expression := expressionWithFunctionInstance(&FunctionInstance{
		name: "count",
		args: []*Expression{
			expressionWithFunctionInstance(&FunctionInstance{
				name: "lower",
				args: []*Expression{
					expressionWithValue("CONTENT"),
				},
			}),
		},
		state: functions.InitialState("count"),
	})
	expressions := Expressions{expressions: []*Expression{expression}}

	values, fullyEvaluated, _, _ := expressions.evaluateWith(nil, functions)

	expected := "1"
	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
	if fullyEvaluated[0] != false {
		t.Fatalf("Expected count to be partially evaluated but was not")
	}
}

func TestExpressionEvaluate5(t *testing.T) {
	functions := context.NewFunctions()
	expression := expressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{
			expressionWithFunctionInstance(&FunctionInstance{
				name:  "count",
				args:  []*Expression{},
				state: functions.InitialState("count"),
			}),
		},
		state: functions.InitialState("count"),
	})
	expressions := Expressions{expressions: []*Expression{expression}}

	values, fullyEvaluated, _, _ := expressions.evaluateWith(nil, functions)

	expected := "1"
	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
	if fullyEvaluated[0] != false {
		t.Fatalf("Expected count to be partially evaluated but was not")
	}
}

func TestExpressionFullyEvaluate(t *testing.T) {
	functions := context.NewFunctions()
	expression := expressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{
			expressionWithFunctionInstance(&FunctionInstance{
				name:  "count",
				args:  []*Expression{},
				state: functions.InitialState("count"),
			}),
		},
		state: functions.InitialState("count"),
	})
	expressions := Expressions{expressions: []*Expression{expression}}

	_, _, allExpressions, _ := expressions.evaluateWith(nil, functions)
	value := allExpressions[0].FullyEvaluate(functions)

	expected := "1"
	actual := value.GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
}
