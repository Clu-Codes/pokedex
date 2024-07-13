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

func (ca *Cache) AddCache(key *string, data []byte) error {
	ca.mu.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
	url := "https://pokeapi.co/api/v2/location-area"
	if key != nil {
		url = *key
	}
	ca.CacheMap[url] = entry
	ca.mu.Unlock()

	if keyData, ok := ca.CacheMap[url]; ok {
		fmt.Printf("%v: cache entry added successfully \n", keyData.createdAt)
	} else {
		return errors.New("failed to add cache entry to cache")

	}

	return nil
}

func (ca *Cache) GetCache(key string) ([]byte, bool) {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	if v, ok := ca.CacheMap[key]; ok {
		fmt.Println("Found cache...")
		return v.val, true
	} else {
		return nil, false
	}
}

func NewCache(interval time.Duration) *Cache {
	ca := &Cache{
		CacheMap: make(map[string]cacheEntry),
		duration: interval,
	}
	go ca.reapLoop()
	return ca
}

func (ca *Cache) reapLoop() {
	ticker := time.NewTicker(ca.duration)
	for range ticker.C {
		ca.mu.Lock()
		for key, entry := range ca.CacheMap {
			if time.Since(entry.createdAt) > ca.duration {
				delete(ca.CacheMap, key)
			}
		}
		ca.mu.Unlock()
	}

}
