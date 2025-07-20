package solutions

import (
	"context"
	"fmt"
	"sync"
)

func FewProducers() {
	const (
		numProducers = 2   // Num of Producers
		numConsumers = 1   // Num of Consumers
		totalNumbers = 100 // Num of generating values
		bufferSize   = 100 // the size of channel's buffer
	)

	bufChan := make(chan int, bufferSize)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// few producers
	for i := 0; i < numProducers; i++ {
		producerWg.Add(1)

		go func() {
			defer producerWg.Done()
			for i := 0; i < totalNumbers/numProducers; i++ {
				val := randIntRange(1, 1000)
				select {
				case bufChan <- val:
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	// consumer
	for i := 0; i < numConsumers; i++ {
		consumerWg.Add(1)
		go func() {
			defer consumerWg.Done()
			idx := 1
			for val := range bufChan {
				fmt.Printf("%d. %d^2 = %d\n", idx, val, pow(val))
				idx++
			}
		}()
	}

	// waiting for producers to finish
	// because we need to trigger channel's closing
	go func() {
		producerWg.Wait()
		close(bufChan)
	}()

	// waiting for consumer to finish
	consumerWg.Wait()
}
