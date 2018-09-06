package linkedqueue

import (
    list "../../list/singlelist"
)

type linkedqueue struct {
	li *list.List
}

func New() *linkedqueue {

	return &linkedqueue {
		li: list.New(),
	}
}

func (queue *linkedqueue) EnQueue(v interface{}) {

	queue.li.Add(queue.li.Size(),v)
}

func (queue *linkedqueue) DeQueue() {

	queue.li.Remove(0)
}

func (queue *linkedqueue) Head() interface{} {

	return queue.li.Get(0)
}

func (queue *linkedqueue) Size() int {

	return queue.li.Size()
}

func (queue *linkedqueue) Empty() bool {

	return queue.li.Empty()
}

func (queue *linkedqueue) String() string {

	str := queue.li.String()

	str = ">>> " + str[1:len(str)-1] + " >>> queue head"

	return str
}