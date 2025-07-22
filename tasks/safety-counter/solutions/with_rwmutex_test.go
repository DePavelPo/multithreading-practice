package solutions_test

import (
	"fmt"
	"testing"

	sl "github.com/DePavelPo/multithreading-practice/tasks/safety-counter/solutions"
)

var cases = []struct {
	totalNumber int
}{
	{totalNumber: 100},
	{totalNumber: 200},
	{totalNumber: 500},
	{totalNumber: 1000},
}

func BenchmarkWithRWMutex(b *testing.B) {
	for _, cs := range cases {
		b.Run(fmt.Sprintf("input_size_%d", cs.totalNumber), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sl.WithRWMutex()
			}
		})
	}
}
