package solutions

import (
	"fmt"
	"sync"
)

type CounterCond struct {
	value int
	mu    sync.RWMutex
	cond  *sync.Cond
}

func NewCounterCond() *CounterCond {
	counter := &CounterCond{}
	counter.cond = sync.NewCond(&counter.mu)
	return counter
}

func (c *CounterCond) Inc() {
	c.mu.Lock()
	c.value++
	c.cond.Signal() // sends signal to one sleepy goroutine
	c.mu.Unlock()
}

// decriment with condition "value cant be negative"
func (c *CounterCond) Dec() {
	c.mu.Lock()
	for c.value == 0 { // needs to recheck condition
		c.cond.Wait() // unlock locker, sleep until signal, lock locker
	}
	c.value--
	c.mu.Unlock()
}

func (c *CounterCond) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func WithCond() {
	counter := NewCounterCond()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			counter.Dec()
			// fmt.Println("After dec value:", counter.GetValue())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			counter.Inc()
			// fmt.Println("After inc value:", counter.GetValue())
		}
	}()

	wg.Wait()
	fmt.Println("result:", counter.GetValue())
}
