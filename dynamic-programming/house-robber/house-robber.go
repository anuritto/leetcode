package houserobber

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxInt(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		withHouse := dp[i-2] + nums[i]
		withoutHouse := dp[i-1]

		dp[i] = maxInt(withHouse, withoutHouse)
	}

	return dp[len(dp)-1]

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}


// https://leetcode.com/problems/house-robber/description/