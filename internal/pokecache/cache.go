package pokecache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	CacheMap map[string]cacheEntry
	mu       sync.Mutex
	duration time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (ca *Cache) AddCache(key string, data []byte) error {
	ca.mu.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
	ca.CacheMap[key] = entry
	ca.mu.Unlock()

	if keyData, ok := ca.CacheMap[key]; ok {
		fmt.Println("cache entry: %v added successfully", keyData)
	} else {
		return errors.New("Failed to add cache entry to Cache")

	}

	return nil
}

func NewCache(interval time.Duration) *Cache {
	ca := &Cache{
		CacheMap: make(map[string]cacheEntry),
		duration: interval,
	}
	// TODO: Create reapLoop() to go here
}

// create reapLoop()
