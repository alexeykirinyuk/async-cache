package main

import (
	"fmt"

	"github.com/alexeykirinyuk/async-cache/internal/cache_impls/cache_v3"
	"github.com/alexeykirinyuk/async-cache/internal/storage"
)

func main() {
	var c storage.Cache = cache_v3.NewCacheImpl()

	c.Set("test", "test-val")
	val, err := c.Get("test")
	fmt.Println(val, err)
}
