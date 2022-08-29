package cache_v3

import (
	"sync"

	"github.com/alexeykirinyuk/async-cache/internal/helpers"
)

type CacheImpl struct {
	m   map[string]string
	mtx *sync.Mutex
}

func NewCacheImpl() *CacheImpl {
	return &CacheImpl{
		m:   make(map[string]string, 100),
		mtx: &sync.Mutex{},
	}
}

func (c *CacheImpl) Set(key, value string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.m[key] = value

	return nil
}

func (c *CacheImpl) Get(key string) (string, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	val, ok := c.m[key]
	if !ok {
		return "", helpers.ErrNotFound
	}

	return val, nil
}

func (c *CacheImpl) Delete(key string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	_, ok := c.m[key]
	if !ok {
		return helpers.ErrNotFound
	}

	delete(c.m, key)

	return nil
}
