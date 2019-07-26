package algorithms

import (
	"testing"
)

func TestBinarySearchForSearchable(t *testing.T) {
	index, ok := BinarySearchForSearchable(is, item)
	if !ok {
		t.Error("not found index, something error")
	}
	index2, ok := BinarySearchForIntSlice(is, item)
	if !ok {
		t.Error("not found index, something error")
	}
	if index != index2 {
		t.Error("index found by BinarySearchForSearchable != index found by BinarySearchForIntSlice")
	}
	t.Log(index, "==", index2)
}

func BenchmarkBinarySearchForIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearchForIntSlice(is, item)
	}
}

//使用自定义接口要比操作[]int慢得多
func BenchmarkBinarySearchForSearchable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearchForSearchable(is, item)
	}
}
