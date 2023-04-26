package binarysearch

const NOT_FOUND = -1

func search(nums []int, target int) int {
	return searchInHalf(nums, target, 0)
}

func searchInHalf(nums []int, target, relativeIndex int) int {
	if len(nums) == 0 {
		return NOT_FOUND
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return relativeIndex
		} else {
			return NOT_FOUND
		}
	}

	midInd := len(nums) / 2
	if nums[midInd] == target {
		return midInd + relativeIndex
	}

	if nums[midInd] > target {
		nums = nums[0:midInd]
	} else {
		nums = nums[midInd+1:]
		relativeIndex += midInd + 1
	}

	res := searchInHalf(nums, target, relativeIndex)
	return res
}
