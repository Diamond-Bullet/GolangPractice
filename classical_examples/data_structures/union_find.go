package data_structures

// UnionFind 并查集
type UnionFind struct {
	Parent   []int
	Size     []int
	N        int
	SetCount int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 1; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{Parent: parent,
		Size:     make([]int, n),
		N:        n,
		SetCount: n,
	}
}

func (uF *UnionFind) FindSet(x int) int {
	if uF.Parent[x] == x {
		return x
	}
	uF.Parent[x] = uF.FindSet(uF.Parent[x])
	return uF.Parent[x]
}

func (uF *UnionFind) Unite(x, y int) bool {
	x, y = uF.FindSet(x), uF.FindSet(y)
	if x == y {
		return false
	}
	if uF.Size[x] < uF.Size[y] {
		x, y = y, x
	}
	uF.Parent[y] = x
	uF.Size[x] += uF.Size[y]
	uF.SetCount -= 1
	return true
}

func (uF *UnionFind) Connected(x, y int) bool {
	return uF.FindSet(x) == uF.FindSet(y)
}
