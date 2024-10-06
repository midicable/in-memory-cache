package cache

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("key does not exist")

type Cache struct {
	store map[string]any
	mu    *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{store: make(map[string]any, 0), mu: &sync.RWMutex{}}
}

func (c *Cache) Get(key string) (any, error) {
	value, ok := c.store[key]

	if !ok {
		return nil, ErrNotFound
	}

	return value, nil
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	c.store[key] = value
	c.mu.Unlock()
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.store, key)
	c.mu.Unlock()
}
