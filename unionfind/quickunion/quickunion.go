package quickunion

type UnionFind struct {
    parent []int
    rank []int
}

func New(n int) *UnionFind {

	ret := &UnionFind{
		parent: make([]int, n),
		rank: make([]int, n),
	}

	for i := 0; i< n; i++ {
		ret.parent[i] = i
		ret.rank[i] = 1
	}

	return ret
}

func (uf *UnionFind) Size() int {

    return len(uf.parent)
}

func (uf *UnionFind) Union(a int, b int) {

	var aRoot int = uf.find(a)
	var bRoot int = uf.find(b)

	if aRoot == bRoot {
		return
	}

	if uf.rank[a] > uf.rank[b] {
        uf.parent[b] = aRoot
	} else if uf.rank[a] < uf.rank[b] {
		uf.parent[a] = bRoot
	} else {
        uf.parent[aRoot] = bRoot
        uf.rank[bRoot] += 1
	}
}

func (uf *UnionFind) IsConnected(a int, b int) bool {

	return uf.find(a) == uf.find(b)
}

func (uf *UnionFind) find(a int) int {

	for a != uf.parent[a] {
		uf.parent[a] = uf.parent[uf.parent[a]]
		a = uf.parent[a]
	}

	return a
}