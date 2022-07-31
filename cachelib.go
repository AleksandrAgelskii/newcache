package newcache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	cache map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()
	c.cache[key] = value
	go func() {
		<-time.NewTimer(ttl).C
		c.Delete(key)
	}()
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()
	_, ok := c.cache[key]
	if !ok {
		return nil, errors.New("data not found")
	}
	return c.cache[key], nil
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()
	_, ok := c.cache[key]
	if !ok {
		return errors.New("data not found")
	}
	delete(c.cache, key)
	return nil
}
