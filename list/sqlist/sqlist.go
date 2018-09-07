package sqlist

import (
	"fmt"
)

type List struct {
	data []interface{}
	size int
}

func New() *List {
	return &List{
		data: make([]interface{}, 10),
		size: 0,
	}
}

func (li *List) Add(index int, v interface{}) {

	if index < 0 || index > li.size {
		panic("index out of bounds")
	}

	if li.size == len(li.data) {
		resize(li, 2*len(li.data))
	}

	for i := li.size - 1; i >= index; i-- {
		li.data[i+1] = li.data[i]
	}

	li.data[index] = v

	li.size++
}

func (li *List) Remove(index int) interface{} {

	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	ret := li.data[index]

	for i := index + 1; i < li.size; i++ {
		li.data[i-1] = li.data[i]
	}

	li.size--

	if li.size <= len(li.data)/2 {
		resize(li, len(li.data)/2)
	}

	return ret
}

func (li *List) Set(index int, v interface{}) {
	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	li.data[index] = v
}

func (li *List) Get(index int) interface{} {
	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	return li.data[index]
}

func (li *List) Contains(v interface{}) bool {

	for _, item := range li.data {
		if item == v {
			return true
		}
	}

	return false
}

func (li *List) IndexOf(v interface{}) int {

	for index, item := range li.data {
		if item == v {
			return index
		}
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

	ret := "[ "

	for i := 0; i < li.size; i++ {

		if i == li.size-1 {
			break
		}

		ret = ret + fmt.Sprintf("%v,", li.data[i])
	}

	if li.size > 0 {
		ret = ret + fmt.Sprintf("%v ]", li.data[li.size-1])
	} else {
		ret = ret + " ]"
	}

	return ret
}

func resize(li *List, newSize int) {

	newList := make([]interface{}, newSize)

	copy(newList, li.data)

	li.data = newList
}
