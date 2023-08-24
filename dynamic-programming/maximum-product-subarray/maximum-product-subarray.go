package maximumproductsubarray

import "math"

func maxProduct(nums []int) int {
	min, max := nums[0], nums[0]
	res := max

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}

		min = int(math.Min(float64(nums[i]), float64(nums[i]*min)))
		max = int(math.Max(float64(nums[i]), float64(nums[i]*max)))

		if max > res {
			res = max
		}
	}
	return res
}

// https://leetcode.com/problems/maximum-product-subarray