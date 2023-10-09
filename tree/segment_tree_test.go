package tree

import (
	"fmt"
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	st := NewSegmentTree(1e9)
	arr := []int{10, 11, 12, 13, 14}
	for i, n := range arr {
		st.Update(i, i, n)
	}
	fmt.Println("st.Query(1, 3)", st.Query(1, 3))
	fmt.Println("st.Query(2, 4)", st.Query(2, 4))
	fmt.Println("st.Query(0, 3)", st.Query(0, 3))

	// arr[2] = 3
	st.Update(2, 2, -9)
	fmt.Println("st.Query(1, 3)", st.Query(1, 3))
	fmt.Println("st.Query(2, 4)", st.Query(2, 4))
	fmt.Println("st.Query(0, 3)", st.Query(0, 3))

	// st.Query(1, 3) 36
	// st.Query(2, 4) 39
	// st.Query(0, 3) 46
	// st.Query(1, 3) 27
	// st.Query(2, 4) 30
	// st.Query(0, 3) 37
}
