package solutions

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type CounterAtomic struct {
	value int32
}

func (c *CounterAtomic) Inc() {
	atomic.AddInt32(&c.value, 1)
}

// decriment with condition "value cant be negative"
func (c *CounterAtomic) Dec() {
	for { // waiting till value is decremented
		val := atomic.LoadInt32(&c.value)
		if val == 0 { // waiting while value is zero
			time.Sleep(time.Millisecond) // too long, but there is no cond-like func in atomic
			continue
		}
		if atomic.CompareAndSwapInt32(&c.value, val, val-1) {
			break
		}
	}
}

func (c *CounterAtomic) GetValue() int32 {
	return atomic.LoadInt32(&c.value)
}

func WithAtomic() {
	var wg sync.WaitGroup
	counter := &CounterAtomic{
		value: 0,
	}

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
