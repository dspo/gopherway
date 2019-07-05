package algorithms

import "testing"

var (
	is = []int{1, 2, 3, 4, 5, 6, 7 ,8, 9, 10}
	item = 8
)

func TestBinarySearchLoop(t *testing.T) {
	index, _ := BinarySearchLoop(is, item)
	if index != 7 {
		t.Error("search failed")
	}
}

