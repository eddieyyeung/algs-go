package sorting

// BubbleSort https://www.geeksforgeeks.org/bubble-sort/
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// if no two elements were swapped by inner loop, then break
		if swapped == false {
			break
		}
	}
}
