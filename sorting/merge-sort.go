package sorting

// MergeSort https://www.geeksforgeeks.org/merge-sort/
func MergeSort(arr []int, l int, r int) {
	if r > l {
		m := (l + r) / 2
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l int, m int, r int) {
	arrL := make([]int, m-l+1)
	arrR := make([]int, r-m)
	copy(arrL, arr[l:m+1])
	copy(arrR, arr[m+1:r+1])
	i, j := 0, 0
	n := l
	for i < len(arrL) && j < len(arrR) {
		if arrL[i] < arrR[j] {
			arr[n] = arrL[i]
			i++
		} else {
			arr[n] = arrR[j]
			j++
		}
		n++
	}
	for i < len(arrL) {
		arr[n] = arrL[i]
		i++
		n++
	}
	for j < len(arrR) {
		arr[n] = arrR[j]
		j++
		n++
	}
}
