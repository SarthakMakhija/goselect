package context

import (
	"fmt"
	"goselect/parser/error/messages"
)

type CountFunctionBlock struct{}
type CountDistinctFunctionBlock struct{}
type SumFunctionBlock struct{}
type AverageFunctionBlock struct{}
type MinFunctionBlock struct{}

func (c *CountFunctionBlock) initialState() *FunctionState {
	return &FunctionState{Initial: zeroUint32Value, isUpdated: false}
}

func (c *CountFunctionBlock) run(initialState *FunctionState, _ ...Value) (*FunctionState, error) {
	return &FunctionState{Initial: Uint32Value(initialState.Initial.uint32Value + 1), isUpdated: true}, nil
}

func (c *CountFunctionBlock) finalValue(currentState *FunctionState, _ []Value) (Value, error) {
	if currentState.isUpdated {
		return currentState.Initial, nil
	}
	return oneUint32Value, nil
}

func (c *CountDistinctFunctionBlock) initialState() *FunctionState {
	return &FunctionState{
		extras:    make(map[interface{}]Value),
		isUpdated: false,
	}
}

func (c *CountDistinctFunctionBlock) run(initialState *FunctionState, args ...Value) (*FunctionState, error) {
	if err := ensureNParametersOrError(args, FunctionNameCountDistinct, 1); err != nil {
		return nil, err
	}
	theOnlyArgument, existenceByValue := args[0].GetAsString(), initialState.extras
	existenceByValue[theOnlyArgument] = trueBooleanValue

	return &FunctionState{
		extras:    existenceByValue,
		isUpdated: true,
	}, nil
}

func (c *CountDistinctFunctionBlock) finalValue(currentState *FunctionState, _ []Value) (Value, error) {
	if currentState.isUpdated {
		return Uint32Value(uint32(len(currentState.extras))), nil
	}
	return oneUint32Value, nil
}

func (c *SumFunctionBlock) initialState() *FunctionState {
	return &FunctionState{
		Initial:   Float64Value(0),
		isUpdated: false,
	}
}

func (c *SumFunctionBlock) run(initialState *FunctionState, args ...Value) (*FunctionState, error) {
	if err := ensureNParametersOrError(args, FunctionNameSum, 1); err != nil {
		return nil, err
	}
	if theOnlyArgument, err := args[0].GetNumericAsFloat64(); err != nil {
		return nil, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameSum, err)
	} else {
		return &FunctionState{
			Initial:   Float64Value(initialState.Initial.float64Value + theOnlyArgument),
			isUpdated: true,
		}, nil
	}
}

func (c *SumFunctionBlock) finalValue(currentState *FunctionState, values []Value) (Value, error) {
	if currentState.isUpdated {
		return currentState.Initial, nil
	}
	if err := ensureNParametersOrError(values, FunctionNameSum, 1); err != nil {
		return EmptyValue, err
	}
	if v, err := values[0].GetNumericAsFloat64(); err != nil {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameSum, err)
	} else {
		return Float64Value(v), nil
	}
}

func (a *AverageFunctionBlock) initialState() *FunctionState {
	return &FunctionState{
		Initial:   Float64Value(0.0),
		extras:    make(map[interface{}]Value),
		isUpdated: false,
	}
}

func (a *AverageFunctionBlock) run(initialState *FunctionState, args ...Value) (*FunctionState, error) {
	if err := ensureNParametersOrError(args, FunctionNameAverage, 1); err != nil {
		return nil, err
	}
	if theOnlyArgument, err := args[0].GetNumericAsFloat64(); err != nil {
		return nil, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameAverage, err)
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
		return EmptyValue, err
	}
	if v, err := values[0].GetNumericAsFloat64(); err != nil {
		return EmptyValue, fmt.Errorf(messages.ErrorMessageFunctionNamePrefixWithExistingError, FunctionNameAverage, err)
	} else {
		return Float64Value(v), nil
	}
}

func (m *MinFunctionBlock) initialState() *FunctionState {
	return &FunctionState{
		Initial:   EmptyValue,
		isUpdated: false,
	}
}

func (m *MinFunctionBlock) run(initialState *FunctionState, args ...Value) (*FunctionState, error) {
	if err := ensureNParametersOrError(args, FunctionNameMin, 1); err != nil {
		return nil, err
	}
	if !initialState.isUpdated {
		return &FunctionState{
			Initial:   args[0],
			isUpdated: true,
		}, nil
	}
	comparisonResult := initialState.Initial.CompareTo(args[0])
	if comparisonResult == CompareToEqual || comparisonResult == CompareToLessThan { //initial == incoming || initial < incoming
		return initialState, nil
	} else {
		return &FunctionState{
			Initial:   args[0],
			isUpdated: true,
		}, nil
	}
}

func (m *MinFunctionBlock) finalValue(currentState *FunctionState, values []Value) (Value, error) {
	if currentState.isUpdated {
		return currentState.Initial, nil
	}
	if err := ensureNParametersOrError(values, FunctionNameMin, 1); err != nil {
		return EmptyValue, err
	}
	return values[0], nil
}
