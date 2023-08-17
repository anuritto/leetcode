package coinchange

import (
	"math"
)

const maxInitValue = math.MaxInt8

// BottomUp DP
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = maxInitValue
	}

	for _, c := range coins {
		for a := 1; a <= amount; a++ {
			if a >= c {
				dp[a] = minInt(dp[a], dp[a - c] +1)
			}
		}
	}

	if dp[amount] == maxInitValue {
		return -1
	}

	return dp[amount]
}

func minInt(a,b int) int {
	if a < b {
		return a
	}
	return b
}

// https://leetcode.com/problems/coin-change/
