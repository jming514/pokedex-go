package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mux *sync.RWMutex
	C   map[string]cacheEntry
}

// NewCache create cache and reap every 5 minutes
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mux: &sync.RWMutex{},
		C:   make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)

	return cache
}

// Add add an entry to cache
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.C[key] = newEntry
}

// Get get an entry from cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	entry, ok := c.C[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(now time.Time, ttl time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for k, v := range c.C {
		lifetime := now.Sub(v.createdAt)
		if lifetime > ttl {
			delete(c.C, k)
		}
	}
}
