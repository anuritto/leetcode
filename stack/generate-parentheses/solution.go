package generateparentheses

type Combination struct {
	str        string
	openCount  int
	closeCount int
}

func generateParenthesis(n int) []string {

	stk := []Combination{Combination{}}

	for {
		width := len(stk)

		for i := 0; i < width; i++ {
			item := stk[i]
			if item.openCount < n {
				stk = append(stk, Combination{
					str:        item.str + "(",
					openCount:  item.openCount + 1,
					closeCount: item.closeCount,
				})
			}

			if item.closeCount < item.openCount {
				stk = append(stk, Combination{
					str:        item.str + ")",
					openCount:  item.openCount,
					closeCount: item.closeCount + 1,
				})
			}
		}
		// not new combination
		if len(stk) == width {
			break
		}

		stk = stk[width:]
	}
	res := make([]string, 0)
	for _, comb := range stk {
		res = append(res, comb.str)
	}

	return res
}
