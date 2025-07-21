package solutions_test

import (
	"fmt"
	"testing"

	"github.com/DePavelPo/multithreading-practice/tasks/chan-data-filter/solutions"
)

var cases = []struct {
	totalNumber int
}{
	{totalNumber: 100},
	{totalNumber: 200},
	{totalNumber: 500},
	{totalNumber: 1000},
}

func BenchmarkDataFilter(b *testing.B) {
	for _, cs := range cases {
		b.Run(fmt.Sprintf("input_size_%d", cs.totalNumber), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				solutions.DataFilter(cs.totalNumber)
			}
		})
	}
}
