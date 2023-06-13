package peakindexinamountainarray

func peakIndexInMountainArray(arr []int) int {
	if len(arr) == 3 {
		return arr[1]
	}
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + ((right - left) / 2)

		// increasing part
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else { // decreasing part
			right = mid - 1
		}
	}

	return left
}
