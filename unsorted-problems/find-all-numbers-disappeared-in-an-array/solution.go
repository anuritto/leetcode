package solution

func findDisappearedNumbers(nums []int) []int {
	max := len(nums)

	sorted := make([]int, max) // we can use also boolean, true == num exist
	for i := 0; i < max; i++ {
		val := nums[i]
		// right position for N is N-1
		sorted[val-1] = val
	}

	notexist := []int{}
	for i := 0; i < max; i++ {
		if sorted[i] == 0 { // we know that all nums > 0
			notexist = append(notexist, i+1) // if notexist[i] is default => num i+1 absent
		}
	}

	return notexist
}
