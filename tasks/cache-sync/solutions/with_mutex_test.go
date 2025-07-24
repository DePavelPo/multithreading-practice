package solutions_test

import (
	"fmt"
	"testing"
	"time"

	sl "github.com/DePavelPo/multithreading-practice/tasks/cache-sync/solutions"
)

var cases = []struct {
	producersCount int
	consumersCount int
}{
	{producersCount: 10, consumersCount: 1000},
	{producersCount: 50, consumersCount: 1000},
	{producersCount: 100, consumersCount: 1000},
	{producersCount: 100, consumersCount: 5000},
}

func BenchmarkSyncWithMutex(b *testing.B) {
	for _, cs := range cases {
		b.Run(fmt.Sprintf("producers_%d_consumers_%d", cs.producersCount, cs.consumersCount), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sl.SyncWithMutex(cs.producersCount, cs.consumersCount)
			}
		})
	}
}

func TestCache_Deadlock(t *testing.T) {
	c := sl.NewCache()
	done := make(chan bool)

	go func() {
		c.Set("key1", "value1")
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Error("Possible deadlock detected")
	}
}
