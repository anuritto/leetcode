package searchinrotatedsortedarray

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (right + left) / 2

		if nums[mid] == target {
			return mid
		}

		// second part of array
		if nums[mid] < nums[left] {
			// evenly rising to the right of the middle
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// evenly descending to the left of the middle
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

	}
	return -1
}

// https://leetcode.com/problems/search-in-rotated-sorted-array/
