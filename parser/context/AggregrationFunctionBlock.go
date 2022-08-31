package context

type CountFunctionBlock struct{}

func (c *CountFunctionBlock) run(initialState *AggregateFunctionState, _ ...Value) (*AggregateFunctionState, error) {
	return &AggregateFunctionState{Initial: Uint32Value(initialState.Initial.uint32Value + 1)}, nil
}

func (c *CountFunctionBlock) initialState() *AggregateFunctionState {
	return &AggregateFunctionState{Initial: Uint32Value(0)}
}
