package cache_v2

import (
	"sync"

	"github.com/alexeykirinyuk/async-cache/internal/storage"
)

type CacheImpl struct {
	m   map[string]string
	mtx *sync.RWMutex
}

func NewCacheImpl() *CacheImpl {
	return &CacheImpl{
		m:   make(map[string]string, 100),
		mtx: &sync.RWMutex{},
	}
}

func (c *CacheImpl) Set(key, value string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.m[key] = value

	return nil
}

func (c *CacheImpl) Get(key string) (string, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	val, ok := c.m[key]
	if !ok {
		return "", storage.ErrNotFound
	}

	return val, nil
}

func (c *CacheImpl) Delete(key string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	_, ok := c.m[key]
	if !ok {
		return storage.ErrNotFound
	}

	delete(c.m, key)

	return nil
}
