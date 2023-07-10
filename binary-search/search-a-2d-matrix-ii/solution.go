package searcha2dmatrixii

const INDEX_NOT_FOUND = -1

func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {

		if matrix[i][0] > target || matrix[i][len(matrix[i])-1] < target {
			continue
		}

		if ind := binarySearch(matrix[i], target); ind != INDEX_NOT_FOUND {
			return true
		}
	}

	return false
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}

	}
	return INDEX_NOT_FOUND
}
