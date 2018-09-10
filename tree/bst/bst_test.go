package bst

import (
	"testing"
	"fmt"
)

func comparator (i interface{}, j interface{}) int {

	ii := i.(int)
	jj := j.(int)

	if ii > jj {
		return 1
	} else if ii == jj {
		return 0
	}

	return -1
}

func TestInit(t *testing.T) {
	bst := New(comparator)

	if !bst.isBST() && bst == nil {
		t.Errorf("TestInit Error\n")
	}
}

func TestAdd(t *testing.T) {
	bst := New(comparator)

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)
	bst.Add(1)
	bst.Add(3)

	if !bst.isBST() && bst.Size() != 5 {
		t.Errorf("TestAdd Error: %d\n",bst.Size())
	}
}

func TestIsBST(t *testing.T) {
	bst := New(comparator)

	if !bst.isBST() {
		t.Errorf("TestIsBST Error: %v",bst.isBST())
	}

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)
	bst.Add(1)
	bst.Add(3)

	if !bst.isBST() && !bst.isBST() {
		t.Errorf("TestIsBST Error: %v",bst.isBST())
	}
}

func TestContains(t *testing.T) {
	bst := New(comparator)

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)
	bst.Add(1)
	bst.Add(3)

	if b := bst.Contains(9); !bst.isBST() && !b {
		t.Errorf("TestContains Error: %v", b)
	}

	if b := bst.Contains(666); !bst.isBST() && b {
		t.Errorf("TestContains Error: %v", b)
	}
}

func TestPreOrder(t *testing.T) {

	bst := New(comparator)

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)
	bst.Add(1)
	bst.Add(3)

	var li []interface{} = bst.PreOrder()

	str := fmt.Sprintln(li)

	if !bst.isBST() && str != "[7 4 1 3 9]\n" {
		t.Errorf("TestPreOrder Error: %s", str)
	}
}

func TestInOrder(t *testing.T) {

	bst := New(comparator)

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)
	bst.Add(1)
	bst.Add(3)

	var li []interface{} = bst.InOrder()

	str := fmt.Sprintln(li)

	if !bst.isBST() && str != "[1 3 4 7 9]\n" {
		t.Errorf("TestInOrder Error: %s", str)
	}
}

func TestPostOrder(t *testing.T) {

	bst := New(comparator)

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)
	bst.Add(1)
	bst.Add(3)

	var li []interface{} = bst.PostOrder()

	str := fmt.Sprintln(li)

	if !bst.isBST() && str != "[3 1 4 9 7]\n" {
		t.Errorf("TestPostOrder Error: %s", str)
	}
}

func TestLevelOrder(t *testing.T) {

	bst := New(comparator)

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)
	bst.Add(1)
	bst.Add(3)

	var li []interface{} = bst.LevelOrder()

	str := fmt.Sprintln(li)

	if !bst.isBST() && str != "[7 4 9 1 3]\n" {
		t.Errorf("TestLevelOrder Error: %s", str)
	}
}

func TestSize(t *testing.T) {

	bst := New(comparator)

	if bst.Size() != 0 {
		t.Errorf("TestSize Error: %v", bst.Size())
	}

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)
	bst.Add(1)
	bst.Add(3)

	if !bst.isBST() && bst.Size() != 5 {
		t.Errorf("TestSize Error: %v", bst.Size())
	}
}


func TestEmpty(t *testing.T) {

	bst := New(comparator)

	if !bst.Empty() {
		t.Errorf("TestSize Error: %v", bst.Empty())
	}

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)
	bst.Add(1)
	bst.Add(3)

	if !bst.isBST() && bst.Empty() {
		t.Errorf("TestSize Error: %v", bst.Empty())
	}
}

func TestGetMax(t *testing.T) {

	bst := New(comparator)

	bst.Add(7)

	if max := bst.GetMax(); !bst.isBST() && max != 7 {
		t.Errorf("TestGetMax Error %v", max)
	}

	bst.Add(9)
	bst.Add(4)

	if max := bst.GetMax(); !bst.isBST() && max != 9 {
		t.Errorf("TestGetMax Error %v", max)
	}

	bst.Add(666)
	bst.Add(3)

	if max := bst.GetMax(); !bst.isBST() && max != 666 {
		t.Errorf("TestGetMax Error %v", max)
	}
}

func TestGetMin(t *testing.T) {

	bst := New(comparator)

	bst.Add(7)

	if min := bst.GetMin(); !bst.isBST() && min != 7 {
		t.Errorf("TestGetMin Error %v", min)
	}

	bst.Add(9)
	bst.Add(4)

	if min := bst.GetMin(); !bst.isBST() && min != 4 {
		t.Errorf("TestGetMin Error %v", min)
	}

	bst.Add(666)
	bst.Add(3)

	if min := bst.GetMin(); !bst.isBST() && min != 3 {
		t.Errorf("TestGetMin Error %v", min)
	}
}

func TestRemoveMax(t *testing.T) {

	bst := New(comparator)

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)

	max := bst.RemoveMax().(int)

	if mmax := bst.GetMax(); !bst.isBST() && max != 9 && mmax != 7 {
		t.Errorf("TestRemoveMax Error %v,%v", max,mmax)
	}

	bst.Add(666)
	bst.Add(3)

	max = bst.RemoveMax().(int)

	if mmax := bst.GetMax(); !bst.isBST() && max != 666 && mmax != 7 {
		t.Errorf("TestRemoveMax Error %v,%v", max,mmax)
	}
}

func TestRemoveMin(t *testing.T) {

	bst := New(comparator)

	bst.Add(7)
	bst.Add(9)
	bst.Add(4)

	min := bst.RemoveMin().(int)

	if mmin := bst.GetMin(); !bst.isBST() && min != 4 && mmin != 7 {
		t.Errorf("TestRemoveMin Error %v,%v", min,mmin)
	}

	bst.Add(-666)
	bst.Add(3)

	min = bst.RemoveMin().(int)

	if mmin := bst.GetMin(); !bst.isBST() && min != -666 && mmin != 3 {
		t.Errorf("TestRemoveMin Error %v,%v", min,mmin)
	}
}