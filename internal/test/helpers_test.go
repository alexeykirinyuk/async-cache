package test

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/alexeykirinyuk/async-cache/internal/helpers"
	"github.com/alexeykirinyuk/async-cache/internal/storage"
	"github.com/stretchr/testify/assert"
)

func emulateLoad(t *testing.T, c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("#%d-key", i)
		val := fmt.Sprintf("#%d-value", i)

		wg.Add(1)
		// write to cache
		go func(key, value string) {
			err := c.Set(key, val)
			assert.NoError(t, err)

			wg.Done()
		}(key, val)

		wg.Add(1)
		go func(key, val string) {
			actualVal, err := c.Get(key)

			if !errors.Is(err, helpers.ErrNotFound) {
				assert.Equal(t, val, actualVal)
			}

			wg.Done()
		}(key, val)

		wg.Add(1)
		go func(key, val string) {
			err := c.Delete(key)

			if !errors.Is(err, helpers.ErrNotFound) {
				assert.NoError(t, err)
			}

			wg.Done()
		}(key, val)
	}

	wg.Wait()
}

func emulateLoadWithMetrix(t *testing.T, c storage.CacheWithMetrics, parallelFactor int) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		emulateLoad(t, c, parallelFactor)
		wg.Done()
	}()

	var min, max int64

	for i := 0; i < parallelFactor; i++ {
		wg.Add(1)
		go func() {
			total := c.TotalAmount()

			if total > max {
				max = total
			}
			if total < min {
				min = total
			}

			wg.Done()
		}()
	}

	t.Logf("min - %d, max - %d", min, max)

	wg.Wait()
}
