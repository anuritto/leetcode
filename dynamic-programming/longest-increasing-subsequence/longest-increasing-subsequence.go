package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	sub := []int{}

	for i := 0; i < len(nums); i++ {
		if len(sub) == 0 {
			sub = []int{nums[i]}
			continue
		}

		if sub[len(sub)-1] < nums[i] {
			sub = append(sub, nums[i])
			continue
		}

		ind := binarySearch(sub, nums[i])

		sub[ind] = nums[i]
	}

	return len(sub)
}

// not strong search
// return not only index if target is found, but also index for target instead of element of this target
func binarySearch(nums []int, target int) int {

	left, right := 0, len(nums) -1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// https://leetcode.com/problems/longest-increasing-subsequence/
