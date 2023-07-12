package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {

	n := len(nums)
	if n <= 1 {
		return n
	}

	left := 0
	for right := 1; right < n; right++ {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
	}

	return left + 1
}
