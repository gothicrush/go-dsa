package sqlist

import (
	"fmt"
)

type sqlist struct {
	data []interface{}
	size int
}

func NewList() *sqlist {
	return &sqlist{
		data: make([]interface{}, 10),
		size: 0,
	}
}

func (li *sqlist) Add(index int, v interface{}) {

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

func (li *sqlist) Remove(index int) interface{} {

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

func (li *sqlist) Set(index int, v interface{}) {
	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	li.data[index] = v
}

func (li *sqlist) Get(index int) interface{} {
	if index < 0 || index >= li.size {
		panic("index out of bounds")
	}

	return li.data[index]
}

func (li *sqlist) Contains(v interface{}) bool {

	for _, item := range li.data {
		if item == v {
			return true
		}
	}

	return false
}

func (li *sqlist) IndexOf(v interface{}) int {

	for index, item := range li.data {
		if item == v {
			return index
		}
	}

	return -1
}

func (li *sqlist) Empty() bool {
	return li.size == 0
}

func (li *sqlist) Size() int {
	return li.size
}

func (li *sqlist) String() string {

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

func resize(li *sqlist, newSize int) {

	newList := make([]interface{}, newSize)

	copy(newList, li.data)

	li.data = newList
}
