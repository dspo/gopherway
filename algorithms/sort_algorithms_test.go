package algorithms

import (
	"sort"
	"testing"
)

var (
	li       = intLi{3, 0, 9, 8, 6, 7, 3, 8, 6, -1, 0, 6}
	sortedLi = make(intLi, li.Len())
)

func init() {
	copy(sortedLi, li)
	sort.Sort(sortedLi)
}

type intLi []int

func (li intLi) Len() int {
	return len(li)
}

func (li intLi) Less(i, j int) bool {
	return li[i] < li[j]
}
func (li intLi) Swap(i, j int) {
	li[i], li[j] = li[j], li[i]
}

func (li intLi) Equal(other intLi) bool {
	if li.Len() != other.Len() {
		return false
	}
	for i := 0; i < li.Len(); i++ {
		if li[i] != other[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	forSorting := li
	BubbleSort(forSorting)
	t.Log(forSorting)
	t.Log(sortedLi)
	if !forSorting.Equal(sortedLi) {
		t.Error("srot failed !")
	}
}

func TestShortBubbleSort(t *testing.T) {
	forSorting := li
	ShortBubbleSort(forSorting)
	t.Log(forSorting)
	t.Log(sortedLi)
	if !forSorting.Equal(sortedLi) {
		t.Error("sort failed !")
	}
}

func TestSelectionSort(t *testing.T) {
	forSorting := li
	SelectionSort(forSorting)
	t.Log(forSorting)
	t.Log(sortedLi)
	if !forSorting.Equal(sortedLi) {
		t.Error("sort failed !")
	}
}

func TestInsertionSort(t *testing.T) {
	forSorting := li
	InsertionSort(forSorting)
	t.Log(forSorting)
	t.Log(sortedLi)
	if !forSorting.Equal(sortedLi) {
		t.Error("sort faild !")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	forSorting := li
	for i := 0; i < b.N; i++ {
		BubbleSort(forSorting)
	}
}

func BenchmarkShortBubbleSort(b *testing.B) {
	forSorting := li
	for i := 0; i < b.N; i++ {
		ShortBubbleSort(forSorting)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	forSorting := li
	for i := 0; i < b.N; i++ {
		SelectionSort(forSorting)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	forSorting := li
	for i := 0; i < b.N; i++ {
		InsertionSort(forSorting)
	}
}
