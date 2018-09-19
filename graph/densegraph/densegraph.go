package densegraph

import (
    "fmt"
)

type DenseGraph struct {
	data [][]bool
	n int
	m int
	directed bool
}

func New(n int, directed bool) *DenseGraph {

	ret := &DenseGraph{
        n: n,
        m: 0,
        data: make([][]bool, n),
        directed: directed,
	}

	for i := 0; i < n; i++ {
		ret.data[i] = make([]bool, n)
	}

	return ret
}

func (g *DenseGraph) AddEdge(p1 int, p2 int) {

	if g.data[p1][p2] {
		return
	}

	g.data[p1][p2] = true
	if !g.directed {
		g.data[p2][p1] = true
	}

	g.m++
}

func (g *DenseGraph) RemoveEdge(p1 int, p2 int) {

	if !g.data[p1][p2] {
		return
	}

	g.data[p1][p2] = false
	if !g.directed {
		g.data[p2][p1] = false
	}

	g.m--
}

func (g *DenseGraph) PointSize() int {
	return g.n
}

func (g *DenseGraph) EdgeSize() int {
	return g.m
}

func (g *DenseGraph) GetIterator(n int) *DenseGraphIterator {
	return &DenseGraphIterator {
		index: -1,
		g: g,
		n: n,
	}
}

func (g *DenseGraph) String() string {

    var ret string = ""

    for i := 0; i < len(g.data); i++ {
    	ret = ret + fmt.Sprintf("%v\n",g.data[i])
    }

    return ret
}

type DenseGraphIterator struct {
	index int
	g *DenseGraph
	n int
}

func (iter *DenseGraphIterator) Begin() int {

	iter.index = -1

	return iter.Next()
}


func (iter *DenseGraphIterator) Next() int {
    
    iter.index++

    for iter.index < iter.g.PointSize() {

    	if iter.g.data[iter.n][iter.index] {
    		return iter.index
    	}

    	iter.index++
    }

    return -1
}

func (iter *DenseGraphIterator) End() bool {
	return iter.index >= iter.g.PointSize()
}