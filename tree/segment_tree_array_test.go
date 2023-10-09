package tree

import (
	"fmt"
	"testing"
)

func TestNewSegmentTreeArray(t *testing.T) {
	str := NewSegmentTreeArray([]int{10, 11, 12, 13, 14})
	fmt.Println("str", str)
	fmt.Println("str.SumRange(1, 3)", str.SumRange(1, 3))
	fmt.Println("str.SumRange(2, 4)", str.SumRange(2, 4))
	fmt.Println("str.SumRange(0, 3)", str.SumRange(0, 3))

	str.Update(2, 3)
	fmt.Println("str.SumRange(1, 3)", str.SumRange(1, 3))
	fmt.Println("str.SumRange(2, 4)", str.SumRange(2, 4))
	fmt.Println("str.SumRange(0, 3)", str.SumRange(0, 3))
}
