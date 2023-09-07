package unionfind

import "testing"

func TestUnionFind(t *testing.T) {
	uf := New(5)
	uf.Union(0, 2)
	uf.Union(4, 2)
	uf.Union(3, 1)
	if uf.Find(4) != uf.Find(0) {
		t.Errorf("4 and 0 should be in union\n")
	}
	if uf.Find(1) == uf.Find(0) {
		t.Errorf("1 and 0 should not be in union")
	}
}
