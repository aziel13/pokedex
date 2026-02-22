package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {

	//cached results from api call
	interval := 1 * time.Second

	actual := NewCacheEntry(interval)

	if actual.items == nil {

		t.Errorf("Expected a initialized empty map but items is nil. ")

	}

	if len(actual.items) != 0 {

		t.Errorf("Expected the newly initialized map to contain no items, got %d", len(actual.items))

	}

}

func setupCacheForTesting() *Cache {

	interval := 10 * time.Second

	cache := NewCacheEntry(interval)

	return cache
}

func TestSetCacheValue(t *testing.T) {

	cache := setupCacheForTesting()

	inputString := "This is a test"

	value := []byte(inputString)
	key := "https://example.com"

	cache.Add(key, value)

	actual, ok := cache.items[key]

	if !ok {
		t.Errorf("Expected a cache item for key %s to exist", key)
	}

	if !bytes.Equal(actual.value, value) {
		t.Errorf("Expected value %s for key %s", string(value), key)
	}

}

func TestGetCache(t *testing.T) {

	createTime := time.Now()
	cache := setupCacheForTesting()

	inputString := "This is a test"

	value := []byte(inputString)
	key := "https://example.com"

	cache.items[key] = cacheEntry{
		createAt: createTime,
		value:    value,
	}

	actual := cache.Get(key)
	if actual == nil {
		t.Errorf("Expected a cache item for key %s to exist", key)
	}

	if !bytes.Equal(actual, value) {
		t.Errorf("Expected value %s for key %s Instead got %s", string(value), key, string(actual))
	}

}

func TestReapLoop(t *testing.T) {
	cache := setupCacheForTesting()

	inputString := "This is a test"

	value := []byte(inputString)
	key := "https://example.com"

	cache.Add(key, value)
	waitTime := 5 * time.Second
	reapInterval := 1 * time.Second

	time.Sleep(waitTime)
	cache.reapLoop(reapInterval)

	ok := cache.Get(key)

	if ok != nil {

		t.Errorf("Expected a cache item for key %s to not exist but it still exists", key)

	}

}
