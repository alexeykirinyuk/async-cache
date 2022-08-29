package test

import (
	"testing"

	"github.com/alexeykirinyuk/async-cache/internal/cache_impls/no_metrix/cache_v2"
	"github.com/stretchr/testify/assert"
)

func Test_Cache_V2(t *testing.T) {
	t.Parallel()

	testCache := cache_v2.NewCacheImpl()

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

}
