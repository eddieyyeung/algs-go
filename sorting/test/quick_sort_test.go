package sorting_test

import (
	"algs-go/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sorting/Test/QuickSort", func() {
	It("[]int{10, 80, 30, 90, 40, 50, 70}", func() {
		arr := []int{10, 80, 30, 90, 40, 50, 70}
		sorting.QuickSort(arr, 0, len(arr)-1)
		Expect(arr).To(Equal([]int{10, 30, 40, 50, 70, 80, 90}))
	})
})
