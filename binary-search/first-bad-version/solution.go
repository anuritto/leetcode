package firstbadversion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
	var left, right int = 0, n
	
	for left <= right {
		mid := (right + left) / 2

		if isBadVersion(mid) {
			if !isBadVersion(mid-1) {
				return mid
			}

			right = mid -1
		} else {
			left = mid+1
		}	
	}

	return -1
}