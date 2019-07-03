package data_structures

import (
	"testing"
)

func TestNewLoopLinkedList(t *testing.T) {
	data := []string{"a", "b", "c", "d", "e"}
	list := NewLoopLinkedList("a", "b", "c", "d", "e")
	curNode := list.head
	for i := 0; i < 20; i++{
		t.Log(curNode.data)
		if curNode.data != data[i % 5] {
			t.Error("error")
		}
		curNode = curNode.next
	}
}

func TestLinkedList_Next(t *testing.T) {
	list := NewLoopLinkedList("a", "b", "c", "d", "e")
	if list.current.data != "a" {
		t.Error("current node error")
	}
	list.Next()
	if list.current.data != "b" {
		t.Error("Next() not move the pointer to next")
	}
	if list.head.data != "a" {
		t.Error("the head moved but it should not move")
	}
	list.Next(); list.Next(); list.Next(); list.Next()
	if list.current.data != "a" {
		t.Error("Next() not move the pointer to loop list head when the last point at the last")
	}
}

func TestLinkedList_InsertFirst(t *testing.T) {
	list := &LinkedList{head:nil, current:nil}
	list.InsertFirst("b")
	t.Log(list.head.data, list.current.data)

	list.InsertFirst("a")
	t.Log(list.head.data, list.current.data)
}
