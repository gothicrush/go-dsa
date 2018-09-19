package quickfind

type UnionFind struct {
	id []int
}

func New(n int) *UnionFind {
	ret := &UnionFind{
        id: make([]int, n),
	}

	for i := 0;  i < n; i++ {
        ret.id[i] = i
	}

	return ret
}

func (uf *UnionFind) Size() int {
	return len(uf.id)
}

func (uf *UnionFind) Union(a int, b int) {
    
	var aRoot int = uf.id[a]
    var bRoot int = uf.id[b]

    if aRoot == bRoot {
    	return
    }

    for i := 0; i < len(uf.id); i++ {
    	if uf.id[i] == bRoot {
    		uf.id[i] = aRoot
    	}
    }
}

func (uf *UnionFind) IsConnected(a int, b int) bool {
    var aRoot int = uf.find(a)
    var bRoot int = uf.find(b)

    return aRoot == bRoot
}

func (uf *UnionFind) find(a int) int {
    return uf.id[a]
}