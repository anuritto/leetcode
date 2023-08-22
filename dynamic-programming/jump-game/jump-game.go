package jumpgame

// calculation paths count for each element
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	dp := make([]int, len(nums))

	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if dp[i] == 0 {
			return false
		}

		for j := i + 1; j < i+1+num && j < len(nums); j++ {

			dp[j]++
		}
	}

	return dp[len(nums)-1] > 0
}

// https://leetcode.com/problems/jump-game
