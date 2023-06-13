package containerwithmostwater

// 0(n^2)
func maxArea(heights []int) int {
	max := 0
	n := len(heights)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			width := j - i
			var height int
			if heights[i] > heights[j] {
				height = heights[j]
			} else {
				height = heights[i]
			}
			area := width * height
			if area > max {
				max = area
			}
		}
	}
	return max
}

func maxArea2(heights []int) int {
	n := len(heights)
	start, end := 0, n-1
	max := 0
	for start < end {
		var currArea int
		width := end - start
		if heights[start] < heights[end] {
			currArea = width * heights[start]
			start++
		} else {
			currArea = width * heights[end]
			end--
		}

		if max < currArea {
			max = currArea
		}
	}
	return max

}
