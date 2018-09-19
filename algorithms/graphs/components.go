package graphs

import (
    "../../graph"
)

type Components struct {
	visited []bool
	count int
	id []int
	g graph.Graph
}

func New(g graph.Graph) *Components {

    ret := &Components {
    	visited: make([]bool, g.PointSize()),
    	id: make([]int, g.PointSize()),
    	count: 0,
    	g: g,
    }

    for i := 0; i < g.PointSize(); i++ {
    	ret.id[i] = -1
    }

    for i := 0; i < g.PointSize(); i++ {
    	if !ret.visited[i] {
    		ret.count++
    		ret.dsa(i)
    	}
    }

    return ret
}

func (c *Components) Size() int {
	return c.count
}

func (c *Components) IsConnected(a int, b int) bool {
    return c.id[a] == c.id[b]
}

func (c *Components) dsa(n int) {

	c.visited[n] = true
    c.id[n] = c.count

    var iter graph.Iterator = c.g.GetIterator(n)

    for i := iter.Begin(); !iter.End(); i = iter.Next() {
    	if !c.visited[i] {
            c.dsa(i)
    	}
    }
}
