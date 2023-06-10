package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {

	i, j := 0, 1
	numbersLen := len(numbers)
	for i < numbersLen-1 {
		for j < numbersLen {
			if (numbers[i] + numbers[j]) == target {
				return []int{i + 1, j + 1}
			}
			j++
		}
		i++
		j = i + 1
	}

	return []int{i + 1, j + 1}
}
