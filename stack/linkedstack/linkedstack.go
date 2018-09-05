package linkedstack

import (
    list "../../list/singlelist"
)

type linkedstack struct {
	li *list.List
}

func New() *linkedstack {

	return &linkedstack {
		li: list.New(),
	}
}

func (stack *linkedstack) Push(v interface{}) {

	stack.li.Add(stack.li.Size(),v)
}

func (stack *linkedstack) Pop() {

	stack.li.Remove(stack.li.Size()-1)
}

func (stack *linkedstack) Top() interface{} {

	return stack.li.Get(stack.li.Size()-1)
}

func (stack *linkedstack) Size() int {

	return stack.li.Size()
}

func (stack *linkedstack) Empty() bool {

	return stack.li.Size() == 0
}

func (stack *linkedstack) String() string {

	str := stack.li.String()
	str = str[:len(str)-1] + "> stack top"

	return str
}