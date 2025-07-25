package solutions

import "fmt"

func fibonacciCalculate(value int) int {
	if value <= 1 {
		return value
	}

	return fibonacciCalculate(value-2) + fibonacciCalculate(value-1)
}

func FuturePromise() {
	outputCh := make(chan int)
	go func() {
		defer close(outputCh)
		fibonacciNumber := fibonacciCalculate(30)
		outputCh <- fibonacciNumber
	}()

	value := <-outputCh
	fmt.Printf("fibonacci of %d: %d\n", 30, value)
}
