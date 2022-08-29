package main

import (
	"fmt"

	"github.com/alexeykirinyuk/async-cache/internal/storage"
)

func main() {
	var c storage.Cache = storage.NewCacheImpl()

	c.Set("test", "test-val")
	val, err := c.Get("test")
	fmt.Println(val, err)
}
