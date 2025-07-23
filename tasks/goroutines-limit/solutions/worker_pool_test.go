package solutions_test

import (
	"fmt"
	"testing"

	sl "github.com/DePavelPo/multithreading-practice/tasks/goroutines-limit/solutions"
)

func BenchmarkWorkerPoolUsing(b *testing.B) {
	for _, ch := range cases {
		b.Run(fmt.Sprintf("input_size_%d_goroutines_%d", ch.totalNumber, ch.goCount), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sl.WorkerPoolUsing(ch.totalNumber, ch.goCount)
			}
		})
	}
}
