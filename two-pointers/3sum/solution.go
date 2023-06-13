package triplesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < n-2; i++ {
		// for duplicates in i
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		for j, k := i+1, n-1; j < k; {
			// for duplicates in k
			if k < n-2 && nums[k] == nums[k+1] {
				k--
				continue
			}

			// for duplicates in j
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
