package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	bottom, top , currRowInd:= 0, len(matrix) -1, 0

	for bottom <= top {
	
		mid := (top + bottom) /2
		// its row is range of target value (not requeire to have a values)
		if matrix[mid][0] <= target &&  matrix[mid][len(matrix[mid]) -1] >= target {
			currRowInd = mid
			break
		}

		if matrix[mid][0] < target {
			bottom = mid+1
		}

		if matrix[mid][0] > target {
			top = mid-1
		}
	}

	currRow := matrix[currRowInd]

	left, right := 0, len(currRow) -1
	for left <= right {
		
		mid := (left + right) / 2
		if currRow[mid] == target {
			return true
		}
		if currRow[mid] < target {
			left = mid +1
		}

		if currRow[mid] > target {
			right = mid -1
		}
	}
	return false
}