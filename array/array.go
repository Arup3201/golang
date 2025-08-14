package array

func Search1DArray(arr []int, x int) int {
	for i := range len(arr) {
		if x == arr[i] {
			return i
		}
	}

	return -1
}
