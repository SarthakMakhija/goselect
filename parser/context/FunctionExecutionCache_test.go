//go:build unit
// +build unit

package context

import "testing"

func TestGetFromTheCacheValueForAnExistingKey(t *testing.T) {
	cache := NewFunctionExecutionCache()
	cache.Put(StringValue("name"), "goselect")

	value, _ := cache.Get(StringValue("name"))
	if value.(string) != "goselect" {
		t.Fatalf("Expected value to be %v, received %v", "goselect", value)
	}
}

func TestGetFromTheCacheValueForANonExistingKey(t *testing.T) {
	cache := NewFunctionExecutionCache()
	cache.Put(StringValue("name"), "goselect")

	value, ok := cache.Get(StringValue("size"))
	if ok == true {
		t.Fatalf("Expected key size to not exist in the cache but it did with value %v", value)
	}
}
