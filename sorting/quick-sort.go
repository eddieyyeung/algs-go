package sorting

// QuickSort https://www.geeksforgeeks.org/quick-sort/
func QuickSort(arr []int, low int, high int) {
	if low < high {
		// pi is partitioning index, arr[pi] is now at right place
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

/*
	partition  This function takes last element as pivot,
	places the pivot element at its correct position in sorted array,
	and places all smaller (smaller than pivot) to left of pivot
	and all greater elements to right of pivot
*/
func partition(arr []int, low int, high int) int {
	// pivot (Element to be placed at right position)
	pivot := arr[high]
	s := low - 1
	for i := low; i <= high-1; i++ {
		if arr[i] < pivot {
			s++
			arr[i], arr[s] = arr[s], arr[i]
		}
	}
	s++
	arr[s], arr[high] = arr[high], arr[s]
	return s
}
