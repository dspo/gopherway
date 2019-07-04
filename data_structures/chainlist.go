package data_structures

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	cursor *Node
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
	cursor := *nodeSlice[0]
	return &LinkedList{nodeSlice[0], &cursor}
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

//make list.cursor a node copy that point to the next node of cursor node
func (list LinkedList) Next() bool {
	if list.cursor == nil {
		return false
	}
	*list.cursor = *list.cursor.next
	return true
}