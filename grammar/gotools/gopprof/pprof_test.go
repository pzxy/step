package gopprof

import "testing"

func BenchmarkSum(b *testing.B) {
	aa := 1
	bb := 2
	for i := 0; i < b.N; i++ {
		Sum(aa, bb)
	}
}
