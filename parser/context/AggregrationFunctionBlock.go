package context

import (
	"fmt"
	"goselect/parser/error/messages"
)

type CountFunctionBlock struct{}
type AverageFunctionBlock struct{}

func (c *CountFunctionBlock) initialState() *FunctionState {
	return &FunctionState{Initial: Uint32Value(0), isUpdated: false}
}

func (c *CountFunctionBlock) run(initialState *FunctionState, _ ...Value) (*FunctionState, error) {
	return &FunctionState{Initial: Uint32Value(initialState.Initial.uint32Value + 1), isUpdated: true}, nil
}

func (c *CountFunctionBlock) finalValue(currentState *FunctionState, _ []Value) (Value, error) {
	if currentState.isUpdated {
		return currentState.Initial, nil
	}
	return Uint32Value(1), nil
}

func (a *AverageFunctionBlock) initialState() *FunctionState {
	return &FunctionState{Initial: Float64Value(0.0), extras: make(map[interface{}]Value), isUpdated: false}
}

func (a *AverageFunctionBlock) run(initialState *FunctionState, args ...Value) (*FunctionState, error) {
	if err := ensureNParametersOrError(args, FunctionNameAverage, 1); err != nil {
		return nil, err
	}
	if theOnlyArgument, err := args[0].GetNumericAsFloat64(); err != nil {
		return nil, fmt.Errorf(messages.ErorMessageFunctionNamePrefixWithExistingError, FunctionNameAverage, err)
	} else {
		existingCount := initialState.extras["count"]
		return &FunctionState{
			Initial:   Float64Value(initialState.Initial.float64Value + theOnlyArgument),
			extras:    map[interface{}]Value{"count": Uint32Value(existingCount.uint32Value + 1)},
			isUpdated: true,
		}, nil
	}
}

func (a *AverageFunctionBlock) finalValue(currentState *FunctionState, values []Value) (Value, error) {
	if currentState.isUpdated {
		asFloat64, _ := currentState.Initial.GetNumericAsFloat64()
		return Float64Value(asFloat64 / ((float64)(currentState.extras["count"].uint32Value))), nil
	}
	if err := ensureNParametersOrError(values, FunctionNameAverage, 1); err != nil {
		return EmptyValue(), err
	}
	if v, err := values[0].GetNumericAsFloat64(); err != nil {
		return EmptyValue(), fmt.Errorf(messages.ErorMessageFunctionNamePrefixWithExistingError, FunctionNameAverage, err)
	} else {
		return Float64Value(v), nil
	}
}
