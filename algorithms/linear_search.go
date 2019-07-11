package algorithms

//线性查找从头找到尾，不要求列表是有序的
//仅适用于[]int
func LinearSearch(data []int, key int) (int, bool) {
	for index, ele := range data{
		if ele == key {
			return index, true
		}
	}
	return -1, false
}

//适用于任何实现了Searchable接口的类型
func LinearSearchForSearchable(data Searchable, item interface{}) (int, bool) {
	for i := 0; i < data.Len(); i++ {
		if data.Equal(i, item) {
			return i, true
		}
	}
	return -1, false
}