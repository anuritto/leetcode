package twosum

import "fmt"


func twoSum(nums []int, target int) []int {
	// key - num from nums, value - index
    numMap :=make(map[int]int)

	// filling map
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}
	fmt.Println(numMap)

	var left, right int
	for i := 0; i < len(nums); i++ {
		if j, exist := numMap[target - nums[i]]; exist && i != j {
			left, right = i,j;
			break
		}
		
	}
	return []int{left,right}
}

// https://leetcode.com/problems/two-sum