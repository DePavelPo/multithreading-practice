package solutions

import (
	"fmt"
	"sync"
)

func SyncOutput() {
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

	// sync output
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d^2 = %d\n", i, results[i])
	}
}
