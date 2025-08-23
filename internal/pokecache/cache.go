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
	mtx     sync.Mutex
	entries map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	fmt.Println("CACHE CREATED")
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if entry, exists := c.entries[key]; exists {
		fmt.Println("CACHE HIT")
		return entry.val, true
	}

	fmt.Println("CACHE MISS")
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)

		c.mtx.Lock()

		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, key)
				fmt.Println("CACHE REAPED AT: ", time.Now(), ", key: ", key)
			}
		}
		c.mtx.Unlock()
	}
}
