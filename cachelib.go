package newcache

//newli
import (
	"errors"
)

type Cache struct {
	cache map[string]interface{}
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value interface{}) {
	c.cache[key] = value
}

func (c Cache) Get(key string) (interface{}, error) {
	_, ok := c.cache[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return c.cache[key], nil
}

func (c Cache) Delete(key string) error {
	_, ok := c.cache[key]
	if !ok {
		return errors.New("key not found")
	}
	delete(c.cache, key)
	return nil
}
