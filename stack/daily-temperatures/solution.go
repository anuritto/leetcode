package dailytemperatures

import (
	"errors"
)

type dayInfo struct {
	index int
	val   int
}

type dayInfoStack struct {
	stk []dayInfo
}

func (this *dayInfoStack) push(day dayInfo) {
	this.stk = append(this.stk, day)
}

func (this *dayInfoStack) pop() (dayInfo, error) {
	n := len(this.stk)
	if n == 0 {
		return dayInfo{}, errors.New("empty stack")
	}

	this.stk = this.stk[:n-1]

	return this.stk[n-1], nil
}

func (this *dayInfoStack) getLast() (dayInfo, error) {
	n := len(this.stk)
	if n == 0 {
		return dayInfo{}, errors.New("empty stack")
	}

	return this.stk[n-1], nil
}

func (this *dayInfoStack) delLast() error {
	n := len(this.stk)
	if n == 0 {
		return errors.New("empty stack")
	}

	this.stk = this.stk[:n-1]

	return nil
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stk := dayInfoStack{
		stk: []dayInfo{},
	}

	for i := len(temperatures) - 1; i >= 0; i-- {
		lastDay, notFound := stk.getLast()
		currentDay := dayInfo{val: temperatures[i], index: i}

		if notFound != nil {
			stk.push(currentDay)
			res[i] = 0
		}
		var err error
		for {
			if lastDay.val > currentDay.val {
				res[i] = lastDay.index - currentDay.index
				stk.push(currentDay)
				break
			}

			stk.delLast()

			lastDay, err = stk.getLast()
			if err != nil {
				stk.push(currentDay)
				res[i] = 0
				break
			}
		}
	}
	return res

}
