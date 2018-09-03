package singlelist

import (
    "fmt"
)

type listNode struct {
	data interface{}
	next *listNode
}

func newNode(v interface{}, n *listNode) *listNode {
	return &listNode {
		data: v,
		next: n,
	}
}

type list struct {
	dummyNode *listNode
	size int
}

func NewList() *list {
	return &list {
		dummyNode: newNode(0, nil),
		size: 0,
	}
}

func (li *list) Add(index int, v interface{}) {
	if index < 0 || index > li.size {
		panic("index out of bounds")
	}

	prev := li.dummyNode

	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = newNode(v, prev.next)

	li.size ++
}

func (li *list) Remove(index int) interface{} {
	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	prev := li.dummyNode

	for i := 0; i < index; i++ {
		prev = prev.next
	}

	val := prev.next.data

	prev.next = prev.next.next

	li.size --

	return val
}

func (li *list) Set(index int, v interface{}) {
	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	cur := li.dummyNode.next

	for i := 0; i < index; i++ {
		cur = cur.next
	}

	cur.data = v
}

func (li *list) Get(index int) interface{} {
	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	cur := li.dummyNode.next

	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.data
}

func (li *list) Contains(v interface{}) bool {
	cur := li.dummyNode.next

	for cur != nil {
		if cur.data == v {
			return true
		}
		cur = cur.next
	}

	return false
}

func (li *list) IndexOf(v interface{}) int {
	cur := li.dummyNode.next

	for i := 0; cur != nil; i++ {
		if cur.data == v {
			return i
		}
		cur = cur.next
	}

	return -1
}

func (li *list) Empty() bool {
	return li.size == 0
}

func (li *list) Size() int {
	return li.size
}

func (li *list) String() string {
	cur := li.dummyNode.next

	var str string = "[ "

	for cur != nil {
		if cur.next == nil {
			break
		}

		str = str + fmt.Sprintf("%v -> ",cur.data)
		cur = cur.next
	}

	if cur != nil {
		str = str + fmt.Sprintf("%v",cur.data) + " ]"
	} else {
		str = str + "]"
	}

	return str
}