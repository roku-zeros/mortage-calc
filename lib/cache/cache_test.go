package cache

import (
	"testing"
)

func TestCache_SetAndGet(t *testing.T) {
	cache := NewCache()

	if _, ok := cache.Get(0); ok {
		t.Errorf("Expected no value for id 0, but got a value")
	}

	value := "testValue"
	id := cache.Set(value)

	gotValue, ok := cache.Get(id)
	if !ok {
		t.Errorf("Expected to get value for id %d, but got none", id)
	}
	if gotValue != value {
		t.Errorf("Expected value %v, but got %v", value, gotValue)
	}
}

func TestCache_Delete(t *testing.T) {
	cache := NewCache()

	value := "testValue"
	id := cache.Set(value)

	cache.Delete(id)

	if _, ok := cache.Get(id); ok {
		t.Errorf("Expected no value for id %d after deletion, but got a value", id)
	}
}

func TestCache_GetCurrID(t *testing.T) {
	cache := NewCache()

	if currID := cache.GetCurrID(); currID != 0 {
		t.Errorf("Expected current ID to be 0, but got %d", currID)
	}

	cache.Set("value1")
	if currID := cache.GetCurrID(); currID != 1 {
		t.Errorf("Expected current ID to be 1, but got %d", currID)
	}

	cache.Set("value2")
	if currID := cache.GetCurrID(); currID != 2 {
		t.Errorf("Expected current ID to be 2, but got %d", currID)
	}
}
