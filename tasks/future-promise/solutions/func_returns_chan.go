package solutions

import "fmt"

// return channel that will contain fibonacci value
func asyncFibonacci(value int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		ch <- fibonacciCalculate(value) // not blocked, because channel is BUFFERED and empty
		close(ch)
	}()

	return ch
}

func FutureWithReturnedChan() {
	outputCh := asyncFibonacci(40)
	value := <-outputCh
	fmt.Printf("fibonacci of %d: %d\n", 40, value)
}
