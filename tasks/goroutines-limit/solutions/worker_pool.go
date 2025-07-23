package solutions

import (
	"fmt"
	"sync"
)

func fibanacciWorker(
	wg *sync.WaitGroup,
	inputCh <-chan int,
	outputCh chan<- fibanacci,
) {
	defer wg.Done()

	for number := range inputCh {
		fibanacciNumber := fibanacciCalculate(number)
		outputCh <- fibanacci{
			Number: number,
			Result: fibanacciNumber,
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

	outputCh := make(chan fibanacci, totalNumber)
	wg := &sync.WaitGroup{}
	// make fixed num of goroutines with calling fibanacci calculating
	go func() {
		for i := 0; i < goCount; i++ {
			wg.Add(1)

			go fibanacciWorker(wg, chanForGenerating, outputCh)
		}
		wg.Wait()
		close(outputCh) // close output channel when all goroutines were done
	}()

	// output calculated numbers
	for output := range outputCh {
		// _ = output
		fmt.Printf("fibanacci of %d = %d\n", output.Number, output.Result)
	}
}
