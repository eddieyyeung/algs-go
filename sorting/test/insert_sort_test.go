package sorting_test

import (
	"algs-go/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sorting/Test/InsertSort", func() {
	It("[]int{12, 11, 13, 5, 6}", func() {
		arr := []int{12, 11, 13, 5, 6}
		sorting.InsertSort(arr)
		Expect(arr).To(Equal([]int{5, 6, 11, 12, 13}))
	})
})
