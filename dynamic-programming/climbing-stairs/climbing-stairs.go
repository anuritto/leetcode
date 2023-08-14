package climbingstairs

func climbStairs(n int) int {
	left, right := 1, 1
	for i := 0; i < n-1; i++ {
		left, right = right, left + right
	}

	return right
}

// https://leetcode.com/problems/climbing-stairs/
