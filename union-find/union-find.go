package union_find

type DisjointSet struct {
	Father []int
	Rank   []int
}

func New(cap int) *DisjointSet {
	father := make([]int, cap)
	rank := make([]int, cap)
	for i := 0; i < len(father); i++ {
		father[i] = i
	}
	return &DisjointSet{
		Father: father,
		Rank:   rank,
	}
}

func (ds *DisjointSet) IsSame(x int, y int) bool {
	return ds.GetFather(x) == ds.GetFather(y)
}

func (ds *DisjointSet) GetFather(i int) int {
	if ds.Father[i] != i {
		ds.Father[i] = ds.GetFather(ds.Father[i])
	}
	return ds.Father[i]
}

func (ds *DisjointSet) Union(x int, y int) {
	fx := ds.GetFather(x)
	fy := ds.GetFather(y)
	if fx == fy {
		return
	}
	if ds.Rank[fx] > ds.Rank[fy] {
		ds.Father[fy] = fx
	} else {
		ds.Father[fx] = fy
		if ds.Rank[fx] == ds.Rank[fy] {
			ds.Rank[fy]++
		}
	}
}
