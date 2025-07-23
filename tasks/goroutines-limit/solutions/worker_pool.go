package solutions

import (
	"fmt"
	"sync"
)

func fibonacciWorker(
	wg *sync.WaitGroup,
	inputCh <-chan int,
	outputCh chan<- fibonacci,
) {
	defer wg.Done()

	for number := range inputCh {
		fibonacciNumber := fibonacciCalculate(number)
		outputCh <- fibonacci{
			Number: number,
			Result: fibonacciNumber,
		}
	}
}

func WorkerPoolUsing(totalNumber, goCount int) {
	chanForGenerating := make(chan int, totalNumber)
	// generate random number
	go func() {
		defer close(chanForGenerating)
		for i := 0; i < totalNumber; i++ {
			chanForGenerating <- randInt()
		}
	}()

	outputCh := make(chan fibonacci, totalNumber)
	wg := &sync.WaitGroup{}
	// make fixed num of goroutines with calling fibonacci calculating
	// !!! we create goCount goroutines and DONT CREATE ANY OTHER AFTER. The number of goroutines is FIXED !!!
	go func() {
		for i := 0; i < goCount; i++ {
			wg.Add(1)

			go fibonacciWorker(wg, chanForGenerating, outputCh)
		}
		wg.Wait()
		close(outputCh) // close output channel when all goroutines were done
	}()

	// output calculated numbers
	for output := range outputCh {
		// _ = output
		fmt.Printf("fibonacci of %d = %d\n", output.Number, output.Result)
	}
}
