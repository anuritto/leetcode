package maximumsubarray

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := maxSum

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > num+currentSum {
			currentSum = num
		} else {
			currentSum += num
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
