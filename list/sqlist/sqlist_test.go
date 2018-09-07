package sqlist

import (
	"testing"
)

func TestInit(t *testing.T) {

	li := New()

	str := li.String()

	if str != "[  ]" {
		t.Errorf("%s", str)
	}
}

func TestAdd(t *testing.T) {

	li := New()

	li.Add(0, 0)
	li.Add(1, 1)
	li.Add(2, 2)
	li.Add(3, 3)

	str := li.String()

	if str != "[ 0,1,2,3 ]" {
		t.Errorf("%s", str)
	}

	li.Add(1, 666)

	str = li.String()

	if str != "[ 0,666,1,2,3 ]" {
		t.Errorf("%s", str)
	}
}

func TestRemove(t *testing.T) {

	li := New()

	li.Add(0, 0)
	li.Add(1, 1)
	li.Add(2, 2)
	li.Add(3, 3)

	li.Remove(0)
	li.Remove(0)

	str := li.String()

	if str != "[ 2,3 ]" {
		t.Errorf("%s", str)
	}

	li.Remove(1)
	li.Remove(0)

	str = li.String()

	if str != "[  ]" {
		t.Errorf("%s", str)
	}
}

func TestSet(t *testing.T) {

	li := New()

	li.Add(0, 0)
	li.Add(1, 1)
	li.Add(2, 2)
	li.Add(3, 3)
	li.Add(4, 4)
	li.Add(5, 5)

	li.Set(0, 666)

	str := li.String()

	if str != "[ 666,1,2,3,4,5 ]" {
		t.Errorf("%s", str)
	}
}

func TestGet(t *testing.T) {

	li := New()

	li.Add(0, 0)
	li.Add(1, 1)
	li.Add(2, "gothicrush")
	li.Add(3, 3)
	li.Add(4, 4)
	li.Add(5, 5)

	str := li.Get(2).(string)

	if str != "gothicrush" {
		t.Errorf("%s", str)
	}

	num := li.Get(5).(int)

	if num != 5 {
		t.Errorf("%v", num)
	}
}

func TestContains(t *testing.T) {

	li := New()

	li.Add(0, 0)
	li.Add(1, 1)
	li.Add(2, "gothicrush")
	li.Add(3, 3)
	li.Add(4, 4)
	li.Add(5, 5)

	b := li.Contains("gothicrush")

	if !b {
		t.Errorf("cannot find existing")
	}

	b = li.Contains("one")

	if b {
		t.Errorf("find someting not existing")
	}
}

func TestIndexOf(t *testing.T) {

	li := New()

	li.Add(0, 0)
	li.Add(1, "gothicrush")
	li.Add(2, 3.14)

	index := li.IndexOf(3.14)

	if index != 2 {
		t.Errorf("%d\n", index)
	}

	index = li.IndexOf("gothicrush")

	if index != 1 {
		t.Errorf("%d\n", index)
	}

	index = li.IndexOf(0)

	if index != 0 {
		t.Errorf("%d\n", index)
	}

	index = li.IndexOf("???")

	if index != -1 {
		t.Errorf("%d\n", index)
	}
}

func TestEmpty(t *testing.T) {

	li := New()

	b := li.Empty()

	if !b {
		t.Errorf("not empty?")
	}

	li.Add(0,0)

	b = li.Empty()

	if b {
		t.Errorf("empty?")
	}
}

func TestSize(t *testing.T) {

	li := New()

	size := li.Size()

	if size != 0 {
		t.Errorf("size != 0")
	}

	li.Add(0,0)

	size = li.Size()

	if size != 1 {
		t.Errorf("size != 1")
	}

	li.Add(0,0)

	size = li.Size()

	if size != 2 {
		t.Errorf("size != 2")
	}
}

func TestResize(t *testing.T) {

	li := New()

	for i := 0; i < 21; i++ {
		li.Add(i,i)
	}

	str := li.String()

	if str != "[ 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20 ]" {
		t.Errorf("%s", str)
	}
}