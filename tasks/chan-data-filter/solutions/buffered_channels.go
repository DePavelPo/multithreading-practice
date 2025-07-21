package solutions

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
)

func DataFilter(totalNumbers int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		wg sync.WaitGroup
		a  = make(chan int, 100)
		b  = make(chan int, 100)
		c  = make(chan int, 100)
	)

	wg.Add(1)
	// making random numbers
	go func() {
		defer wg.Done()
		defer close(a)
		for i := 0; i < totalNumbers; i++ {
			select {
			case a <- randIntRange(1, 1000):
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Add(1)
	// choosing even numbers
	go func() {
		defer wg.Done()
		defer close(b)
		for val := range a {
			if isEven(val) {
				b <- val
			}
		}
	}()

	wg.Add(1)
	// doing x2
	go func() {
		defer wg.Done()
		defer close(c)
		for val := range b {
			c <- val * 2
		}
	}()

	wg.Add(1)
	// output numbers
	go func() {
		defer wg.Done()
		for val := range c {
			fmt.Println("Got value:", val)
		}
	}()

	wg.Wait()
}

func randIntRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func isEven(n int) bool {
	return n&1 == 0
}
