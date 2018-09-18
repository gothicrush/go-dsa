package graph

type GraphIterator interface {
	begin() int
	next() int
	end() bool
}