package solutions

import (
	"fmt"
	"sync"
)

type CounterRW struct {
	Value int
	mu    sync.RWMutex
}

func (c *CounterRW) Int() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value++
}

func (c *CounterRW) Dec() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value--
}

func (c *CounterRW) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Value
}

func WithRWMutex() {
	var wg sync.WaitGroup

	counter := CounterRW{
		Value: 0,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			counter.Int()
			// fmt.Println("After inc value:", counter.GetValue())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			counter.Dec()
			// fmt.Println("After dec value:", counter.GetValue())
		}
	}()

	wg.Wait()
	fmt.Println("result:", counter.GetValue())
}
