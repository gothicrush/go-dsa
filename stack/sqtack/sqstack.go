package sqstack

import (
    list "../../list/sqlist"
)

type sqstack struct {
	li *list.List
}

func New() *sqstack {

	return &sqstack {
		li: list.New(),
	}
}

func (stack *sqstack) Push(v interface{}) {

	stack.li.Add(stack.li.Size(),v)
}

func (stack *sqstack) Pop() {

	stack.li.Remove(stack.li.Size()-1)
}

func (stack *sqstack) Top() interface{} {

	return stack.li.Get(stack.li.Size()-1)
}

func (stack *sqstack) Size() int {

	return stack.li.Size()
}

func (stack *sqstack) Empty() bool {

	return stack.li.Size() == 0
}

func (stack *sqstack) String() string {

	str := stack.li.String()
	str = str[:len(str)-1] + "> stack top"

	return str
}