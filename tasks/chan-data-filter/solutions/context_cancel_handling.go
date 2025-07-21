package solutions

import (
	"context"
	"fmt"
	"sync"
)

func ContextCancelHandling(totalNumbers int) {
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
		for {
			select {
			case val, ok := <-a:
				if !ok {
					return
				}
				if isEven(val) {
					select {
					case b <- val:
					case <-ctx.Done():
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Add(1)
	// doing x2
	go func() {
		defer wg.Done()
		defer close(c)
		for {
			select {
			case val, ok := <-b:
				if !ok {
					return
				}
				select {
				case c <- val * 2:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Add(1)
	// output numbers
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-c:
				if !ok {
					return
				}
				fmt.Println("Got value:", val)
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
}
