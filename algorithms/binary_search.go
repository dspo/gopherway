package algorithms

//二分查找要求列表是有序的
//这个二分搜索函数仅适用于[]int
func BinarySearchForIntSlice(sortedArray []int, item int) (int, bool) {
	init, end := 0, len(sortedArray)-1
	for init <= end {
		middle := (end + init) >> 1
		if sortedArray[middle] == item {
			return middle, true
		}
		if sortedArray[middle] < item {
			init = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1, false
}

//这个二分搜索函数适用于实现了Searchable接口的类型
//见binary_search_test.go
func BinarySearchForSearchable(sorted Searchable, item interface{}) (int, bool) {
	init, end := 0, sorted.Len()-1
	for init <= end {
		middle := (end + init) >> 1
		if sorted.Equal(middle, item) {
			return middle, true
		}
		if sorted.LessThanItem(middle, item) {
			init = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1, false
}
