package bench

import (
	"testing"

	"github.com/alexeykirinyuk/async-cache/internal/cache_impls/no_metrix/cache_v1"
	"github.com/alexeykirinyuk/async-cache/internal/cache_impls/no_metrix/cache_v2"
	"github.com/alexeykirinyuk/async-cache/internal/cache_impls/no_metrix/cache_v3"
)

const parallelFactor = 10_000

func Benchmark_NoMutexBalanced(b *testing.B) {
	b.Skip("it's not working")

	c := cache_v1.NewCacheImpl()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, parallelFactor)
	}
}

// Benchmark_RWMutex-8, 79, 15008872 ns/op
func Benchmark_RWMutexBalanced(b *testing.B) {
	c := cache_v2.NewCacheImpl()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, parallelFactor)
	}
}

// Benchmark_Mutex-8, 135, 8690496 ns/op
func Benchmark_MutexBalanced(b *testing.B) {
	c := cache_v3.NewCacheImpl()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, parallelFactor)
	}
}

// Benchmark_RWMutexReadIntersive-8, 309, 3819435 ns/op
func Benchmark_RWMutexReadIntersive(b *testing.B) {
	c := cache_v2.NewCacheImpl()

	for i := 0; i < b.N; i++ {
		emulateReadIntersiveLoad(c, parallelFactor)
	}
}

// Benchmark_MutexReadIntersive-8, 310, 3816771 ns/op
func Benchmark_MutexReadIntersive(b *testing.B) {
	c := cache_v3.NewCacheImpl()

	for i := 0; i < b.N; i++ {
		emulateReadIntersiveLoad(c, parallelFactor)
	}
}
