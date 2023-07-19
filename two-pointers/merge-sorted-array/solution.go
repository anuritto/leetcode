package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int)  {
    first, second:= 0,0
	res := []int{}
	for first < len(nums1) || second < len(nums2) {

		if first >= len(nums1) || nums1[first] <= nums2[second]  {
			res = append(res, nums1[first])
			first++
			continue
		}


		if second >= len(nums2) || nums1[first] > nums2[second]  {
			res = append(res, nums2[second])
			second++
			continue
		}
	}

	return res
}