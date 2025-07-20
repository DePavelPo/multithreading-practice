package solutions

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func UnbufferedChannel() {
	unbufChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	// producer (put rand int into channel)
	go func() {
		defer wg.Done()
		defer close(unbufChan) // close channel after finish sending to the channel
		for i := 0; i < 100; i++ {
			unbufChan <- randIntRange(1, 1000)
		}
	}()

	wg.Add(1)
	// consumer (read channel and output the square of received int)
	go func() {
		defer wg.Done()
		idx := 1
		for val := range unbufChan { // getting channel's data. Automatically finishes if close() is called
			fmt.Printf("%d. %d^2 = %d\n", idx, val, pow(val))
			idx++
		}
	}()

	wg.Wait()
}

func pow(n int) int {
	return n * n
}

func randIntRange(min, max int) int {
	return rand.IntN(max-min) + min
}
