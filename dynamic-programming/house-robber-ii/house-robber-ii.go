package houserobberii

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return maxInt(nums[0], nums[1])
	}

	return maxInt(robWithoutCircle(nums[1:]), robWithoutCircle(nums[:len(nums)-1]))
}


// it's solution of https://leetcode.com/problems/house-robber/
func robWithoutCircle(nums []int) int {

	oneBefore := maxInt(nums[0], nums[1])
	twoBefore := nums[0]

	for i := 2; i < len(nums); i++ {
		withHouse := twoBefore + nums[i]
		withoutHouse := oneBefore

		max := maxInt(withHouse, withoutHouse)
		oneBefore, twoBefore = max, oneBefore
	}

	return oneBefore

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}


// https://leetcode.com/problems/house-robber-ii