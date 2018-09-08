package priorityqueue

import (
	"../../assist"
    "../../heap/maxheap"
)

type Queue struct {
	heap *maxheap.Heap
}

func New(comparator assist.Comparator) *Queue {

	return &Queue {
		heap: maxheap.New(comparator),
	}
}

func (queue *Queue) EnQueue(v interface{}) {

	queue.heap.Add(v)
}

func (queue *Queue) DeQueue() {

	queue.heap.ExtractMax()
}

func (queue *Queue) Head() interface{} {

	return queue.heap.PeekMax()
}

func (queue *Queue) Size() int {

	return queue.heap.Size()
}

func (queue *Queue) Empty() bool {

	return queue.heap.Empty()
}