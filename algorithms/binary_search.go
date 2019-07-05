package algorithms


//二分查找要求列表是有序的
func BinarySearchLoop(sortedArray []int, item int) (int, bool) {
	init, end := 0, len(sortedArray) - 1
	for init <= end{
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

