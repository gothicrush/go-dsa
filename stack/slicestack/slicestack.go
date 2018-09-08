package slicestack

import (
    "fmt"
)

type Stack struct {
	li []interface{}
	size int
}

func New() *Stack {

	return &Stack {
		li: make([]interface{},10),
		size: 0,
	}
}

func (stack *Stack) Push(v interface{}) {

	if stack.size == len(stack.li) {
		var newSlice []interface{} = make([]interface{}, 2 * len(stack.li))
		copy(newSlice, stack.li)
		stack.li = newSlice
	}

	stack.li[stack.size] = v
	stack.size++	
}

func (stack *Stack) Pop() {

	stack.size--

	if stack.size < len(stack.li) / 2 {
		var newSlice []interface{} = make([]interface{}, len(stack.li) / 2)
		copy(newSlice, stack.li)
		stack.li = newSlice
	}
}

func (stack *Stack) Top() interface{} {
	return stack.li[stack.size-1]
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Empty() bool {
	return stack.size == 0
}

func (stack *Stack) String() string {

	str := "[ "

	for i := 0; i < stack.size-1; i++ {
		str = str + fmt.Sprintf("%v -> ", stack.li[i])
	}

	if stack.size > 0 {
		str = str + fmt.Sprintf("%v > stack top", stack.li[stack.size-1])
	} else {
		str = str + "> stack top"
	}

	return str
}