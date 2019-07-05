package algorithms

import "testing"

func TestLinearSearch(t *testing.T) {
	data := []int{1,2,8,9,3,4,6,5,3,8,9,6,8}
	key := 8
	index, in := LinearSearch(data, key)
	t.Log(index, in)
}