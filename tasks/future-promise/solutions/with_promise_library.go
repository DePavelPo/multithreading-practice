package solutions

import (
	"context"
	"fmt"
	"time"

	"github.com/chebyrash/promise"
)

func FutureWithLibrary() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	p := promise.New(func(resolve func(int), reject func(error)) {
		resolve(fibonacciCalculate(40)) // calling the fibonacci calculation
	})

	value, err := p.Await(ctx) // return err if ctx was cancelled or if we get an exception
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if value == nil {
		fmt.Println("Empty result")
		return
	}
	fmt.Printf("fibonacci of %d: %d\n", 40, *value)
}
