package trappingrainwater

func trap(heights []int) int {
	n := len(heights)
	var maxInd, maxHeight int
	for i := 0; i < n; i++ {
		if heights[i] > maxHeight {
			maxHeight = heights[i]
			maxInd = i
		}
	}

	res := 0
	if maxInd > 0 {
		for left, maxLeftHeight := 0, 0; left < maxInd; left++ {
			if heights[left] >= maxLeftHeight {
				maxLeftHeight = heights[left]

			} else {
				res += maxLeftHeight - heights[left]
			}
		}
	}

	if maxInd < n-1 {
		for right, maxRightHeight := n-1, 0; right > maxInd; right-- {
			if heights[right] >= maxRightHeight {
				maxRightHeight = heights[right]
			} else {
				res += maxRightHeight - heights[right]
			}
		}
	}
	return res
}
