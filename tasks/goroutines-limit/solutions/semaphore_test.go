package solutions_test

import (
	"fmt"
	"testing"

	sl "github.com/DePavelPo/multithreading-practice/tasks/goroutines-limit/solutions"
)

var cases = []struct {
	totalNumber int
	goCount     int
}{
	{totalNumber: 100, goCount: 5},
	{totalNumber: 200, goCount: 5},
	{totalNumber: 100, goCount: 10},
	{totalNumber: 100, goCount: 10},
	{totalNumber: 100, goCount: 100},
}

func BenchmarkSemaphoreUsing(b *testing.B) {
	for _, ch := range cases {
		b.Run(fmt.Sprintf("input_size_%d_goroutines_%d", ch.totalNumber, ch.goCount), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sl.SemaphoreUsing(ch.totalNumber, ch.goCount)
			}
		})
	}
}
