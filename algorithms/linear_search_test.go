package algorithms

import "testing"

func TestLinearSearch(t *testing.T) {
	index, in := LinearSearch(is, item)
	t.Log(index, in)
}

func TestLinearSearchForSearchable(t *testing.T) {
	index, in := LinearSearchForSearchable(is, item)
	t.Log(index, in)
}

func BenchmarkLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LinearSearch(is, item)
	}
}

func BenchmarkLinearSearchForSearchable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LinearSearchForSearchable(is, item)
	}
}