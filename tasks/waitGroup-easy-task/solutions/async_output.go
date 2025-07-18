package solutions

import (
	"fmt"
	"sync"
)

func RandomOutput() {
	numbers := make([]int, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = i + 1
	}

	var (
		wg      sync.WaitGroup
		mu      sync.RWMutex
		results = make(map[int]int)
	)

	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			square := pow(n)

			// without mutex cause of data race
			mu.Lock()
			results[n] = square
			mu.Unlock()
		}(number)
	}
	wg.Wait()

	// async output
	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			mu.RLock()
			val := results[n]
			mu.RUnlock()

			fmt.Printf("%d^2 = %d\n", n, val)
		}(number)
	}
	wg.Wait()
}

func pow(a int) int {
	return a * a
}
