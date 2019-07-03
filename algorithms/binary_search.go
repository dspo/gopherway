package algorithms

func BinarySearchLoop(sortedArray []int, item int) int {
	init, end := 0, len(sortedArray) - 1
	for init <= end{
		middle := (end + init) >> 1
		if sortedArray[middle] == item {
			return middle
		}
		if sortedArray[middle] < item {
			init = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}

