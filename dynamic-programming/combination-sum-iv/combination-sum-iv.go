package combinationsumiv

// bottom up method
// count variation of target T = sum of count(target T-num1) +  count(target T-num2) + ...  count(target T-numN)
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // for target==0 only one variation
	for t := 1; t <= target; t++ {
		for _, num := range nums {
			if t >= num {
				dp[t] += dp[t-num]
			}
		}
	}
	return dp[target]
}

// https://leetcode.com/problems/combination-sum-iv