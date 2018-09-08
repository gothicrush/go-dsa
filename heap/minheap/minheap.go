package minheap

import (
	"../../assist"
	list "../../list/sqlist"
)

type Heap struct {
	data *list.List
	comparator assist.Comparator
}

func New(comparator assist.Comparator) *Heap {

	return &Heap{
		data: list.New(),
		comparator: comparator,
	}
}

func Heapify(li []interface{}, comparator assist.Comparator) *Heap {

	newList := list.New()

	for index, val := range li {
		newList.Add(index, val)
	}

	heap := &Heap{
		data: newList,
		comparator: comparator,
	}

	index := parents(heap.data.Size() - 1)
	for ; index >= 0; index-- {
		heap.shiftDown(index)

	}

	return heap
}

func (heap *Heap) Add(v interface{}) {

	heap.data.Add(heap.data.Size(), v)

	heap.shiftUp(heap.data.Size() - 1)
}

func (heap *Heap) ExtractMin() interface{} {

	var ret interface{} = heap.data.Get(0)

	heap.data.Set(0, heap.data.Get(heap.data.Size()-1))

	heap.data.Remove(heap.data.Size() - 1)

	heap.shiftDown(0)

	return ret
}

func (heap *Heap) PeekMin() interface{} {

	return heap.data.Get(0)
}

func (heap *Heap) Replace(v interface{}) interface{} {

	var ret interface{} = heap.data.Get(0)

	heap.data.Set(0, v)

	heap.shiftDown(0)

	return ret
}

func (heap *Heap) Size() int {
	return heap.data.Size()
}

func (heap *Heap) Empty() bool {
	return heap.data.Empty()
}

func (heap *Heap) shiftDown(k int) {

	for leftChild(k) < heap.data.Size() {

		var lor int = leftChild(k)

		if lor+1 < heap.data.Size() && assist.Compare(heap.data.Get(lor),
				                                      heap.data.Get(lor + 1),
				                                      heap.comparator) > 0 {
			lor += 1
		}

		if assist.Compare(heap.data.Get(k),
			              heap.data.Get(lor),
			              heap.comparator) <= 0 {
			break
		}

		temp := heap.data.Get(k)
		heap.data.Set(k, heap.data.Get(lor))
		heap.data.Set(lor, temp)

		k = lor
	}
}

func (heap *Heap) shiftUp(k int) {

	for k > 0 && assist.Compare(heap.data.Get(parents(k)),
			                    heap.data.Get(k),
			                    heap.comparator) > 0 {

			temp := heap.data.Get(parents(k))
			heap.data.Set(parents(k), heap.data.Get(k))
			heap.data.Set(k, temp)

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
