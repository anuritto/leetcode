package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	indexes := []int{}
	values := []int{}

	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(values) == 0 {
			indexes = append(indexes, i)
			values = append(values, temperatures[i])
			res[i] = 0
			continue
		}

		for {
			if temperatures[i] < values[len(values)-1] {
				res[i] = indexes[len(indexes)-1] - i
				values = append(values, temperatures[i])
				indexes = append(indexes, i)
				break
			}

			values = values[:len(values)-1]
			indexes = indexes[:len(indexes)-1]

			if len(values) == 0 {
				res[i] = 0
				values = append(values, temperatures[i])
				indexes = append(indexes, i)
				break
			}
		}
	}
	return res

}
