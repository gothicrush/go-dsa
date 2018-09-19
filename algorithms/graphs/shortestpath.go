package graphs

import (
    "../../graph"
    "../../queue/linkedqueue"
    "../../stack/slicestack"
    "fmt"
    "strconv"
)

type ShortestPath struct {
	visited []bool
	ord []int
	from []int
    g graph.Graph
    start int
}

func New(g graph.Graph, start int) *ShortestPath {

	ret := &ShortestPath {
		visited: make([]bool, g.PointSize()),
		ord: make([]int, g.PointSize()),
		from: make([]int, g.PointSize()),
		g: g,
		start: start,
	}

	for i :=0; i < g.PointSize(); i++ {
		ret.visited[i] = false
		ret.ord[i] = -1
		ret.from[i] = -1
	}

	var q *linkedqueue.Queue = linkedqueue.New()

	q.EnQueue(start)
    ret.visited[start] = true
    ret.ord[start] = 0

    for !q.Empty() {

    	var v int = q.Head().(int)
    	q.DeQueue()

        var iter graph.Iterator = g.GetIterator(start)
        for i := iter.Begin(); !iter.End(); i = iter.Next() {
        	if !ret.visited[i] {
        		q.EnQueue(i)
        		ret.visited[i] = true
        		ret.from[i] = v
        		ret.ord[i] = ret.ord[v] + 1
        	}
        }
    }


	return ret
}

func (sp *ShortestPath) HasPath(target int) bool {
	return sp.visited[target]
}

func (sp *ShortestPath) GetPath(target int) []int {

	var st *slicestack.Stack = slicestack.New()

	for p := target; p != -1; p = sp.from[p] {
        st.Push(p)
	}

	var ret []int

	for !st.Empty() {
		ret = append(ret, st.Top().(int))
		st.Pop()
	}
    
    return ret
}

func (sp *ShortestPath) ShowPath(target int) string {
    if !sp.HasPath(target) {
    	return fmt.Sprintf("cannot reach %v", target)
    }

    var paths []int = sp.GetPath(target)

    var ret string = ""

    for index, item := range paths {
    	if index != len(paths) - 1 {
    		ret = ret + strconv.Itoa(item) + " ->"
		} else {
			ret = ret + strconv.Itoa(item)
		}
    }

    return ret
}