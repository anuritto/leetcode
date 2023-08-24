package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, target int) int {
	dp := map[int]int{0: 1}
	n := len(nums)

	for i := 0; i < n; i++ {
		tempDp := make(map[int]int)
		for sum, count := range dp {
			tempDp[sum+nums[i]] += count
			tempDp[sum-nums[i]] += count
		}
		dp = tempDp
	}
	return dp[target]
}

// https://leetcode.com/problems/target-sum

// my first solution
func findTargetSumWays(nums []int, target int) int {
	dp := make([][]int, len(nums)+1)
	dp[0] = []int{0}
	for i := 1; i <= len(nums); i++ {
		dp[i] = make([]int, 0, int(math.Pow(2, float64(i))))
	}

	for i := 0; i < len(nums); i++ {
		prev := dp[i] // because dp = [0, ...other]
		num := nums[i]
		for j := 0; j < len(prev); j++ {
			dp[i+1] = append(dp[i+1], prev[j]-num)
			dp[i+1] = append(dp[i+1], prev[j]+num)
		}
	}

	res := 0
	for i := 0; i < len(dp[len(nums)]); i++ {
		if dp[len(nums)][i] == target {
			res++
		}
	}

	return res
}
