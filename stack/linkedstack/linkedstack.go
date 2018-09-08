package linkedstack

import (
    list "../../list/singlelist"
)

type Stack struct {
	li *list.List
}

func New() *Stack {

	return &Stack {
		li: list.New(),
	}
}

func (stack *Stack) Push(v interface{}) {

	stack.li.Add(stack.li.Size(),v)
}

func (stack *Stack) Pop() {

	stack.li.Remove(stack.li.Size()-1)
}

func (stack *Stack) Top() interface{} {

	return stack.li.Get(stack.li.Size()-1)
}

func (stack *Stack) Size() int {

	return stack.li.Size()
}

func (stack *Stack) Empty() bool {

	return stack.li.Size() == 0
}

func (stack *Stack) String() string {

	str := stack.li.String()
	str = str[:len(str)-1] + "> stack top"

	return str
}