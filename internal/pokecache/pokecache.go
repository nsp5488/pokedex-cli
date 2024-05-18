package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		data: make(map[string]cacheEntry),
		mu:   sync.Mutex{},
	}
	go c.reapLoop(interval)
	return &c
}

type Cache struct {
	data map[string]cacheEntry
	mu   sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) (data []byte, success bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mu.Lock()
		now := time.Now().Add(-interval)
		for key, entry := range c.data {
			if entry.createdAt.Before(now) {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
