package unionfind

// UnionFind
// - 并查集: https://oi-wiki.org/ds/dsu/
// - Introduction to Disjoint Set (Union-Find Algorithm): https://www.geeksforgeeks.org/introduction-to-disjoint-set-data-structure-or-union-find-algorithm/
//
// Optimizations:
// - Path Compression
// - Union by Rank
//
// LeetCode:
// - 952. Largest Component Size by Common Factor: https://leetcode.cn/problems/largest-component-size-by-common-factor/
type UnionFind struct {
	Parent []int
	Rank   []int
}

func New(cap int) *UnionFind {
	parent := make([]int, cap)
	rank := make([]int, cap)
	for i := 0; i < cap; i++ {
		parent[i] = i
	}
	return &UnionFind{
		Parent: parent,
		Rank:   rank,
	}
}

func (uf UnionFind) Find(x int) int {
	if uf.Parent[x] != x {
		// path compression
		uf.Parent[x] = uf.Find(uf.Parent[x])
	}
	return uf.Parent[x]
}

func (uf UnionFind) Union(x, y int) {
	x, y = uf.Find(x), uf.Find(y)
	if x == y {
		return
	}
	// union by rank
	if uf.Rank[x] > uf.Rank[y] {
		uf.Parent[y] = x
	} else if uf.Rank[x] < uf.Rank[y] {
		uf.Parent[x] = y
	} else {
		uf.Parent[x] = y
		uf.Rank[y]++
	}
}

func (uf UnionFind) IsConnected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
