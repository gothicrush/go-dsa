package graph

type Graph interface {
	AddEdge(p1 int, p2 int)
	RemoveEdge(p1 int, p2 int)
	PointSize() int
	EdgeSize() int
	GetIterator(n int) Iterator
	String() string
}

