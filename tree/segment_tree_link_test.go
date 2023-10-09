package tree

import (
	"fmt"
	"testing"
)

func TestNewSegmentTreeLink(t *testing.T) {
	stl := NewSegmentTreeLink([]int{10, 11, 12, 13, 14})
	fmt.Println("str", stl)
	fmt.Println("str.SumRange(1, 3)", stl.SumRange(1, 3))
	fmt.Println("str.SumRange(2, 4)", stl.SumRange(2, 4))
	fmt.Println("str.SumRange(0, 3)", stl.SumRange(0, 3))

	stl.Update(2, 3)
	fmt.Println("str.SumRange(1, 3)", stl.SumRange(1, 3))
	fmt.Println("str.SumRange(2, 4)", stl.SumRange(2, 4))
	fmt.Println("str.SumRange(0, 3)", stl.SumRange(0, 3))
}
