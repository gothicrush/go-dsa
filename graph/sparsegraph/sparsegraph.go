package sparsegraph

import (
    "fmt"
)

type SparseGraph struct {
	data [][]int
	n int
	m int
	directed bool
}

func New(n int, directed bool) *SparseGraph {

	return &SparseGraph{
        n: n,
        m: 0,
        data: make([][]int, n),
        directed: directed,
	}
}

func (g *SparseGraph) AddEdge(p1 int, p2 int) {

	if g.hasEdge(p1, p2) {
		return
	}

	g.data[p1] = append(g.data[p1], p2)
	if !g.directed {
		g.data[p2] = append(g.data[p2], p1)
	}

	g.m++
}

func (g *SparseGraph) hasEdge(p1 int, p2 int) bool {

    for i := 0; i < len(g.data[p1]); i++ {
    	if g.data[p1][i] == p2 {
    		return true
    	}
    }

    return false
}

func (g *SparseGraph) RemoveEdge(p1 int, p2 int) {

	if !g.hasEdge(p1, p2) {
		return
	}

	for i := 0; i < len(g.data[p1]); i++ {
		if g.data[p1][i] == p2 {
			g.data[p1] = append(g.data[p1][:i], g.data[p1][i+1:]...)
		}
	}


	if !g.directed {
		for i := 0; i < len(g.data[p2]); i++ {
			if g.data[p2][i] == p1 {
				g.data[p2] = append(g.data[p2][:i], g.data[p2][i+1:]...)
			}
		}
	}

	g.m--
}

func (g *SparseGraph) PointSize() int {
	return g.n
}

func (g *SparseGraph) EdgeSize() int {
	return g.m
}

func (g *SparseGraph) GetIterator(n int) *SparseGraphIterator {
	return &SparseGraphIterator {
		index: -1,
		g: g,
		n: n,
	}
}

func (g *SparseGraph) String() string {

    var ret string = ""

    for i := 0; i < len(g.data); i++ {
    	ret = ret + fmt.Sprintf("%v: %v\n", i, g.data[i])
    }

    return ret
}

type SparseGraphIterator struct {
	index int
	g *SparseGraph
	n int
}

func (iter *SparseGraphIterator) Begin() int {

	iter.index = -1

	return iter.Next()
}


func (iter *SparseGraphIterator) Next() int {
    
    iter.index++

    if iter.End() {
        return -1
    }

    return iter.g.data[iter.n][iter.index]
}

func (iter *SparseGraphIterator) End() bool {
	return iter.index >= len(iter.g.data[iter.n])
}