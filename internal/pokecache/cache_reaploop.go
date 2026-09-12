package pokecache

import (
	"time"
)

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.duration)
	defer ticker.Stop()

	for range ticker.C {
		actualTime := time.Now()
		c.mu.Lock()
		for id, entry := range c.mapCache {
			duration := actualTime.Sub(entry.createdAt)
			if c.duration < duration {
				delete(c.mapCache, id)
			}
		}
		c.mu.Unlock()
	}
}
