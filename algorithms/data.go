package algorithms

import (
	"math/rand"
	"sort"
)

var (
	is   IntSlice
	item int
)

//生成一个长度为10000整数数组，
//向这个数组随机地填充满整数，
//并对这个数组排序（因为二分搜索只适用于有序列）
func init()  {
	length := 10000
	maxNum := 1000000
	is = make(IntSlice, length)
	for i := 0; i < length; i++ {
		is[i] = rand.Intn(maxNum)
	}
	sort.Sort(is)
	item = is[rand.Intn(length)]
}

//定义一个可搜索接口
//实现了本接口的类型，都是可搜索的
type Searchable interface {
	Len() int
	LessThanItem(i int, item interface{}) bool
	Equal(i int, item interface{}) bool
}

//这个类型实现了可排序接口和可搜索接口
type IntSlice []int

func (is IntSlice) Len() int {
	return len(is)
}

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is IntSlice) LessThanItem(i int, item interface{}) bool {
	return is[i] < item.(int)
}

func (is IntSlice) Equal(i int, item interface{}) bool {
	return is[i] == item.(int)
}


