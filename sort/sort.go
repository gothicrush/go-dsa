package sort

import (
    "../assist"
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