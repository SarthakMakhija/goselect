package context

type CountFunctionBlock struct {
	mutableCount uint32
}

type CountDistinctFunctionBlock struct {
	values map[Value]bool
}

func (c *CountFunctionBlock) run(_ ...Value) (Value, error) {
	c.mutableCount = c.mutableCount + 1
	return Uint32Value(c.mutableCount), nil
}

func (c *CountFunctionBlock) finalState() Value {
	return Uint32Value(c.mutableCount)
}

func (c *CountDistinctFunctionBlock) run(args ...Value) (Value, error) {
	if err := ensureNParametersOrError(args, FunctionNameCountDistinct, 1); err != nil {
		return EmptyValue(), err
	}
	c.values[args[0]] = true
	return Uint32Value(uint32(len(c.values))), nil
}

func (c *CountDistinctFunctionBlock) finalState() Value {
	return Uint32Value(uint32(len(c.values)))
}
