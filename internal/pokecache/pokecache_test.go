package pokecache

import (
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	// Test adding and retrieving an item
	cache.Add("test_key", []byte("test_value"))
	val, ok := cache.Get("test_key")
	if !ok {
		t.Errorf("Expected to find key 'test_key'")
	}
	if string(val) != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", string(val))
	}

	// Test retrieving a non-existent item
	_, ok = cache.Get("non_existent_key")
	if ok {
		t.Errorf("Did not expect to find key 'non_existent_key'")
	}

	// Add more tests as needed
}
