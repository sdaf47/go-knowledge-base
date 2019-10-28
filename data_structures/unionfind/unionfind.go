package unionfind

type UnionFind struct {
	root []int
}

func NewUnionFind(size int) *UnionFind {
	uf := new(UnionFind)
	uf.root = make([]int, size)

	for i := 0; i < size; i++ {
		uf.root[i] = i
	}

	return uf
}

func (uf *UnionFind) Union(p, q int) {
	qRoot := uf.Root(q)
	pRoot := uf.Root(p)

	uf.root[qRoot] = uf.root[pRoot]
	uf.root[pRoot] = uf.root[qRoot]
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.Root(p) == uf.Root(q)
}

func (uf *UnionFind) Root(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}

	for uf.root[p] != p {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}

	return p
}

func (uf *UnionFind) Find(p int) int {
	return uf.Root(p)
}
