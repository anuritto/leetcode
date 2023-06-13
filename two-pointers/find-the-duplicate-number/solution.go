package findtheduplicatenumber

// floyd algorithm
func findDuplicate(nums []int) int {
	hare, turtle := 0, 0

	for {
		turtle = nums[turtle]
		hare = nums[nums[hare]]

		if hare == turtle {
			break
		}
	}

	// now turle = x[i], hare = X[2i], i multiple of cycle length
	// so if tutrle will start by first node,
	// they will meet in cycle start (dublicate)

	turtle = 0

	for hare != turtle {
		turtle = nums[turtle]
		hare = nums[hare]
	}

	return hare
}
