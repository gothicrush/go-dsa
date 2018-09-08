package maxheap

import (
    "testing"
    "time"
    "math/rand"
)

func compare(i interface{}, j interface{}) int {
	var ii int = i.(int)
	var jj int = j.(int)

	return ii - jj
}

func TestInit(t *testing.T) {

	heap := New(compare)

	if heap == nil {
		t.Errorf("TestInit Error")
	}
}

func TestAdd(t *testing.T) {

	heap := New(compare)

	heap.Add(2)
	heap.Add(5)
	heap.Add(0)
	heap.Add(3)
	heap.Add(1)

	if heap == nil && heap.Size() != 5 {
		t.Errorf("TestAdd Error,%v, %v", heap, heap.Size())
	}
}

func TestExtractMax(t *testing.T) {

	heap := New(compare)

	heap.Add(2)
	heap.Add(5)
	heap.Add(0)
	heap.Add(3)
	heap.Add(1)

	max := heap.ExtractMax().(int)

	if max != 5 {
		t.Errorf("TestExtractMax Error %v", max)
	}
}

func TestPeekMax(t *testing.T) {

	heap := New(compare)

	heap.Add(2)
	heap.Add(5)
	heap.Add(0)
	heap.Add(3)
	heap.Add(1)
	heap.Add(4)

	max := heap.ExtractMax().(int)

	max = heap.PeekMax().(int)

	if max != 4 {
		t.Errorf("TestPeekMax Error %v", max)
	}
}

func TestReplace(t *testing.T) {

	heap := New(compare)

	heap.Add(2)
	heap.Add(5)
	heap.Add(0)
	heap.Add(3)
	heap.Add(1)
	heap.Add(4)

	max := heap.Replace(666).(int)

	if max != 5 {
		t.Errorf("TestReplace Error %v", max)
	}

	max = heap.PeekMax().(int)

	if max != 666 {
		t.Errorf("TestReplace Error %v", max)
	}
}

func TestEmpty(t *testing.T) {

	heap := New(compare)

	if !heap.Empty() {
		t.Errorf("TestEmpty Error empty: %v", heap.Empty())
	}

	heap.Add(2)
	heap.Add(5)
	heap.Add(0)
	heap.Add(3)
	heap.Add(1)
	heap.Add(4)

	if heap.Empty() {
		t.Errorf("TestEmpty Error empty: %v", heap.Empty())
	}
}

func TestSize(t *testing.T) {

	heap := New(compare)

	if heap.Size() != 0 {
		t.Errorf("TestSize Error size: %v", heap.Size())
	}

	heap.Add(2)
	heap.Add(5)
	heap.Add(0)
	heap.Add(3)
	heap.Add(1)
	heap.Add(4)

	if heap.Size() != 6 {
		t.Errorf("TestSize Error size: %v", heap.Size())
	}
}

func TestHeapify(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var nums []int

	for i := 0; i < 300000; i++ {
		nums = append(nums,r.Intn(300000))
	}



	var arr []interface{}

	for _, val := range nums {
		arr = append(arr, val)
	}


	var result []interface{}

	heap := Heapify(arr, compare)

	for !heap.Empty() {
		result = append(result, heap.ExtractMax())
	}


	for i := 1; i < len(result); i++ {
		
		num1 := result[i-1].(int)
		num2 := result[i].(int)

		if num2 > num1 {
			t.Errorf("TestHeapify Error %v > %v", num2, num1)
		}
	}
}