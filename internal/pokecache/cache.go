package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(timeout time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.readLoop(timeout)
	return c
}

func (c Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c Cache) readLoop(timeout time.Duration) {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > timeout {
				delete(c.cache, key)
			}
		}
		c.mux.Unlock()
	}
}
