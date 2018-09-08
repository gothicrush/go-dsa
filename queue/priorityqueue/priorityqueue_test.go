package priorityqueue

import (
    "testing"
)

func compare(i interface{}, j interface{}) int {
	var ii int = i.(int)
	var jj int = j.(int)

	return ii - jj
}

func TestInit(t *testing.T) {

	queue := New(compare)

	if queue == nil {
		t.Errorf("TestInit Error")
	}
}

func TestEnQueue(t *testing.T) {

	queue := New(compare)

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	if queue.Size() != 3 {
		t.Errorf("TestEnQueue Error: size: %v", queue.Size())
	}

	if queue.Head().(int) != 3 {
		t.Errorf("TestEnQueue Error: Head: %v", queue.Head().(int))
	}
}

func TestDeQueue(t *testing.T) {

	queue := New(compare)

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(342)
	queue.EnQueue(5)
	queue.EnQueue(22)

	if queue.Head().(int) != 342 {
		t.Errorf("TestDequeue Error: Head(): %v", queue.Head().(int))
	}

	queue.DeQueue()

	if queue.Head().(int) != 22 {
		t.Errorf("TestDequeue Error: %v", queue.Head().(int))
	}
}

func TestHead(t *testing.T) {

	queue := New(compare)

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(342)
	queue.EnQueue(5)
	queue.EnQueue(22)

	if queue.Head().(int) != 342 {
		t.Errorf("TestHead Error: Head(): %v", queue.Head().(int))
	}
}

func TestSize(t *testing.T) {

	queue := New(compare)

	queue.EnQueue(1)
	queue.EnQueue(3)
	queue.EnQueue(2)

	if queue.Size() != 3 {
		t.Errorf("TestSize Error: size != 3")
	}

	queue.DeQueue()
	queue.DeQueue()

	if queue.Size() != 1 {
		t.Errorf("TestSize Error: size != 1")
	}
}

func TestEmpty(t *testing.T) {
	
	queue := New(compare)

	if !queue.Empty() {
		t.Errorf("TestEmpty Error: not empty")
	}

	queue.EnQueue(1)

	if queue.Empty() {
		t.Errorf("TestEmpty Error: empty")
	}

	queue.DeQueue()
	
	if !queue.Empty() {
		t.Errorf("TestEmpty Error: not empty")
	}
}