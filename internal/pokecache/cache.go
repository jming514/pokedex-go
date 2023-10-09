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
	c   map[string]cacheEntry
}

// NewCache create cache and reap every 5 minutes
func NewCache(interval time.Duration) *Cache {
	_ = time.NewTicker(interval * time.Second)

	cache := Cache{
		mux: &sync.RWMutex{},
		c:   make(map[string]cacheEntry),
	}

	go cache.reapLoop()

	return &cache
}

// Add add an entry to cache
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.c[key] = newEntry
}

// Get get an entry from cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	for k, v := range c.c {
		if k == key {
			return v.val, true
		}
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	interval := 5 * time.Minute

	for {
		c.mux.Lock()
		for k, v := range c.c {
			lifetime := time.Now().Sub(v.createdAt)
			if lifetime >= interval {
				delete(c.c, k)
			}
		}
		defer c.mux.Unlock()

		time.Sleep(interval)
	}
}
