package solutions_test

import (
	"testing"

	"github.com/DePavelPo/multithreading-practice/tasks/chan-producer-consumer/solutions"
)

func BenchmarkBufferedChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solutions.BufferedChannel()
	}
}
