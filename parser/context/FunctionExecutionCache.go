package context

type FunctionExecutionCache struct {
	entries map[Value]interface{}
}

func NewFunctionExecutionCache() *FunctionExecutionCache {
	return &FunctionExecutionCache{
		entries: make(map[Value]interface{}),
	}
}

func (functionExecutionCache *FunctionExecutionCache) Put(key Value, value interface{}) {
	functionExecutionCache.entries[key] = value
}

func (functionExecutionCache *FunctionExecutionCache) Get(key Value) (interface{}, bool) {
	value, ok := functionExecutionCache.entries[key]
	return value, ok
}
