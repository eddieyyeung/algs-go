package sorting_test

import (
	"algs-go/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sorting/Test/BubbleSort", func() {
	It("[]int{64, 34, 25, 12, 22, 11, 90}", func() {
		arr := []int{64, 34, 25, 12, 22, 11, 90}
		sorting.BubbleSort(arr)
		Expect(arr).To(Equal([]int{11, 12, 22, 25, 34, 64, 90}))
	})
})
