package bench

import (
	"testing"

	"github.com/alexeykirinyuk/async-cache/internal/cache_impls/with_metrix/with_atomic"
	"github.com/alexeykirinyuk/async-cache/internal/cache_impls/with_metrix/with_int"
)

// Benchmark_WithMetrics_Atomic-8 - 265, 4293433 ns/op
func Benchmark_WithMetrics_Atomic(b *testing.B) {
	c := with_atomic.NewCacheImpl()

	for i := 0; i < b.N; i++ {
		emulateReadIntersiveLoad(c, parallelFactor)
	}
}

// Benchmark_WithMetrics_Int-8 - 267, 4427924 ns/op
func Benchmark_WithMetrics_Int(b *testing.B) {
	c := with_int.NewCacheImpl()

	for i := 0; i < b.N; i++ {
		emulateReadIntersiveLoad(c, parallelFactor)
	}
}
