package solutions

import (
	"fmt"
	"sync"
)

type Result struct {
	Number int
	Square int
}

func BufferedChan() {
	numbers := make([]int, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = i + 1
	}

	var wg sync.WaitGroup
	results := make(chan Result, 100)
	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- Result{n, pow(n)}
		}(number)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	resultMap := make(map[int]int)
	for result := range results {
		resultMap[result.Number] = result.Square
	}

	fmt.Println("result len:", len(resultMap))

	for i := 1; i <= len(resultMap); i++ {
		fmt.Printf("%d^2 = %d\n", i, resultMap[i])
	}
}
