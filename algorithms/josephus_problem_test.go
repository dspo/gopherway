package algorithms

import (
	"testing"
)

func TestNewJCircle(t *testing.T) {
	list := NewLinkedList("a", "b", "c", "d", "e")
	curNode := list.head
	for i := 0; i < 20; i++{
		t.Log(curNode.data)
		curNode = curNode.next
	}
}
