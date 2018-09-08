package linkedqueue

import (
    list "../../list/singlelist"
)

type Queue struct {
	li *list.List
}

func New() *Queue {

	return &Queue {
		li: list.New(),
	}
}

func (queue *Queue) EnQueue(v interface{}) {

	queue.li.Add(queue.li.Size(),v)
}

func (queue *Queue) DeQueue() {

	queue.li.Remove(0)
}

func (queue *Queue) Head() interface{} {

	return queue.li.Get(0)
}

func (queue *Queue) Size() int {

	return queue.li.Size()
}

func (queue *Queue) Empty() bool {

	return queue.li.Empty()
}

func (queue *Queue) String() string {

	str := queue.li.String()

	str = ">>> " + str[1:len(str)-1] + " >>> queue head"

	return str
}