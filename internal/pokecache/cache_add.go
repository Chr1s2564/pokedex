package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mapCache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
