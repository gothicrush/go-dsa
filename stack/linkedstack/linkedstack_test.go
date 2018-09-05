package linkedstack

import (
    "testing"
)

func TestInit(t *testing.T) {

	stack := New()

	str := stack.String()
           
	if str != "[ > stack top" {
		t.Errorf("%s\n",str)
	}
}

func TestPush(t *testing.T) {

	stack := New()

	stack.Push(1)
	stack.Push("2")
	stack.Push(true)

	str := stack.String()
           
	if str != "[ 1 -> 2 -> true > stack top" {
		t.Errorf("%s\n",str)
	}
}

func TestPop(t *testing.T) {

	stack := New()

	stack.Push(1)
	stack.Push("2")
	stack.Push(true)

	stack.Pop()
	stack.Pop()

	str := stack.String()
           
	if str != "[ 1 > stack top" {
		t.Errorf("%s\n",str)
	}
}

func TestTop(t *testing.T) {

	stack := New()

	stack.Push(1)
	stack.Push("2")
	stack.Push(true)

	stack.Pop()
	str := stack.Top().(string)

	if str != "2" {
		t.Errorf("TestTop Error: str != 2")
	}
}

func TestSize(t *testing.T) {

	stack := New()

	stack.Push(1)
	stack.Push("2")
	stack.Push(true)

	if stack.Size() != 3 {
		t.Errorf("TestSize Error: size != 3")
	}

	stack.Pop()
	stack.Pop()

	if stack.Size() != 1 {
		t.Errorf("TestSize Error: size != 1")
	}
}

func TestEmpty(t *testing.T) {
	
	stack := New()

	if !stack.Empty() {
		t.Errorf("TestEmpty Error: not empty")
	}

	stack.Push(1)

	if stack.Empty() {
		t.Errorf("TestEmpty Error: empty")
	}

	stack.Pop()
	
	if !stack.Empty() {
		t.Errorf("TestEmpty Error: not empty")
	}
}