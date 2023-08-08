package productofarrayexceptself


func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	leftMultiplier := 1 
	for i := range nums {
		res[i] = leftMultiplier
		leftMultiplier *= nums[i]
	}

	rightMultiplier := 1
	for i := len(nums) -1 ; i >= 0; i--{
		res[i] *= rightMultiplier
		rightMultiplier*=nums[i]
	}
	return res
}

func productExceptSelfOld(nums []int) []int {
	product := 1
	zeroCount:= 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			product *= nums[i]
		} else {
			zeroCount++
		}

	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if zeroCount > 1 {
			res = append(res, 0)
			continue
		}

		if zeroCount == 1 && nums[i] != 0 {
			res = append(res, 0)
		} else if zeroCount == 1 && nums[i] == 0 {
			res = append(res, product)
		} else {
			res = append(res, product/nums[i])
		}

	}
	return res

}
