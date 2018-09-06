package linkedqueue

import (
    "testing"
)

func TestInit(t *testing.T) {

	queue := New()

	str := queue.String()
               
	if str != ">>>   >>> queue head" {
		t.Errorf("%s\n",str)
	}
}

func TestEnQueue(t *testing.T) {

	queue := New()

	queue.EnQueue(1)
	queue.EnQueue("2")
	queue.EnQueue(true)

	str := queue.String()
              
	if str != ">>>  1 -> 2 -> true  >>> queue head" {
		t.Errorf("%s\n",str)
	}
}

func TestDeQueue(t *testing.T) {

	queue := New()

	queue.EnQueue(1)
	queue.EnQueue("2")
	queue.EnQueue(true)

	queue.DeQueue()
	queue.DeQueue()

	str := queue.String()
           
	if str != ">>>  true  >>> queue head" {
		t.Errorf("%s\n",str)
	}
}

func TestHead(t *testing.T) {

	queue := New()

	queue.EnQueue(1)
	queue.EnQueue("2")
	queue.EnQueue(true)

	queue.DeQueue()
	str := queue.Head().(string)

	if str != "2" {
		t.Errorf("TestTop Error: str != 2")
	}
}

func TestSize(t *testing.T) {

	queue := New()

	queue.EnQueue(1)
	queue.EnQueue("2")
	queue.EnQueue(true)

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
	
	queue := New()

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