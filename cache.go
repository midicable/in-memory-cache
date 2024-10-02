package cache

import (
	"errors"
)

var ErrNotFound = errors.New("key does not exist")

type Cache struct {
	store map[string]any
}

func NewCache() *Cache {
	return &Cache{store: make(map[string]any, 0)}
}

func (c *Cache) Get(key string) (any, error) {
	value, ok := c.store[key]

	if !ok {
		return nil, ErrNotFound
	}

	return value, nil
}

func (c *Cache) Set(key string, value any) {
	c.store[key] = value
}

func (c *Cache) Delete(key string) {
	delete(c.store, key)
}
