package manifold

import (
	"testing"
)

func TestShardedData(t *testing.T) {
	data := NewShardedData()

	// Test setting and getting a value.
	data.Set("key1", "value1")
	val, ok := data.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("Expected %s, got %s", "value1", val)
	}

	// Test updating a value.
	data.Set("key1", "newvalue1")
	val, ok = data.Get("key1")
	if !ok || val != "newvalue1" {
		t.Errorf("Expected %s, got %s", "newvalue1", val)
	}

	// Test that a non-existent key returns an empty string.
	val, ok = data.Get("nonexistent")
	if ok || val != "" {
		t.Errorf("Expected an empty string, got %s", val)
	}

	// Test concurrency safety.
	const numGoroutines = 100
	done := make(chan bool)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			key := "key" + string(rune(i))
			data.Set(key, "value")
			done <- true
		}(i)
	}
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
	for i := 0; i < numGoroutines; i++ {
		key := "key" + string(rune(i))
		val, ok = data.Get(key)
		if !ok || val != "value" {
			t.Errorf("Expected %s, got %s for key %s", "value", val, key)
		}
	}
}
