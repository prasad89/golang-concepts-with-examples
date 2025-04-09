// https://go.dev/play/p/A4OzZDjhwCM

//go:build ignore

package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu    sync.RWMutex
	items map[string]string
}

func (c *Cache) Set(key, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = val
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, exist := c.items[key]
	return val, exist
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func main() {
	userData := &Cache{items: make(map[string]string)}

	userData.Set("name", "prasad")

	if val, exist := userData.Get("name"); exist {
		fmt.Println(val)
	} else {
		fmt.Println("Not found")
	}

	userData.Delete("name")

	if val, exist := userData.Get("name"); exist {
		fmt.Println(val)
	} else {
		fmt.Println("Not found")
	}
}
