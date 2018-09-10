package avl

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
	avl := New(comparator)

	if !avl.isBalance() && !avl.isBST() && avl == nil {
		t.Errorf("TestInit Error\n")
	}
}

func TestAdd(t *testing.T) {
	avl := New(comparator)

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	if !avl.isBalance() && !avl.isBST() && avl.Size() != 5 {
		t.Errorf("TestAdd Error: %d\n",avl.Size())
	}
}

func TestIsBST(t *testing.T) {
	avl := New(comparator)

	if !avl.isBST() {
		t.Errorf("TestIsBST Error: %v",avl.isBST())
	}

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	if !avl.isBST() {
		t.Errorf("TestIsBST Error: %v",avl.isBST())
	}
}

func TestIsBalance(t *testing.T) {
	avl := New(comparator)

	if !avl.isBST() {
		t.Errorf("TestIsBalance Error: %v",avl.isBST())
	}

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	if !avl.isBalance() {
		t.Errorf("TestIsBalance Error: %v",avl.isBST())
	}
}

func TestContains(t *testing.T) {
	avl := New(comparator)

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	if b := avl.Contains(9); !avl.isBalance() && !avl.isBST() && !b {
		t.Errorf("TestContains Error: %v", b)
	}

	if b := avl.Contains(666); !avl.isBalance() && !avl.isBST() && b {
		t.Errorf("TestContains Error: %v", b)
	}
}

func TestPreOrder(t *testing.T) {

	avl := New(comparator)

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	var li []interface{} = avl.PreOrder()

	str := fmt.Sprintln(li)

	if !avl.isBalance() && !avl.isBST() && str != "[7 4 1 3 9]\n" {
		t.Errorf("TestPreOrder Error: %s", str)
	}
}

func TestInOrder(t *testing.T) {

	avl := New(comparator)

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	var li []interface{} = avl.InOrder()

	str := fmt.Sprintln(li)

	if !avl.isBalance() && !avl.isBST() && str != "[1 3 4 7 9]\n" {
		t.Errorf("TestInOrder Error: %s", str)
	}
}

func TestPostOrder(t *testing.T) {

	avl := New(comparator)

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	var li []interface{} = avl.PostOrder()

	str := fmt.Sprintln(li)

	if !avl.isBalance() && !avl.isBST() && str != "[3 1 4 9 7]\n" {
		t.Errorf("TestPostOrder Error: %s", str)
	}
}

func TestLevelOrder(t *testing.T) {

	avl := New(comparator)

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	var li []interface{} = avl.LevelOrder()

	str := fmt.Sprintln(li)

	if !avl.isBalance() && !avl.isBST() && str != "[7 4 9 1 3]\n" {
		t.Errorf("TestLevelOrder Error: %s", str)
	}
}

func TestSize(t *testing.T) {

	avl := New(comparator)

	if avl.Size() != 0 {
		t.Errorf("TestSize Error: %v", avl.Size())
	}

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	if !avl.isBalance() && !avl.isBST() && avl.Size() != 5 {
		t.Errorf("TestSize Error: %v", avl.Size())
	}
}


func TestEmpty(t *testing.T) {

	avl := New(comparator)

	if !avl.Empty() {
		t.Errorf("TestSize Error: %v", avl.Empty())
	}

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)
	avl.Add(1)
	avl.Add(3)

	if !avl.isBalance() && !avl.isBST() && avl.Empty() {
		t.Errorf("TestSize Error: %v", avl.Empty())
	}
}

func TestGetMax(t *testing.T) {

	avl := New(comparator)

	avl.Add(7)

	if max := avl.GetMax(); !avl.isBalance() && !avl.isBST() && max != 7 {
		t.Errorf("TestGetMax Error %v", max)
	}

	avl.Add(9)
	avl.Add(4)

	if max := avl.GetMax(); !avl.isBalance() && !avl.isBST() && max != 9 {
		t.Errorf("TestGetMax Error %v", max)
	}

	avl.Add(666)
	avl.Add(3)

	if max := avl.GetMax(); !avl.isBalance() && !avl.isBST() && max != 666 {
		t.Errorf("TestGetMax Error %v", max)
	}
}

func TestGetMin(t *testing.T) {

	avl := New(comparator)

	avl.Add(7)

	if min := avl.GetMin(); !avl.isBalance() && !avl.isBST() && min != 7 {
		t.Errorf("TestGetMin Error %v", min)
	}

	avl.Add(9)
	avl.Add(4)

	if min := avl.GetMin(); !avl.isBalance() && !avl.isBST() && min != 4 {
		t.Errorf("TestGetMin Error %v", min)
	}

	avl.Add(666)
	avl.Add(3)

	if min := avl.GetMin(); !avl.isBalance() && !avl.isBST() && min != 3 {
		t.Errorf("TestGetMin Error %v", min)
	}
}

func TestRemoveMax(t *testing.T) {

	avl := New(comparator)

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)

	max := avl.RemoveMax().(int)

	if mmax := avl.GetMax(); !avl.isBalance() && !avl.isBST() && max != 9 && mmax != 7 {
		t.Errorf("TestRemoveMax Error %v,%v", max,mmax)
	}

	avl.Add(666)
	avl.Add(3)

	max = avl.RemoveMax().(int)

	if mmax := avl.GetMax(); !avl.isBalance() && !avl.isBST() && max != 666 && mmax != 7 {
		t.Errorf("TestRemoveMax Error %v,%v", max,mmax)
	}
}

func TestRemoveMin(t *testing.T) {

	avl := New(comparator)

	avl.Add(7)
	avl.Add(9)
	avl.Add(4)

	min := avl.RemoveMin().(int)

	if mmin := avl.GetMin(); !avl.isBalance() && !avl.isBST() && min != 4 && mmin != 7 {
		t.Errorf("TestRemoveMin Error %v,%v", min,mmin)
	}

	avl.Add(-666)
	avl.Add(3)

	min = avl.RemoveMin().(int)

	if mmin := avl.GetMin(); !avl.isBST() && min != -666 && mmin != 3 {
		t.Errorf("TestRemoveMin Error %v,%v", min,mmin)
	}
}