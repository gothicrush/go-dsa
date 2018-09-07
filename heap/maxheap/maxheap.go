package maxheap

import (
	"../../assist"
	list "../../list/sqlist"
)

type maxheap struct {
	data *list.List
	comparator assist.Comparator
}

func New(comparator assist.Comparator) *maxheap {

	return &maxheap{
		data: list.New(),
		comparator: comparator,
	}
}

func Heapify(li *[]interface{}, comparator assist.Comparator) *maxheap {

	newList := list.New()

	for index, val := range *li {
		newList.Add(index, val)
	}

	heap := &maxheap{
		data: newList,
		comparator: comparator,
	}

	index := parents(heap.data.Size() - 1)
	for ; index >= 0; index-- {
		heap.shiftDown(index)

	}

	return heap
}

func (heap *maxheap) Add(v interface{}) {

	heap.data.Add(heap.data.Size(), v)

	heap.shiftUp(heap.data.Size() - 1)
}

func (heap *maxheap) ExtractMax() interface{} {

	var ret interface{} = heap.data.Get(0)

	heap.data.Set(0, heap.data.Get(heap.data.Size()-1))

	heap.data.Remove(heap.data.Size() - 1)

	heap.shiftDown(0)

	return ret
}

func (heap *maxheap) PeekMax() interface{} {

	return heap.data.Get(0)
}

func (heap *maxheap) Replace(v interface{}) interface{} {

	var ret interface{} = heap.data.Get(0)

	heap.data.Set(0, v)

	heap.shiftDown(0)

	return ret
}

func (heap *maxheap) Size() int {
	return heap.data.Size()
}

func (heap *maxheap) Empty() bool {
	return heap.data.Empty()
}

func (heap *maxheap) shiftDown(k int) {

	for leftChild(k) < heap.data.Size() {

		var lor int = leftChild(k)

		if lor+1 < heap.data.Size() {

			leftVal := heap.data.Get(lor)
			rightVal := heap.data.Get(lor + 1)

			if assist.Compare(leftVal,rightVal,heap.comparator) < 0 {
				lor += 1
			}
		}

		var kVal interface{} = heap.data.Get(k)
		var lorVal interface{} = heap.data.Get(lor)

		if assist.Compare(kVal,lorVal,heap.comparator) >= 0 {
			break
		}

		heap.data.Set(k, lorVal)
		heap.data.Set(lor, kVal)

		k = lor
	}
}

func (heap *maxheap) shiftUp(k int) {

	for k > 0 {

		var pVal interface{} = heap.data.Get(parents(k))
		var kVal interface{} = heap.data.Get(k)

		if assist.Compare(pVal,kVal,heap.comparator) < 0 {
			heap.data.Set(parents(k), &kVal)
			heap.data.Set(k, &pVal)
		}

		k = parents(k)
	}
}

func parents(k int) int {
	return (k - 1) / 2
}

func leftChild(k int) int {
	return k*2 + 1
}

func rightChild(k int) int {
	return k*2 + 2
}
