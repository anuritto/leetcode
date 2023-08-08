package containsduplicate

func containsDuplicate(nums []int) bool {
	existedMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if exists := existedMap[num]; exists {
			return true
		}
		existedMap[num] = true

	}
	return false
}

// https://leetcode.com/problems/contains-duplicate/description/
