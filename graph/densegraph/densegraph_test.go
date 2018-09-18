package densegraph

import (
    "testing"
    "fmt"
)

func TestInit(t *testing.T) {

	g := New(10, false)

	if g == nil {
		t.Errorf("TestInit Error")
	}
}

func TestAddEdge(t *testing.T) {

	g := New(5, false)

	g.AddEdge(0,1)
	g.AddEdge(1,0)
	g.AddEdge(1,3)
	g.AddEdge(0,3)
	g.AddEdge(3,0)
	g.AddEdge(3,4)
	g.AddEdge(2,3)
	g.AddEdge(2,4)
	g.AddEdge(1,2)

	if g.PointSize() != 5 || g.EdgeSize() != 7 {
		t.Errorf("TestAddEdge Error: %v %v", g.PointSize(), g.EdgeSize())
	}
 
	fmt.Println(g)
}

func TestRemoveEdge(t *testing.T) {

	g := New(5, false)

	g.AddEdge(0,1)
	g.AddEdge(1,0)
	g.AddEdge(1,3)
	g.AddEdge(0,3)
	g.AddEdge(3,0)
	g.AddEdge(3,4)
	g.AddEdge(2,3)
	g.AddEdge(2,4)
	g.AddEdge(1,2)

    g.RemoveEdge(2,3)

    if g.PointSize() != 5 || g.EdgeSize() != 6 {
		t.Errorf("TestAddEdge Error: %v %v", g.PointSize(), g.EdgeSize())
	}

	fmt.Println(g)
}

func TestPointSize(t *testing.T) {

	g := New(5, false)

	g.AddEdge(0,1)
	g.AddEdge(1,0)
	g.AddEdge(1,3)
	g.AddEdge(0,3)
	g.AddEdge(3,0)
	g.AddEdge(3,4)
	g.AddEdge(2,3)
	g.AddEdge(2,4)
	g.AddEdge(1,2)

	if g.PointSize() != 5 {
		t.Errorf("TestPointSize Error: %v", g.PointSize())
	}
}

func TestEdgeSize(t *testing.T) {

	g := New(5, false)

	g.AddEdge(0,1)
	g.AddEdge(1,0)
	g.AddEdge(1,3)
	g.AddEdge(0,3)
	g.AddEdge(3,0)
	g.AddEdge(3,4)
	g.AddEdge(2,3)
	g.AddEdge(2,4)
	g.AddEdge(1,2)

	if g.EdgeSize() != 7 {
		t.Errorf("TestEdgeSize Error: %v", g.EdgeSize())
	}
}

func TestIterator(t *testing.T) {

	g := New(5, false)

	g.AddEdge(0,1)
	g.AddEdge(1,0)
	g.AddEdge(1,3)
	g.AddEdge(0,3)
	g.AddEdge(3,0)
	g.AddEdge(3,4)
	g.AddEdge(2,3)
	g.AddEdge(2,4)
	g.AddEdge(1,2)

	iter := g.GetIterator(3)

	if iter == nil {
		t.Errorf("TestIterator Error: %v", iter)
	}

    var i int

    i = iter.Begin()
    if iter.End() || i != 0 {
        t.Errorf("TestIterator Begin Error: %v, %v", iter.End(), i)
    }

    i = iter.Next()
    if iter.End() || i != 1 {
        t.Errorf("TestIterator Next Error: %v", i)	
    }

    i = iter.Next()
    if iter.End() || i != 2 {
        t.Errorf("TestIterator Next Error: %v", i)	
    }

    i = iter.Next()
    if iter.End() || i != 4 {
        t.Errorf("TestIterator Next Error: %v", i)	
    }

    i = iter.Next()
    if !iter.End() || i != -1 {
        t.Errorf("TestIterator Next Error: %v, %v", iter.End(), i)	
    }
}