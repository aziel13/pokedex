package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	value    []byte
}

type Cache struct {
	mu            sync.RWMutex
	reap_interval time.Duration
	items         map[string]cacheEntry
}

func NewCacheEntry(interval time.Duration) *Cache {
	newCache := &Cache{

		items:         make(map[string]cacheEntry),
		reap_interval: interval,
	}

	return newCache
}

func (c *Cache) Add(key string, value []byte) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = cacheEntry{

		createAt: time.Now(),
		value:    value,
	}

}

func (c *Cache) Get(key string) []byte {

	c.reapLoop(c.reap_interval)

	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.items[key]

	if !ok {
		return nil
	}

	return entry.value
}

func (c *Cache) reapLoop(interval time.Duration) {

	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.items {
		if time.Since(entry.createAt) > interval {
			delete(c.items, key)
		}
	}

}
