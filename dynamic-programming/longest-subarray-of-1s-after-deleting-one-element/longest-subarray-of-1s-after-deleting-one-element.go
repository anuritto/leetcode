package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 0, 1, 1, 1, 0, 1}))
}

func longestSubarray(nums []int) int {
	zeroCount := 0
	l, r := 0, 0
	for ; r < len(nums); r++ {
		if nums[r] == 0 {
			zeroCount++
		}

		if zeroCount > 1 {
			if nums[l] == 0 {
				zeroCount--
			}
			l++
		}
	}

	return r - l - 1
}

func longestSubarray(nums []int) int {
	subArrs := []int{}
	currLeng := 0
	for _, num := range nums {
		if num > 0 {
			currLeng++
		} else {
			if currLeng == 0 {
				subArrs = append(subArrs, 0)
			} else {
				subArrs = append(subArrs, currLeng)
				currLeng = 0
			}
		}
	}
	if currLeng != 0 {
		subArrs = append(subArrs, currLeng)
	}

	n := len(subArrs)
	if n == 1 {
		if subArrs[0] == len(nums) {
			return subArrs[0] - 1
		} else {
			return subArrs[0]
		}
	}

	max := 0
	for i := 0; i < n-1; i++ {
		sumWithNext := subArrs[i] + subArrs[i+1]
		if sumWithNext > max {
			max = sumWithNext
		}
	}

	return max
}
