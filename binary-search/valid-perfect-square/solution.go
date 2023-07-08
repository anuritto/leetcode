package validperfectsquare


func isPerfectSquare(num int) bool {
    
	var left,right int = 0, num

	for left <= right {
		mid := (left +right ) / 2

		midPow := mid*mid
		if midPow == num {
			return true
		} else if midPow > num {
			right = mid -1
		} else {
			left = mid +1
		}
	}
	return false
}
