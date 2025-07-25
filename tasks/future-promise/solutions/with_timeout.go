package solutions

import (
	"context"
	"fmt"
	"time"
)

// return fibonacci value or error if context was cancelled
func futureWithContext(ctx context.Context, value int) (int, error) {
	ch := make(chan int, 1)

	go func() {
		ch <- fibonacciCalculate(value)
		close(ch)
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func FutureWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	value, err := futureWithContext(ctx, 40)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("fibonacci of %d: %d\n", 40, value)
}
