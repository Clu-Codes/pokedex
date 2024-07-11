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
		fmt.Printf("cache entry: %v added successfully", keyData)
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
	ca.reapLoop()
	return ca
}

func (ca *Cache) reapLoop() {
	tick := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-tick.C:
				for key, val := range ca.CacheMap {
					if time.Duration(time.Since(val.createdAt).Seconds()) > ca.duration {
						delete(ca.CacheMap, key)
					}
				}
			}
		}
	}()

	time.Sleep(2 * time.Minute) // safeguard in place until I better understand how to close channel
	done <- true

}
