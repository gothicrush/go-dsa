package minheap

import (
    "testing"
    "time"
    "math/rand"
    "fmt"
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

func TestExtractMin(t *testing.T) {

	heap := New(compare)

	heap.Add(2)
	heap.Add(5)
	heap.Add(0)
	heap.Add(3)
	heap.Add(1)

	min := heap.ExtractMin().(int)

	if min != 0 {
		t.Errorf("TestExtractMin Error %v", min)
	}
}

func TestPeekMin(t *testing.T) {

	heap := New(compare)

	heap.Add(2)
	heap.Add(5)
	heap.Add(0)
	heap.Add(3)
	heap.Add(1)
	heap.Add(4)

	min := heap.ExtractMin().(int)

	min = heap.PeekMin().(int)

	if min != 1 {
		t.Errorf("TestPeekMin Error %v", min)
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

	min := heap.Replace(-666).(int)

	if min != 0 {
		t.Errorf("TestReplace Error %v", min)
	}

	min = heap.PeekMin().(int)

	if min != -666 {
		t.Errorf("TestReplace Error %v", min)
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

	for i := 0; i < 1000; i++ {
		fmt.Printf("%v ", arr[i])

		if i % 20 == 0 {
			fmt.Println()
		}
	}


	var result []interface{}

	heap := Heapify(arr, compare)

	for !heap.Empty() {
		result = append(result, heap.ExtractMin())
	}


	fmt.Printf("\n\n\n\n\n")
	for i := 0; i < 1000; i++ {
		fmt.Printf("%v ", result[i])

		if i % 20 == 0 {
			fmt.Println()
		}
	}


	for i := 1; i < len(result); i++ {
		
		num1 := result[i-1].(int)
		num2 := result[i].(int)

		if num2 < num1 {
			t.Errorf("TestHeapify Error %v < %v", num2, num1)
		}
	}
}