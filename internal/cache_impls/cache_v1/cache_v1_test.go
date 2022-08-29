package cache_v1

import (
	"fmt"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	c := NewCacheImpl()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)
		val := fmt.Sprintf("val_%d", i)

		err := c.Set(key, val)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	c := NewCacheImpl()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)
		val := fmt.Sprintf("val_%d", i)

		err := c.Set(key, val)
		if err != nil {
			b.Error(err)
		}
	}

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)

		_, err := c.Get(key)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	c := NewCacheImpl()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)
		val := fmt.Sprintf("val_%d", i)

		err := c.Set(key, val)
		if err != nil {
			b.Error(err)
		}
	}

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)

		err := c.Delete(key)
		if err != nil {
			b.Error(err)
		}
	}
}
