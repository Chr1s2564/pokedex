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
	mapCache map[string]cacheEntry
	mu       sync.Mutex
	duration time.Duration
}

func GetMapCache() map[string]cacheEntry {
	return make(map[string]cacheEntry)
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mapCache: GetMapCache(),
		mu:       sync.Mutex{},
		duration: interval,
	}
	go cache.reapLoop()
	return cache
}
