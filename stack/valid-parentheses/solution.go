package validparentheses

var bracketsPairs = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(str string) bool {
	if len(str) == 0 || len(str)%2 == 1 {
		return false
	}

	stk := []rune(str)

	for _, br := range str {
		if _, isOpening := bracketsPairs[br]; isOpening {
			stk = append(stk, br)

			continue
		}

		n := len(stk)
		if n == 0 {
			return false
		}

		lst := stk[n-1]
		stk = stk[:n-1]

		if closingBr := bracketsPairs[lst]; br != closingBr {
			return false
		}

	}

	return true
}
