package sort

import (
    "../../assist"
    "../../heap/minheap"
)

type Sort struct {
    comparator assist.Comparator
}

func New(comparator assist.Comparator) *Sort {
	return &Sort {
		comparator: comparator,
	}
}

func swap(list *[]interface{}, i int, j int) {
	var temp interface{} = (*list)[i]
	(*list)[i] = (*list)[j]
	(*list)[j] = temp
}

func (sort *Sort) BubbleSort(list *[]interface{}) {

	for i := 0; i < len(*list) - 1; i++ {

		for j := i+1; j < len(*list); j++ {

			cmp := assist.Compare( (*list)[i], (*list)[j], sort.comparator  )

			if cmp > 0 {
				swap(list, i, j)
			}
		}
	}
}

func (sort *Sort) SelectionSort(list *[]interface{}) {

	for i := 0; i < len(*list) - 1; i++ {

		minIndex := i

		for j := i+1; j < len(*list); j++ {

			cmp := assist.Compare( (*list)[minIndex], (*list)[j], sort.comparator  )

			if cmp >= 0 {
				minIndex = j
			}
		}

		if i != minIndex {
			swap(list, i, minIndex)
		}
	}
}

func (sort *Sort) InsertionSort(list *[]interface{}) {

	for i := 1; i < len(*list); i++ {

		j := i

		for j > 0 && assist.Compare( (*list)[j], (*list)[j-1], sort.comparator ) < 0 {

			(*list)[j] = (*list)[j-1]

			j--
		}
	}
}

func (sort *Sort) ShellSort(list *[]interface{}) {

	N := len(*list)
	h := 1
	for h < N / 3 {
		h = h * 3 + 1
	}

	for h >= 1 {

		for i := h; i < len(*list); i++ {

			j := i

			for j > h-1 && assist.Compare( (*list)[j], (*list)[j-h], sort.comparator ) < 0 {

				(*list)[j] = (*list)[j-h]

				j = j - h
			}
		}

		h = h / 3
	}
}

func (sort *Sort)  HeapSort(list *[]interface{}) {
    
    heap := minheap.Heapify(*list, sort.comparator)

    var newList []interface{}

    for !heap.Empty() {
    	newList = append(newList, heap.ExtractMin())
    }

    *list = newList
}

func (sort *Sort) MergeSort(list *[]interface{}) {

    sort.mergeSort(list, 0, len(*list) - 1)
}

func (sort *Sort) mergeSort(list *[]interface{},  left int, right int) {

    if left >= right {
    	return
    }

    mid := ((right - left) / 2 ) + left

    sort.mergeSort(list, left, mid)
    sort.mergeSort(list, mid+1, right)

    if assist.Compare((*list)[mid], (*list)[mid+1], sort.comparator) > 0 {
        sort.merge(list, left, mid, right)
    }
}

func (sort *Sort) merge(list *[]interface{}, left int, mid int, right int) {
    
    var aux []interface{} = make([]interface{}, right - left + 1)

    for i := left; i <= right; i++ {
    	aux[i-left] = (*list)[i]
    }

    var i, j = left, mid + 1

    for k := left; k <= right; k++ {

    	if i > mid {
    		(*list)[k] = aux[j-left]
    		j++
    	} else if j > right {
    		(*list)[k] = aux[i-left]
    		i++
    	} else if assist.Compare(aux[i-left], aux[j-left], sort.comparator) < 0 {
    		(*list)[k] = aux[i-left]
    		i++
    	} else {
    		(*list)[k] = aux[j - left]
    		j++
    	}
    }
}

func (sort *Sort) QuickSort1(list *[]interface{}) {

	sort.quickSort1(list, 0, len(*list)-1)
}

func (sort *Sort) quickSort1(list *[]interface{}, left int, right int) {

	if left >= right {
		return
	}

	var p int = sort.partition1(list, left, right)
	sort.quickSort1(list,left,p-1)
	sort.quickSort1(list,p+1,right)
}

func (sort *Sort) partition1(list *[]interface{}, left int, right int) int {

	var v interface{} = (*list)[left]

	var j int = left

	for i := left + 1; i <= right; i++ {
		if assist.Compare((*list)[i], v, sort.comparator) < 0 {
			j++

			swap(list, i, j)
		}
	}

	swap(list, left, j)

	return j
}

func (sort *Sort) quickSort2(list *[]interface{}, left int, right int) int {}