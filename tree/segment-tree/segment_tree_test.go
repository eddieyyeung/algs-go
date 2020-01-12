package segmenttree_test

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"

	// . "github.com/onsi/gomega"

	segmenttree "algs-go/tree/segment-tree"
)

var _ = Describe("SegmentTree", func() {
	arr := []int{1, 3, 5, 7, 9, 11}
	st := segmenttree.New(arr)
	st.UpdateValue(2, 4, 5)
	fmt.Println(st.GetSum(1, 4))
	spew.Dump(st)
})
