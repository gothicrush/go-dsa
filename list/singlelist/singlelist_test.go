package singlelist

import (
    "testing"
)

func TestInit(t *testing.T) {
	li := New()

	str := li.String()

	if (str != "[ ]") {
		t.Errorf(str)
	}
}

func TestAdd(t *testing.T) {
	li := New()

	li.Add(0,0)
	li.Add(1,1)
	li.Add(2,2)
	li.Add(3,3)
	li.Add(4,4)

	str := li.String()

	if (str != "[ 0 -> 1 -> 2 -> 3 -> 4 ]") {
		t.Errorf(str)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("cannot do panic")
		}	
	}()

	li.Add(-1,-1)
}

func TestRemove(t *testing.T) {
	li := New()

	li.Add(0,0)
	li.Add(1,1)
	li.Add(2,2)
	li.Add(3,3)
	li.Add(4,4)

	li.Remove(1)
	li.Remove(3)

	str := li.String()

	if (str != "[ 0 -> 2 -> 3 ]") {
		t.Errorf(str)
	}

	li.Remove(0)
	li.Remove(0)
	li.Remove(0)

	str = li.String()

	if (str != "[ ]") {
		t.Errorf(str)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("cannot do panic")
		}	
	}()

	li.Remove(6666)
}

func TestSet(t *testing.T) {
	li := New()

	li.Add(0,0)
	li.Add(1,1)
	li.Add(2,2)

	li.Set(1,666)

	str := li.String()

	if (str != "[ 0 -> 666 -> 2 ]") {
		t.Errorf(str)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("cannot do panic")
		}	
	}()

	li.Set(666,666)
}

func TestGet(t *testing.T) {
	// t.SkipNow()
	li := New()

	li.Add(0,"1")
	li.Add(1,"testGet")
	li.Add(2,2)

	str := li.Get(1)

	if str != "testGet" {
		t.Errorf("%v", str)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("cannot do panic")
		}	
	}()

	li.Get(-1)
}

func TestContains(t *testing.T) {

	li := New()

	li.Add(0,0)
	li.Add(1,1)
	li.Add(2,2)

	b := li.Contains(1)

	if !b {
		t.Errorf("%v",b)
	}

	b = li.Contains(666)

	if b {
		t.Errorf("%v",b)
	}
}

func TestEmpty(t *testing.T) {
	li := New()

	b := li.Empty()

	if !b {
		t.Errorf("%v",b)
	}

	li.Add(0,0)

	b = li.Empty()

	if b {
		t.Errorf("%v",b)
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
}

func TestIndexOf(t *testing.T) {
	li := New()

	li.Add(0,0)
	li.Add(1,"666")
	li.Add(2,2)

	index := li.IndexOf("666")

	if index != 1 {
		t.Error("index != 1")
	}

	index = li.IndexOf(666)

	if index != -1 {
		t.Error("index != -1")
	}
}