package searchinsertposition

func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1

	for left <= right {
		mid:= (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			if  mid == 0 || nums[mid-1] < target  {
				return mid
			}
			right = mid -1
		} 

		if nums[mid] < target {
			left = mid +1
		}
		
	}
	return len(nums)
}