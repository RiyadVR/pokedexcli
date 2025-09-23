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
	mu       *sync.Mutex
	data     map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		mu:       &sync.Mutex{},
		data:     map[string]cacheEntry{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.data {
			if now.Sub(v.createdAt) > c.interval {
				delete(c.data, k)
			}
		}
		c.mu.Unlock()
	}
}
