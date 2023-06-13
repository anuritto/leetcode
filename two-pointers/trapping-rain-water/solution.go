package trappingrainwater

func trap(heights []int) int {
	n := len(heights)

	left, right := 0, n-1
	var leftMax, rightMax, res int
	for left < right {
		if heights[left] <= heights[right] {

			if heights[left] < leftMax {
				res += leftMax - heights[left]
			} else {
				leftMax = heights[left]
			}
			left++
		} else {
			if heights[right] < rightMax {
				res += rightMax - heights[right]
			} else {
				rightMax = heights[right]
			}
			right--
		}
	}

	return res
}

// left and right parts of maximum
func trapOld(heights []int) int {
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
