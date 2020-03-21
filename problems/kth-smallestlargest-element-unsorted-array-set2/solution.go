package kth_smallestlargest_element_unsorted_array_set2

func kthSmallest(arr []int, k int) int {
	mh := New(len(arr))
	for _, v := range arr {
		mh.Insert(v)
	}
	for i := 0; i < k-1; i++ {
		mh.ExtractMin()
	}
	min, _ := mh.ExtractMin()
	return min
}
