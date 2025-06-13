package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entry: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("Nouvelle entrée dans cache %s\n", key)
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	// fmt.Print("New cache entry : ")
	// fmt.Println(c.entry[key])
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("GET entrée dans cache %s\n", key)
	entry, exists := c.entry[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			c.mu.Lock()
			for key, val := range c.entry {
				if t.After(val.createdAt.Add(interval)) {
					delete(c.entry, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
