package algorithms

import (
	"sort"
)

//冒泡排序：
//总是遍历列表未排序部分
//但浪费大量交换操作
func BubbleSort(li sort.Interface) {
	for i := li.Len() - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if li.Less(j + 1, j) {
				li.Swap(j, j + 1)
			}
		}
	}
}

//短冒泡排序：
//可检查已排序列表以便提前结束
func ShortBubbleSort(li sort.Interface) {
	exchange := true
	passNum := li.Len() - 1
	for passNum > 0 && exchange {
		exchange = false
		for i := 0; i < passNum; i++ {
			if li.Less(i + 1, i) {
				exchange = true
				li.Swap(i, i + 1)
			}
		}
		passNum--
	}
}

//选择排序：
//总是找找到子表中的最大元素，放到表首
func SelectionSort(li sort.Interface) {
	for i := li.Len() - 1; i > 0; i-- {
		positionOfMax := 0
		for j := 1; j < i + 1; j++ {
			if li.Less(positionOfMax, j) {
				positionOfMax = j
			}
		}
		li.Swap(positionOfMax, i)
	}
}

func ChenSort(li sort.Interface) {
	i := 0
	for i < li.Len() - 1{
		for j := i + 1; j < li.Len(); j++ {
			if li.Less(j, i) {
				li.Swap(i, j)
				i--
				break
			}
		}
		i++
	}
}

func InsertionSort(li sort.Interface) {
	for i := 1; i < li.Len(); i++ {
		for j := i; j > 0 && li.Less(j, j - 1); j-- {
			li.Swap(j - 1, j)
		}
	}
}

func GoOfficialInsertionSrot(li sort.Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && li.Less(j, j-1); j-- {
			li.Swap(j, j-1)
		}
	}
}

func GoOfficialHeapSort(li sort.Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(li, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		li.Swap(first, first+i)
		siftDown(li, lo, i, first)
	}
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(data sort.Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}