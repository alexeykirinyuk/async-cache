package test

import (
	"testing"

	"github.com/alexeykirinyuk/async-cache/internal/cache_impls/with_metrix/with_atomic"
	"github.com/stretchr/testify/assert"
)

func Test_CacheWithAtomicMetrics(t *testing.T) {
	t.Parallel()

	testCache := with_atomic.NewCacheImpl()

	t.Run("correctly stored value", func(t *testing.T) {
		t.Parallel()
		key := "test_key"
		value := "test_value"

		err := testCache.Set(key, value)
		assert.NoError(t, err)

		storedVal, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedVal)
	})

	t.Run("correctly updated value", func(t *testing.T) {
		t.Parallel()
		key := "test_key"
		value := "test_value"

		err := testCache.Set(key, value)
		assert.NoError(t, err)

		storedVal, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedVal)

		value2 := "test_value_2"

		err = testCache.Set(key, value2)
		assert.NoError(t, err)

		storedVal, err = testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value2, storedVal)
	})

	t.Run("correctly deleted value", func(t *testing.T) {
		t.Parallel()
		key := "test_key"
		value := "test_value"

		err := testCache.Set(key, value)
		assert.NoError(t, err)

		storedVal, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedVal)

		value2 := "test_value_2"

		err = testCache.Set(key, value2)
		assert.NoError(t, err)

		storedVal, err = testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value2, storedVal)

		err = testCache.Delete(key)
		assert.NoError(t, err)
	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()

		parallelFactory := 100_000

		emulateLoad(t, testCache, parallelFactory)
	})

	t.Run("no data races with metrics", func(t *testing.T) {
		t.Parallel()

		parallelFactory := 100_000

		emulateLoadWithMetrix(t, testCache, parallelFactory)
	})
}
