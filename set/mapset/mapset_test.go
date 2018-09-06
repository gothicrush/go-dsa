package mapset

import (
	"testing"
)

func TestInit(t *testing.T) {

	s := New()

	str := s.String()

	if str != "( )" {
		t.Errorf("TestInit Error: %s", str)
	}
}

func TestContains(t *testing.T) {

	s := New()

	s.Add(1)
	s.Add(2)
	s.Add(666)

	b := s.Contains(666)

	if !b {
		t.Errorf("TestContains Error: %v", b)
	}

	b = s.Contains(777)

	if b {
		t.Errorf("TestContains Error: %v", b)
	}
}


func TestAdd(t *testing.T) {

	s := New()

	s.Add(1)
	s.Add(2)
	s.Add(666)

	if !s.Contains(1) || !s.Contains(2) || !s.Contains(666) {
		t.Errorf("TestAdd Error")
	}
}

func TestRemove(t *testing.T) {

	s := New()

	s.Add(1)
	s.Add(2)
	s.Add(666)

	if !s.Contains(1) || !s.Contains(2) || !s.Contains(666) {
		t.Errorf("TestRemove Error")
	}

	s.Remove(2)
	s.Remove(1)

	if s.Contains(1) || s.Contains(2) || !s.Contains(666) {
		t.Errorf("TestRemove Error")
	}
}


func TestEquals(t *testing.T) {

	s := New()

	s.Add(1)
	s.Add(2)
	s.Add(666)

	ss := New()

	ss.Add(1)
	ss.Add(2)
	ss.Add(666)

	b := s.Equals(ss)

	if !b {
		t.Errorf("TestEquals Error: %v", b)
	}

	ss.Add(777)

	b = s.Equals(ss)

	if b {
		t.Errorf("TestEquals Error: %v", b)
	}
}

func TestSubset(t *testing.T) {

	s := New()

	s.Add(1)
	s.Add(2)
	s.Add(666)

	ss := New()

	ss.Add(1)

	b := s.Subset(ss)

	if !b {
		t.Errorf("TestSubset Error: %v", b)
	}

	ss.Add(777)

	b = s.Subset(ss)

	if b {
		t.Errorf("TestSubset Error: %v", b)
	}	
}

func TestInter(t *testing.T) {

	s := New()

	s.Add(1)
	s.Add(2)
	s.Add(666)

	ss := New()

	ss.Add(1)
	ss.Add(5)
	ss.Add(6)

	sss := s.Inter(ss)

	if !sss.Contains(1) || sss.Contains(2) || sss.Contains(666) || 
	    sss.Contains(5) || sss.Contains(6) {

		t.Errorf("TestInter Error")
	}
}

func TestDiff(t *testing.T) {

	s := New()

	s.Add(1)
	s.Add(2)
	s.Add(666)

	ss := New()

	ss.Add(1)
	ss.Add(5)
	ss.Add(6)

	sss := s.Diff(ss)

	if  sss.Contains(1) || !sss.Contains(2) || !sss.Contains(666) || 
	   !sss.Contains(5) || !sss.Contains(6) {

		t.Errorf("TestDiff Error")
	}
}

func TestUnion(t *testing.T) {

	s := New()

	s.Add(1)
	s.Add(2)
	s.Add(666)

	ss := New()

	ss.Add(1)
	ss.Add(5)
	ss.Add(6)

	sss := s.Union(ss)

	if !sss.Contains(1) || !sss.Contains(2) || !sss.Contains(666) || 
	   !sss.Contains(5) || !sss.Contains(6) {

		t.Errorf("TestUnion Error")
	}
}

func TestEmpty(t *testing.T) {

	s := New()

	if b := s.Empty(); !b {
		t.Errorf("TestEmpty Error: %v", b)
	}

	s.Add(1)
	s.Add(2)
	s.Add(666)

	if b := s.Empty(); b {
		t.Errorf("TestEmpty Error: %v", b)
	}
}

func TestClear(t *testing.T) {

	s := New()

	if b := s.Empty(); !b {
		t.Errorf("TestClear Error: %v", b)
	}

	s.Add(1)
	s.Add(2)
	s.Add(666)

	if b := s.Empty(); b {
		t.Errorf("TestClear Error: %v", b)
	}

	s.Clear()

	if b := s.Empty(); !b {
		t.Errorf("TestClear Error: %v", b)
	}
}

func TestSize(t *testing.T) {

	s := New()

	if b := s.Size(); b != 0 {
		t.Errorf("TestClear Error: %v", b)
	}

	s.Add(1)
	s.Add(2)
	s.Add(666)

	if b := s.Size(); b != 3 {
		t.Errorf("TestClear Error: %v", b)
	}

	s.Clear()

	if b := s.Size(); b != 0 {
		t.Errorf("TestClear Error: %v", b)
	}

}