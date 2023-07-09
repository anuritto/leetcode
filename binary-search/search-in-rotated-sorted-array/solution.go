package searchinrotatedsortedarray


func search(nums []int, target int) int {
	l, r := 0, len(nums) -1

	for l <= r {

		mid := (l + r) /2
		
		if nums[mid] == target {
			return mid
		}

		// mid in right part of slice
		if nums[mid] < nums[l] {

			// target in right part of slice
			if nums[mid] < target && target <= nums[r] {
				l = mid +1
			} else {  // target in left part of slice
				r = mid -1
			}
		} else { // mid in left part of slice
			if nums[mid] > target && target >= nums[l] { // target in left part
				r = mid-1
			} else {
				l = mid +1
			}
		 }
	}
	return -1
}