package evaluatereversepolishnotation

import "strconv"

func evalRPN(tokens []string) int {

	stk := []int{}
	for _, str := range tokens {
		if num, err := strconv.Atoi(str); err == nil {
			stk = append(stk, num)
			continue
		}

		n := len(stk)
		left := stk[n-2]
		right := stk[n-1]
		stk = stk[:n-2]

		switch str {
		case "+":
			stk = append(stk, left+right)
			break
		case "-":
			stk = append(stk, left-right)
			break
		case "*":
			stk = append(stk, left*right)
			break
		case "/":
			stk = append(stk, left/right)
			break
		}
	}

	return stk[0]

}
