package solutions

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

type semaphore struct {
	Channel chan struct{}
}

func (s *semaphore) Acquire() {
	s.Channel <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.Channel
}

type fibonacci struct {
	Number int
	Result int
}

func fibonacciCalculate(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciCalculate(n-2) + fibonacciCalculate(n-1)
}

func SemaphoreUsing(totalNumber, goCount int) {
	sem := semaphore{
		Channel: make(chan struct{}, goCount),
	}

	chanForGenerating := make(chan int, totalNumber)
	// generate random number
	go func() {
		defer close(chanForGenerating)
		for i := 0; i < totalNumber; i++ {
			chanForGenerating <- randInt()
		}
	}()

	var wg sync.WaitGroup
	outputCh := make(chan fibonacci, totalNumber)
	// calculating fibonacci numbers with goroutines limit by semaphore
	// !!! we finish goroutines and create new goroutines but ONLY goCount goroutines work concurently !!!
	for number := range chanForGenerating {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()

			fibonacciNumber := fibonacciCalculate(num)

			outputCh <- fibonacci{
				Number: num,
				Result: fibonacciNumber,
			}
		}(number)
	}

	// needs to close output channel
	go func() {
		wg.Wait()
		close(outputCh)
	}()

	// output calculated numbers
	for output := range outputCh {
		// _ = output
		fmt.Printf("fibonacci of %d = %d\n", output.Number, output.Result)
	}
}

func randInt() int {
	return rand.IntN(40)
}
