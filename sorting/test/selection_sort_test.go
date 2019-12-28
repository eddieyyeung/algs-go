package sorting_test

import (
	"algs-go/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sorting/Test/SelectionSort", func() {
	It("[]int{5, 1, 4, 2, 8}", func() {
		arr := []int{5, 1, 4, 2, 8}
		sorting.SelectionSort(arr)
		Expect(arr).To(Equal([]int{1, 2, 4, 5, 8}))
	})
})
