//go:build ignore

package main

import (
	"fmt"
	"time"
)

type Cache struct {
	items  map[string]float32
	expiry map[string]time.Time
}

func NewCache() *Cache {
	return &Cache{
		items:  make(map[string]float32),
		expiry: make(map[string]time.Time),
	}
}

func (c *Cache) Get(key string) float32 {
	expTime, exists := c.expiry[key]
	if !exists || time.Now().After(expTime) {
		delete(c.items, key)
		delete(c.expiry, key)
		return -1.0
	}
	return c.items[key]
}

func (c *Cache) Put(key string, val float32) {
	c.items[key] = val
	c.expiry[key] = time.Now().Add(10 * time.Second)
}

func main() {
	cache := NewCache()
	cache.Put("temperature", 98.6)

	fmt.Println("First get:", cache.Get("temperature"))

	time.Sleep(11 * time.Second)

	fmt.Println("Second get (after expiration):", cache.Get("temperature"))
}
