package algorithms

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList(data ...interface{}) *LinkedList {
	dataSlice := append([]interface{}{}, data...)
	nodeSlice := make([]*Node, len(dataSlice))
	for i := len(dataSlice) - 1; i >= 0; i--{
		nodeSlice[i] = &Node{data:dataSlice[i]}
		if i + 1 < len(dataSlice) {
			nodeSlice[i].next = nodeSlice[i + 1]
		}
	}
	nodeSlice[len(dataSlice) - 1].next = nodeSlice[0]
	return &LinkedList{nodeSlice[0]}
}

func (list *LinkedList) InsertFirst(i interface{}) {
	data := &Node{data: i}
	if list.head != nil {
		data.next = list.head
	}
	list.head = data
}

func (list *LinkedList) InsertLast(i interface{}) {
	data := &Node{data: i}
	if list.head == nil {
		list.head = data
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = data
}

func (list *LinkedList) RemoveByValue(i interface{}) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == i {
		list.head = list.head.next
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == i {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head = list.head.next
		return true
	}
	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	current.next = current.next.next
	return true
}

func (list *LinkedList) SearchValue(i interface{}) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if current.data == i {
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) GetFirst() (interface{}, bool) {
	if list.head == nil {
		return 0, false
	}
	return list.head.data, true
}

func (list *LinkedList) GetLast() (interface{}, bool) {
	if list.head == nil {
		return 0, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current.data, true
}

func (list *LinkedList) GetSize() int {
	count := 0
	current := list.head
	for current != nil {
		count += 1
		current = current.next
	}
	return count
}

func (list *LinkedList) GetItems() []interface{} {
	var items []interface{}
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}