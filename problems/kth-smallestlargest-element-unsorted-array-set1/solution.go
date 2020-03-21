// https://www.geeksforgeeks.org/kth-smallestlargest-element-unsorted-array/
package kth_smallestlargest_element_unsorted_array_set1

import "sort"

func kthSmallest(arr []int, k int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[k-1]
}
