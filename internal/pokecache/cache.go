package pokecache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	CacheMap map[string]cacheEntry
	mu       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (ca *Cache) AddCache(key string, data []byte) error {
	ca.mu.Lock()
	ca.CacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
	ca.mu.Unlock()

	if keyData, ok := ca.CacheMap[key]; ok {
		fmt.Printf("%v: cache entry added successfully \n", keyData.createdAt)
	} else {
		return errors.New("failed to add cache entry to cache")

	}

	return nil
}

func (ca *Cache) GetCache(key string) ([]byte, bool) {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	v, ok := ca.CacheMap[key]
	if ok {
		fmt.Println("Found cache...")
	}
	return v.val, ok
}

func NewCache(interval time.Duration) Cache {
	ca := Cache{
		CacheMap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}
	go ca.reapLoop(interval)
	return ca
}

func (ca *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		ca.mu.Lock()
		for key, entry := range ca.CacheMap {
			if time.Since(entry.createdAt) > interval {
				delete(ca.CacheMap, key)
			}
		}
		ca.mu.Unlock()
	}

}
