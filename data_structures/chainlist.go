package data_structures

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	current *Node
}

//make a last-to-head loop chain list
func NewLoopLinkedList(data ...interface{}) *LinkedList {
	dataSlice := append([]interface{}{}, data...)
	nodeSlice := make([]*Node, len(dataSlice))
	for i := len(dataSlice) - 1; i >= 0; i--{
		nodeSlice[i] = &Node{data:dataSlice[i]}
		if i + 1 < len(dataSlice) {
			nodeSlice[i].next = nodeSlice[i + 1]
		}
	}
	nodeSlice[len(dataSlice) - 1].next = nodeSlice[0]
	var current Node
	current = *nodeSlice[0]
	return &LinkedList{nodeSlice[0], &current}
}

//make an un-loop chain list
func NewLinkedList(data ...interface{}) *LinkedList {
	dataSlice := append([]interface{}{}, data...)
	nodeSlice := make([]*Node, len(dataSlice))
	for i := len(dataSlice) - 1; i >= 0; i--{
		nodeSlice[i] = &Node{data:dataSlice[i]}
		if i + 1 < len(dataSlice) {
			nodeSlice[i].next = nodeSlice[i + 1]
		}
	}
	return &LinkedList{nodeSlice[0], nodeSlice[0]}
}

//make list.current a node copy that point to the next node of current node
func (list LinkedList) Next() bool {
	if list.current == nil {
		return false
	}
	*list.current = *list.current.next
	return true
}

//make data a node, and insert it into the list before the first node,
//and let it be head.
//this method is only for un-loop chain list.
func (list *LinkedList) InsertFirst(data interface{}) {
	if list.head == nil {
		list.head = &Node{data:data, next:nil}
		list.current = &Node{data:data, next:nil}
		return
	}
	node := Node{data:data, next:list.head}
	*list.head = node
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