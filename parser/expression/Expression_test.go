package expression

import (
	"goselect/parser/context"
	"reflect"
	"testing"
)

func TestExpressionsDisplayableAttributesWithAttributeName(t *testing.T) {
	expressions := Expressions{Expressions: []*Expression{ExpressionWithAttribute("name")}}
	attributes := expressions.DisplayableAttributes()
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
					args: []*Expression{ExpressionWithAttribute("uid")},
				},
			},
		},
	}
	expressions := Expressions{Expressions: []*Expression{ExpressionWithFunctionInstance(fun)}}
	attributes := expressions.DisplayableAttributes()
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
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{
			ExpressionWithValue("CONTENT"),
		},
	})
	expressions := Expressions{Expressions: []*Expression{expression}}
	functions := context.NewFunctions()

	values, _, _, _ := expressions.EvaluateWith(nil, functions)

	expected := "content"
	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
}

func TestExpressionEvaluate2(t *testing.T) {
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "concat",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name: "lower",
				args: []*Expression{
					ExpressionWithValue("CONTENT"),
				},
			}),
			ExpressionWithValue("##"),
			ExpressionWithFunctionInstance(&FunctionInstance{
				name: "upper",
				args: []*Expression{
					ExpressionWithValue("value"),
				},
			}),
		},
	})
	expressions := Expressions{Expressions: []*Expression{expression}}
	functions := context.NewFunctions()

	values, _, _, _ := expressions.EvaluateWith(nil, functions)

	expected := "content##VALUE"
	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
}

func TestExpressionEvaluate3(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name:  "Count",
		args:  []*Expression{},
		state: functions.InitialState("Count"),
	})
	expressions := Expressions{Expressions: []*Expression{expression}}

	values, fullyEvaluated, _, _ := expressions.EvaluateWith(nil, functions)

	expected := "1"
	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
	if fullyEvaluated[0] != false {
		t.Fatalf("Expected Count to be partially evaluated but was not")
	}
}

func TestExpressionEvaluate4(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "Count",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name: "lower",
				args: []*Expression{
					ExpressionWithValue("CONTENT"),
				},
			}),
		},
		state: functions.InitialState("Count"),
	})
	expressions := Expressions{Expressions: []*Expression{expression}}

	values, fullyEvaluated, _, _ := expressions.EvaluateWith(nil, functions)

	expected := "1"
	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
	if fullyEvaluated[0] != false {
		t.Fatalf("Expected Count to be partially evaluated but was not")
	}
}

func TestExpressionEvaluate5(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name:  "Count",
				args:  []*Expression{},
				state: functions.InitialState("Count"),
			}),
		},
	})
	expressions := Expressions{Expressions: []*Expression{expression}}

	_, fullyEvaluated, _, _ := expressions.EvaluateWith(nil, functions)

	if fullyEvaluated[0] != false {
		t.Fatalf("Expected Count to be partially evaluated but was not")
	}
}

func TestExpressionFullyEvaluate(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name:  "Count",
				args:  []*Expression{},
				state: functions.InitialState("Count"),
			}),
		},
		state: functions.InitialState("Count"),
	})
	expressions := Expressions{Expressions: []*Expression{expression}}

	_, _, allExpressions, _ := expressions.EvaluateWith(nil, functions)
	value, _ := allExpressions[0].FullyEvaluate(functions)

	expected := "1"
	actual := value.GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
}

func TestExpressionEvaluateWithNestingOfCount(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name: "Count",
				args: []*Expression{
					ExpressionWithFunctionInstance(&FunctionInstance{
						name:  "Count",
						args:  []*Expression{},
						state: functions.InitialState("Count"),
					}),
				},
				state: functions.InitialState("Count"),
			}),
		},
	})
	expressions := Expressions{Expressions: []*Expression{expression}}
	allExpressions1, allExpressions2 := simulate2RowExecution(expressions, functions)

	value1, _ := allExpressions1[0].FullyEvaluate(functions)
	if value1.GetAsString() != "1" {
		t.Fatalf("Expected lower(Count(Count)) to be %v, received %v", "1", value1.GetAsString())
	}

	value2, _ := allExpressions2[0].FullyEvaluate(functions)
	if value2.GetAsString() != "1" {
		t.Fatalf("Expected lower(Count(Count)) to be %v, received %v", "1", value2.GetAsString())
	}
}

func TestExpressionEvaluateWithLowerFunctionInsideCount(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "Count",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name: "lower",
				args: []*Expression{
					ExpressionWithValue("NAME"),
				},
			}),
		},
		state: functions.InitialState("Count"),
	})
	expressions := Expressions{Expressions: []*Expression{expression}}
	allExpressions1, allExpressions2 := simulate2RowExecution(expressions, functions)

	value1, _ := allExpressions1[0].FullyEvaluate(functions)
	if value1.GetAsString() != "2" {
		t.Fatalf("Expected Count(lower) to be %v, received %v", "2", value1.GetAsString())
	}

	value2, _ := allExpressions2[0].FullyEvaluate(functions)
	if value2.GetAsString() != "2" {
		t.Fatalf("Expected Count(lower) to be %v, received %v", "2", value2.GetAsString())
	}
}

func TestExpressionEvaluateWithNestingOfAverage(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "avg",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name: "avg",
				args: []*Expression{
					ExpressionWithFunctionInstance(&FunctionInstance{
						name: "len",
						args: []*Expression{
							ExpressionWithValue("CONTENT"),
						},
					}),
				},
				state: functions.InitialState("avg"),
			}),
		},
		state: functions.InitialState("avg"),
	})
	expressions := Expressions{Expressions: []*Expression{expression}}
	allExpressions1, allExpressions2 := simulate2RowExecution(expressions, functions)

	value1, _ := allExpressions1[0].FullyEvaluate(functions)
	if value1.GetAsString() != "7.00" {
		t.Fatalf("Expected avg(avg(len('CONTENT'))) to be %v, received %v", "7.00", value1.GetAsString())
	}

	value2, _ := allExpressions2[0].FullyEvaluate(functions)
	if value2.GetAsString() != "7.00" {
		t.Fatalf("Expected avg(avg(len('CONTENT'))) to be %v, received %v", "7.00", value2.GetAsString())
	}
}

func simulate2RowExecution(expressions Expressions, functions *context.AllFunctions) ([]*Expression, []*Expression) {
	_, _, allExpressions1, _ := expressions.EvaluateWith(nil, functions)
	_, _, allExpressions2, _ := expressions.EvaluateWith(nil, functions)

	return allExpressions1, allExpressions2
}
