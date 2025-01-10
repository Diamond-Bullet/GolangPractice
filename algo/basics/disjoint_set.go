package basics

// DisjointSet union-find algorithm
// https://www.geeksforgeeks.org/introduction-to-disjoint-set-data-structure-or-union-find-algorithm/
type DisjointSet struct {
	Parent   []int
	Size     []int
	N        int
	SetCount int
}

func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	for i := 1; i < n; i++ {
		parent[i] = i
	}
	return &DisjointSet{Parent: parent,
		Size:     make([]int, n),
		N:        n,
		SetCount: n,
	}
}

func (d *DisjointSet) Find(x int) int {
	if d.Parent[x] == x {
		return x
	}
	d.Parent[x] = d.Find(d.Parent[x])
	return d.Parent[x]
}

// Union let y become the child node of x.
func (d *DisjointSet) Union(x, y int) bool {
	x, y = d.Find(x), d.Find(y)
	if x == y {
		return false
	}
	if d.Size[x] < d.Size[y] {
		x, y = y, x
	}
	d.Parent[y] = x
	d.Size[x] += d.Size[y]
	d.SetCount -= 1
	return true
}

func (d *DisjointSet) Connected(x, y int) bool {
	return d.Find(x) == d.Find(y)
}
