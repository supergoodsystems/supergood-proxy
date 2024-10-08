package cache

import "sync"

func New() Cache {
	return Cache{
		cache: map[string]*CacheVal{},
		mutex: new(sync.RWMutex),
	}
}

func (c *Cache) Get(key string) *CacheVal {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	val := c.cache[key]
	return val
}

func (c *Cache) Set(key string, val *CacheVal) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = val
}
