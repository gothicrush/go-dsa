package graph

type Iterator interface {
	Begin() int
	Next() int
	End() bool
}