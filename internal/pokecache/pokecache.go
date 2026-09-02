/* SPDX-License-Identifier: GPL-3.0-or-later */
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	interval time.Duration
	caches   map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		interval: interval,
		caches:   make(map[string]cacheEntry),
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.caches[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.caches[key]
	if exists {
		return entry.val, exists
	} else {
		return nil, exists
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.caches {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.caches, key)
			}
		}
		c.mu.Unlock()
	}
}
