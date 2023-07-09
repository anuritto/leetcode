package searchinrotatedsortedarray

import "fmt"

func search(nums []int, target int) int {
	length, k := len(nums), 0
	fmt.Println("length= ",length)
	// if k != 0
	if nums[0] > nums[length-1] {

		left, right := 0, length -1
		for left <= right  {
			mid := (left + right) /2
			fmt.Println("left=", left, " right=",right,  " mid=", mid)

			if mid == 0 {
				k = length- 1
				break
			}

			if nums[mid] <= nums[mid-1] {
				k = length- mid
				break
			}
			if nums[mid] > nums[0] {
				left = mid +1
			}
			if nums[mid] < nums[0] {
				right = mid -1
			}
		}
	}

	fmt.Println("k=", k)
	sortedNums := append(nums[length - k:], nums[:length - k]...)

	left, right := 0, length -1
	for left <= right {
		mid := (left + right) /2
		fmt.Println("left=", left, " right=",right,  " mid=", mid)
		if sortedNums[mid] == target {
			unsortedInd := mid - k
			if unsortedInd < 0 {
				unsortedInd += length
			}
			return unsortedInd
		}

		if sortedNums[mid] > target {
			right = mid -1
		}
		
		if sortedNums[mid] < target {
			left = mid +1
		}
	}
	return -1

}