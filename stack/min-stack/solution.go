package minstack

type minValue struct {
	index int
	prev  *minValue
}

type MinStack struct {
	stack []int
	min   *minValue
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
	}
}

func (this *MinStack) Push(val int) {

	if this.min == nil {
		this.min = &minValue{
			prev:  nil,
			index: 0,
		}
	} else if val < this.stack[this.min.index] {
		newMin := &minValue{
			prev:  this.min,
			index: len(this.stack),
		}
		this.min = newMin
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	n := len(this.stack)
	if n == 0 {
		return
	}

	if this.min.index == n-1 {
		this.min = this.min.prev
	}

	this.stack = this.stack[:n-1]
}

func (this *MinStack) Top() int {
	n := len(this.stack)
	if n == 0 {
		return 0
	}
	return this.stack[n-1]
}

func (this *MinStack) GetMin() int {
	if this.min == nil {
		return 0
	}
	return this.stack[this.min.index]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
