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

type List struct {
	dummyNode *listNode
	size int
}

func New() *List {
	return &List {
		dummyNode: newNode(0, nil),
		size: 0,
	}
}

func (li *List) Add(index int, v interface{}) {
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

func (li *List) Remove(index int) interface{} {
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

func (li *List) Set(index int, v interface{}) {
	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	cur := li.dummyNode.next

	for i := 0; i < index; i++ {
		cur = cur.next
	}

	cur.data = v
}

func (li *List) Get(index int) interface{} {
	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	cur := li.dummyNode.next

	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.data
}

func (li *List) Contains(v interface{}) bool {
	cur := li.dummyNode.next

	for cur != nil {
		if cur.data == v {
			return true
		}
		cur = cur.next
	}

	return false
}

func (li *List) IndexOf(v interface{}) int {
	cur := li.dummyNode.next

	for i := 0; cur != nil; i++ {
		if cur.data == v {
			return i
		}
		cur = cur.next
	}

	return -1
}

func (li *List) Empty() bool {
	return li.size == 0
}

func (li *List) Size() int {
	return li.size
}

func (li *List) String() string {
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