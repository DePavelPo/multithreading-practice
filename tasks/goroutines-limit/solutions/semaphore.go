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

type fibanacci struct {
	Number int
	Result int
}

func fibanacciCalculate(n int) int {
	if n <= 1 {
		return n
	}

	return fibanacciCalculate(n-2) + fibanacciCalculate(n-1)
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
	outputCh := make(chan fibanacci, totalNumber)
	// calculating fibanacci numbers with goroutines limit by semaphore
	for number := range chanForGenerating {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()

			fibanacciNumber := fibanacciCalculate(num)

			outputCh <- fibanacci{
				Number: num,
				Result: fibanacciNumber,
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
		fmt.Printf("fibanacci of %d = %d\n", output.Number, output.Result)
	}
}

func randInt() int {
	return rand.IntN(40)
}
