package solutions

import (
	"context"
	"fmt"
	"sync"
)

func BufferedChannel() {
	bufChan := make(chan int, 100)
	ctx, cancel := context.WithCancel(context.Background()) // cancel context by some triggers
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	// producer (put rand int into channel)
	go func() {
		defer wg.Done()
		defer close(bufChan)
		for i := 0; i < 100; i++ {
			val := randIntRange(1, 1000)
			select {
			case bufChan <- val:
			case <-ctx.Done(): // helps to finish goroutine if context was canceled
				return
			}
		}
	}()

	wg.Add(1)
	// consumer (read channel and output the square of received int)
	go func() {
		defer wg.Done()
		idx := 1
		for { // finishes by some triggers
			select {
			case val, ok := <-bufChan:
				if !ok { // finish goroutine if channel is empty AND closed
					return
				}
				fmt.Printf("%d. %d^2 = %d\n", idx, val, pow(val))
				idx++
			case <-ctx.Done(): // helps to finish goroutine if context was canceled
				return
			}
		}
	}()

	wg.Wait()
}
