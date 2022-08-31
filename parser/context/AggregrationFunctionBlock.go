package context

type CountFunctionBlock struct{}
type AverageFunctionBlock struct{}

func (c *CountFunctionBlock) initialState() *FunctionState {
	return &FunctionState{Initial: Uint32Value(0)}
}

func (c *CountFunctionBlock) run(initialState *FunctionState, _ ...Value) (*FunctionState, error) {
	return &FunctionState{Initial: Uint32Value(initialState.Initial.uint32Value + 1)}, nil
}

func (c *CountFunctionBlock) finalValue(currentState *FunctionState) Value {
	return currentState.Initial
}

func (a *AverageFunctionBlock) initialState() *FunctionState {
	return &FunctionState{Initial: Float64Value(0.0), extras: make(map[interface{}]Value)}
}

func (a *AverageFunctionBlock) run(initialState *FunctionState, args ...Value) (*FunctionState, error) {
	if err := ensureNParametersOrError(args, FunctionNameAverage, 1); err != nil {
		return nil, err
	}
	if theOnlyArgument, err := args[0].GetNumericAsFloat64(); err != nil {
		return nil, err
	} else {
		existingCount := initialState.extras["count"]
		return &FunctionState{
			Initial: Float64Value(initialState.Initial.float64Value + theOnlyArgument),
			extras:  map[interface{}]Value{"count": Uint32Value(existingCount.uint32Value + 1)},
		}, nil
	}
}

func (a *AverageFunctionBlock) finalValue(currentState *FunctionState) Value {
	return Float64Value(currentState.Initial.float64Value / ((float64)(currentState.extras["count"].uint32Value)))
}
