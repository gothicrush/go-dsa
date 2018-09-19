package graphiterator

type GraphIterator interface {
	Begin() int
	Next() int
	End() bool
}