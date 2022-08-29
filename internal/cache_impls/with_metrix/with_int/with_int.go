package with_int

import (
	"sync"
	"sync/atomic"

	"github.com/alexeykirinyuk/async-cache/internal/helpers"
)

type CacheImpl struct {
	m   map[string]string
	t   int64
	mtx *sync.RWMutex
}

func NewCacheImpl() *CacheImpl {
	return &CacheImpl{
		m:   make(map[string]string, 100),
		t:   int64(0),
		mtx: &sync.RWMutex{},
	}
}

func (c *CacheImpl) Set(key, value string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.m[key] = value
	c.t++

	return nil
}

func (c *CacheImpl) Get(key string) (string, error) {
	atomic.AddInt64(&c.t, 1)

	c.mtx.RLock()
	defer c.mtx.RUnlock()

	val, ok := c.m[key]
	if !ok {
		return "", helpers.ErrNotFound
	}

	return val, nil
}

func (c *CacheImpl) Delete(key string) error {
	atomic.AddInt64(&c.t, 1)

	c.mtx.Lock()
	defer c.mtx.Unlock()

	_, ok := c.m[key]
	if !ok {
		return helpers.ErrNotFound
	}

	delete(c.m, key)
	c.t--

	return nil
}

func (c *CacheImpl) TotalAmount() int64 {
	return c.t
}
