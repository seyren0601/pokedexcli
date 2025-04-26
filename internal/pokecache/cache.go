package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	caches map[string]*cacheEntry
	mu     *sync.RWMutex
	ticker time.Ticker
}

func (c Cache) Add(key string, value []byte) {
	newCache := &cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

	c.mu.Lock()
	c.caches[key] = newCache
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if entry, ok := c.caches[key]; ok {
		// Refresh entry creation time
		entry.createdAt = time.Now()
		return entry.val, true
	}

	return nil, false
}

func (c Cache) reapLoop(duration time.Duration) {
	for range c.ticker.C {
		for key, entry := range c.caches {
			timePassed := time.Since(entry.createdAt)
			if timePassed > duration {
				c.mu.Lock()
				delete(c.caches, key)
				c.mu.Unlock()
			}
		}
	}
}

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		caches: map[string]*cacheEntry{},
		mu:     &sync.RWMutex{},
		ticker: *time.NewTicker(duration),
	}

	go cache.reapLoop(duration)

	return cache
}
