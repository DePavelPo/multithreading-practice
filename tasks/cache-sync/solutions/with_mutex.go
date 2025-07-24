package solutions

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

type cache struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCache() *cache {
	return &cache{
		data: make(map[string]string),
	}
}

// get value from cache
func (c *cache) Get(key string) (string, bool) {
	c.mu.RLock()
	value, ok := c.data[key]
	c.mu.RUnlock()
	return value, ok
}

// set value in cahce
func (c *cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func SyncWithMutex(producersCount, consumersCount int) {
	cache := NewCache()

	wg := sync.WaitGroup{}
	// concurrently trying to set value in cache
	for i := 0; i < producersCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			keyStr := fmt.Sprintf("key%d", i)
			valueStr := fmt.Sprintf("value%d", i)
			cache.Set(keyStr, valueStr)

		}(i)
	}

	// concurrently trying to get value from cache
	for j := 0; j < consumersCount; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			// get a random key depend on generated keys
			keyStr := fmt.Sprintf("key%d", randInt(producersCount))
			_, _ = cache.Get(keyStr)
			// if !ok {
			// 	fmt.Println("Not found")
			// } else {
			// 	fmt.Printf("%s: %s\n", keyStr, value)
			// }
		}(j)
	}

	wg.Wait()
}

func randInt(max int) int {
	return rand.IntN(max)
}
