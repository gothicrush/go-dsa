package sort

import (
	"../../assist"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func f(i interface{}, j interface{}) int {

	return i.(int) - j.(int)
}

func TestInit(t *testing.T) {

	s := New(f)

	if s == nil {
		t.Errorf("TestInit Error: %v", s)
	}
}

func TestBubbleSort(t *testing.T) {

	var data []interface{} = generateRandomList(10000)

	s := New(f)

	s.BubbleSort(&data)

	if !check(&data) {
		t.Errorf("TestBubbleSort Error: %v", check(&data))
		printList(&data)
	}
}

func TestSelectionSort(t *testing.T) {

	var data []interface{} = generateRandomList(10000)

	s := New(f)

	s.SelectionSort(&data)

	if !check(&data) {
		t.Errorf("TestSelectionSort Error: %v", check(&data))
		printList(&data)
	}
}

func TestInsertionSort(t *testing.T) {

	var data []interface{} = generateRandomList(10000)

	s := New(f)

	s.InsertionSort(&data)

	if !check(&data) {
		t.Errorf("TestInsertionSort Error: %v", check(&data))
		printList(&data)
	}
}

func TestShellSort(t *testing.T) {

	var data []interface{} = generateRandomList(10000)

	s := New(f)

	s.ShellSort(&data)

	if !check(&data) {
		t.Errorf("TestShellSort Error: %v", check(&data))
		printList(&data)
	}
}

func TestHeapSort(t *testing.T) {

	var data []interface{} = generateRandomList(1000000)

	s := New(f)

	s.HeapSort(&data)

	if !check(&data) {
		t.Errorf("TestHeapSort Error: %v", check(&data))
		printList(&data)
	}
}

func TestMergeSort(t *testing.T) {

	var data []interface{} = generateRandomList(1000000)

	s := New(f)

	s.MergeSort(&data)

	if !check(&data) {
		t.Errorf("TestMergeSort Error: %v", check(&data))
		printList(&data)
	}
}

func TestQuickSort1(t *testing.T) {

	var data []interface{} = generateRandomList(1000000)

	s := New(f)

	s.QuickSort1(&data)

	if !check(&data) {
		t.Errorf("TestQuickSort1 Error: %v", check(&data))
		printList(&data)
	}
}

func generateRandomList(n int) []interface{} {

	rand.Seed(time.Now().Unix())

	var ret []interface{}

	for i := 0; i < n; i++ {
		ret = append(ret, rand.Intn(n))
	}

	return ret
}

func check(list *[]interface{}) bool {

	for i := 1; i < len(*list); i++ {
		if assist.Compare((*list)[i], (*list)[i-1], f) < 0 {
			return false
		}
	}

	return true
}

func printList(list *[]interface{}) {

	for index, item := range *list {
		if index%10 == 0 {
			fmt.Println()
		}

		fmt.Printf("%v ", item)
	}
}
