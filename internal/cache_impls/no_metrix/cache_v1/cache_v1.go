package cache_v1

import (
	"github.com/alexeykirinyuk/async-cache/internal/helpers"
)

type CacheImpl struct {
	m map[string]string
}

func NewCacheImpl() *CacheImpl {
	return &CacheImpl{
		m: make(map[string]string, 100),
	}
}

func (c *CacheImpl) Set(key, value string) error {
	c.m[key] = value

	return nil
}

func (c *CacheImpl) Get(key string) (string, error) {
	val, ok := c.m[key]
	if !ok {
		return "", helpers.ErrNotFound
	}

	return val, nil
}

func (c *CacheImpl) Delete(key string) error {
	_, ok := c.m[key]
	if !ok {
		return helpers.ErrNotFound
	}

	delete(c.m, key)

	return nil
}
