package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.mapCache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
