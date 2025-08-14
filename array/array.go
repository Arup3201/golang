package array

func Search1DArray(arr []int, x int) int {
	for i := range len(arr) {
		if x == arr[i] {
			return i
		}
	}

	return -1
}

func BinarySearch(arr []int, x int) int {
	var low, high int = 0, len(arr)

	for low <= high {
		mid := low + (high-low)/2
		if x == arr[mid] {
			return mid
		} else if x > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
