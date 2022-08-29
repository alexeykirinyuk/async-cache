package bench

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alexeykirinyuk/async-cache/internal/helpers"
	"github.com/alexeykirinyuk/async-cache/internal/storage"
)

func emulateLoad(c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("#%d-key", i)
		val := fmt.Sprintf("#%d-value", i)

		wg.Add(1)
		go func(key, value string) {
			err := c.Set(key, val)
			if err != nil {
				panic(err)
			}

			wg.Done()
		}(key, val)

		wg.Add(1)
		go func(key, val string) {
			_, err := c.Get(key)
			if err != nil && !errors.Is(err, helpers.ErrNotFound) {
				panic(err)
			}

			wg.Done()
		}(key, val)

		wg.Add(1)
		go func(key, val string) {
			err := c.Delete(key)

			if err != nil && !errors.Is(err, helpers.ErrNotFound) {
				panic(err)
			}

			wg.Done()
		}(key, val)
	}

	wg.Wait()
}

func emulateReadIntersiveLoad(c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor/10; i++ {
		key := fmt.Sprintf("#%d-key", i)
		val := fmt.Sprintf("#%d-value", i)

		wg.Add(1)
		go func(key, value string) {
			err := c.Set(key, val)
			if err != nil {
				panic(err)
			}

			wg.Done()
		}(key, val)

		wg.Add(1)
		go func(key, val string) {
			err := c.Delete(key)

			if err != nil && !errors.Is(err, helpers.ErrNotFound) {
				panic(err)
			}

			wg.Done()
		}(key, val)
	}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("#%d-key", i)
		val := fmt.Sprintf("#%d-value", i)

		wg.Add(1)
		go func(key, val string) {
			_, err := c.Get(key)
			if err != nil && !errors.Is(err, helpers.ErrNotFound) {
				panic(err)
			}

			wg.Done()
		}(key, val)
	}

	wg.Wait()
}
