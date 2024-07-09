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

type Ticker struct {
	C <-chan time
}

func NewTicker(d time.Duration) *Ticker {
	ticker := &Ticker{
		C: d,
	}
	return ticker
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

func (ca *Cache) reapLoop(interval time.Duration) {
	tick := time.NewTicker(5 * time.Second)
	// TODO: Better understand how to create a Ticker.
	// Determine how the ticker will run -- should I use infinite for loop?
	// How do I determine when to stop the ticker / how do I stop the ticker?
	// I need to also add the logic that will handle the clearing of cache if it is greater than the interval.
}
