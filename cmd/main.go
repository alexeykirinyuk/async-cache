package main

import (
	"fmt"

	cache_v2 "github.com/alexeykirinyuk/async-cache/internal/cache_impls/cache_v2_"
	"github.com/alexeykirinyuk/async-cache/internal/storage"
)

func main() {
	var c storage.Cache = cache_v2.NewCacheImpl()

	c.Set("test", "test-val")
	val, err := c.Get("test")
	fmt.Println(val, err)
}
