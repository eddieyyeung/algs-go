package union_find

import "testing"

func TestDisjointSet(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ds := New(5)
		ds.Union(0, 2)
		ds.Union(4, 2)
		ds.Union(3, 1)
		if got := ds.GetFather(4) == ds.GetFather(0); got != true {
			t.Errorf("ds.GetFather(4) == ds.GetFather(0) = %v, want %v", got, true)
		}
		if got := ds.GetFather(1) == ds.GetFather(0); got != false {
			t.Errorf("ds.GetFather(1) == ds.GetFather(0) = %v, want %v", got, true)
		}
	})
}
