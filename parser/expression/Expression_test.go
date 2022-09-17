package expression

import (
	"goselect/parser/context"
	"reflect"
	"testing"
)

func TestExpressionsWithAttributes(t *testing.T) {
	expressions := ExpressionsWithAttributes([]string{"name", "size"})
	attributes := Expressions{Expressions: expressions}.DisplayableAttributes()
	expected := []string{"name", "size"}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestExpressionsWithFunction(t *testing.T) {
	expression := ExpressionWithFunctionInstance(FunctionInstanceWith(
		"lower",
		[]*Expression{ExpressionWithValue("content")},
		nil,
		false,
	))
	attributes := Expressions{Expressions: []*Expression{expression}}.DisplayableAttributes()
	expected := []string{"lower(content)"}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestExpressionsDisplayableAttributesWithAttributeName(t *testing.T) {
	expressions := Expressions{Expressions: []*Expression{ExpressionWithAttribute("name")}}
	attributes := expressions.DisplayableAttributes()
	expected := []string{"name"}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestExpressionsDisplayableAttributesWithValue(t *testing.T) {
	expressions := Expressions{Expressions: []*Expression{ExpressionWithValue("2")}}
	attributes := expressions.DisplayableAttributes()
	expected := []string{"2"}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestExpressionsDisplayableAttributesFunctionContainingAValue(t *testing.T) {
	fun := &FunctionInstance{
		name: "lower",
		args: []*Expression{ExpressionWithValue("2")},
	}
	expressions := Expressions{Expressions: []*Expression{ExpressionWithFunctionInstance(fun)}}
	attributes := expressions.DisplayableAttributes()
	expected := []string{"lower(2)"}
	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes to be %v, received %v", expected, attributes)
	}
}

func TestExpressionsDisplayableAttributesFunctionWithoutAnyParamaterValues(t *testing.T) {
	fun := &FunctionInstance{
		name: "now",
	}
	expressions := Expressions{Expressions: []*Expression{ExpressionWithFunctionInstance(fun)}}
	attributes := expressions.DisplayableAttributes()
	expected := []string{"now()"}
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

func TestExpressionEvaluate6(t *testing.T) {
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
	expressions := Expressions{Expressions: []*Expression{expression, ExpressionWithValue("content")}}
	functions := context.NewFunctions()

	values, _, _, _ := expressions.EvaluateWith(nil, functions)
	expected := "content##VALUE"

	actual := values[0].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}

	expected = "content"
	actual = values[1].GetAsString()
	if actual != expected {
		t.Fatalf("Expected function evaluation to return %v, received %v", expected, actual)
	}
}

func TestExpressionEvaluate7(t *testing.T) {
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "concat",
	})
	expressions := Expressions{Expressions: []*Expression{expression}}
	functions := context.NewFunctions()

	_, _, _, err := expressions.EvaluateWith(nil, functions)
	if err == nil {
		t.Fatalf("Expected an error while evaluating the expression but received none")
	}
}

func TestExpressionEvaluate8(t *testing.T) {
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "concat",
	})
	functions := context.NewFunctions()

	_, err, _ := expression.Evaluate(nil, functions)
	if err == nil {
		t.Fatalf("Expected an error while evaluating the expression but received none")
	}
}

func TestExpressionEvaluate9(t *testing.T) {
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "min",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name: "lower",
			}),
		},
	})
	functions := context.NewFunctions()

	_, err, _ := expression.Evaluate(nil, functions)
	if err == nil {
		t.Fatalf("Expected an error while evaluating the expression but received none")
	}
}

func TestExpressionEvaluate10(t *testing.T) {
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "min",
	})
	functions := context.NewFunctions()

	_, err, _ := expression.Evaluate(nil, functions)
	if err == nil {
		t.Fatalf("Expected an error while evaluating the expression but received none")
	}
}

func TestExpressionFullyEvaluateWithError(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name:  "min",
				state: functions.InitialState("min"),
			}),
		},
	})
	expressions := Expressions{Expressions: []*Expression{expression}}
	_, err := expressions.Expressions[0].FullyEvaluate(functions)

	if err == nil {
		t.Fatalf("Expected an error while fully evaluating min without any parameter value")
	}
}

func TestExpressionFullyEvaluateWithAnAggregateFunctionHavingAValue(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "upper",
		args: []*Expression{
			ExpressionWithFunctionInstance(&FunctionInstance{
				name:  "min",
				args:  []*Expression{ExpressionWithValue("someValue")},
				state: functions.InitialState("min"),
			}),
		},
	})
	expressions := Expressions{Expressions: []*Expression{expression}}
	value, _ := expressions.Expressions[0].FullyEvaluate(functions)

	if value.GetAsString() != "SOMEVALUE" {
		t.Fatalf("Expected final fully evaluate to return %v, received %v", "SOMEVALUE", value.GetAsString())
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

func TestExpressionHasAnAggregateFalse1(t *testing.T) {
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{},
	})
	hasAnAggregate := expression.HasAnAggregate()
	if hasAnAggregate != false {
		t.Fatalf("Expected hasAnAggregate to be false, received true")
	}
}

func TestExpressionHasAnAggregateFalse2(t *testing.T) {
	expression := ExpressionWithAttribute("name")
	hasAnAggregate := expression.HasAnAggregate()
	if hasAnAggregate != false {
		t.Fatalf("Expected hasAnAggregate to be false, received true")
	}
}

func TestExpressionHasAnAggregateTrue(t *testing.T) {
	functions := context.NewFunctions()
	expression := ExpressionWithFunctionInstance(&FunctionInstance{
		name: "lower",
		args: []*Expression{
			ExpressionWithValue("CONTENT"),
			ExpressionWithFunctionInstance(&FunctionInstance{
				name:        "min",
				args:        []*Expression{ExpressionWithValue("45")},
				state:       functions.InitialState("min"),
				isAggregate: true,
			}),
		},
	})
	hasAnAggregate := expression.HasAnAggregate()
	if hasAnAggregate != true {
		t.Fatalf("Expected hasAnAggregate to be true, received false")
	}
}

func TestAggregationCount1(t *testing.T) {
	functions := context.NewFunctions()
	expressions := Expressions{Expressions: []*Expression{
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
				ExpressionWithFunctionInstance(&FunctionInstance{
					name:        "min",
					args:        []*Expression{ExpressionWithValue("45")},
					state:       functions.InitialState("min"),
					isAggregate: true,
				}),
			},
		}),
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
			},
		}),
	}}

	aggregationCount := expressions.AggregationCount()
	if aggregationCount != 1 {
		t.Fatalf("Expected aggregation count to be %v, received %v", 1, aggregationCount)
	}
}

func TestAggregationCount2(t *testing.T) {
	expressions := Expressions{Expressions: []*Expression{
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
			},
		}),
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
			},
		}),
	}}

	aggregationCount := expressions.AggregationCount()
	if aggregationCount != 0 {
		t.Fatalf("Expected aggregation count to be %v, received %v", 0, aggregationCount)
	}
}

func TestAggregationCount3(t *testing.T) {
	functions := context.NewFunctions()
	expressions := Expressions{Expressions: []*Expression{
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
				ExpressionWithFunctionInstance(&FunctionInstance{
					name:        "min",
					args:        []*Expression{ExpressionWithValue("45")},
					state:       functions.InitialState("min"),
					isAggregate: true,
				}),
			},
		}),
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
				ExpressionWithFunctionInstance(&FunctionInstance{
					name:        "min",
					args:        []*Expression{ExpressionWithValue("45")},
					state:       functions.InitialState("min"),
					isAggregate: true,
				}),
			},
		}),
	}}

	aggregationCount := expressions.AggregationCount()
	if aggregationCount != 2 {
		t.Fatalf("Expected aggregation count to be %v, received %v", 2, aggregationCount)
	}
}

func TestCountOfExpression1(t *testing.T) {
	fun := &FunctionInstance{
		name: "now",
	}
	expressions := Expressions{Expressions: []*Expression{ExpressionWithFunctionInstance(fun)}}
	count := expressions.Count()

	if count != 1 {
		t.Fatalf("Expected expression count to be %v, received %v", 1, count)
	}
}

func TestCountOfExpression2(t *testing.T) {
	expressions := Expressions{Expressions: []*Expression{
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
			},
		}),
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
			},
		}),
	}}
	count := expressions.Count()

	if count != 2 {
		t.Fatalf("Expected expression count to be %v, received %v", 2, count)
	}
}

func TestExpressionAtAValidIndex(t *testing.T) {
	expressions := Expressions{Expressions: []*Expression{
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
			},
		}),
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "upper",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
			},
		}),
	}}
	expression := expressions.ExpressionAt(0)
	attribute := Expressions{Expressions: []*Expression{expression}}.DisplayableAttributes()[0]

	if attribute != "lower(CONTENT)" {
		t.Fatalf("Expected expression at index 0 to be %v, received %v", "lower(CONTENT)", attribute)
	}
}

func TestExpressionAtAnInvalidIndex(t *testing.T) {
	expressions := Expressions{Expressions: []*Expression{
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "lower",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
			},
		}),
		ExpressionWithFunctionInstance(&FunctionInstance{
			name: "upper",
			args: []*Expression{
				ExpressionWithValue("CONTENT"),
			},
		}),
	}}
	expression := expressions.ExpressionAt(2)
	if expression != nil {
		t.Fatalf("Expected expression to be nil, but was not")
	}
}

func simulate2RowExecution(expressions Expressions, functions *context.AllFunctions) ([]*Expression, []*Expression) {
	_, _, allExpressions1, _ := expressions.EvaluateWith(nil, functions)
	_, _, allExpressions2, _ := expressions.EvaluateWith(nil, functions)

	return allExpressions1, allExpressions2
}
