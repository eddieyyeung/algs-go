package sorting_test

import (
	"algs-go/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sorting/Test/MergeSort", func() {
	It("[]int{38, 27, 43, 3, 9, 82, 10}", func() {
		arr := []int{38, 27, 43, 3, 9, 82, 10}
		sorting.MergeSort(arr, 0, len(arr)-1)
		Expect(arr).To(Equal([]int{3, 9, 10, 27, 38, 43, 82}))
	})
})
