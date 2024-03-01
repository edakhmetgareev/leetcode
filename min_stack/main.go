package main

func main() {
	//	["MinStack",
	minStack := Constructor()

	//	"push","push","push",
	// [2147483646],[2147483646],[2147483647]
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)

	//
	//	"top","pop",
	r := minStack.Top() // get the last = 2
	minStack.Pop()
	//
	//	"getMin","pop","getMin","pop",
	r = minStack.GetMin() // return 1
	minStack.Pop()        // remove 2
	r = minStack.GetMin() // return 1
	minStack.Pop()        // remove 2

	//	"push","top",
	minStack.Push(2147483647)
	minStack.Top()

	//	"getMin","push",
	r = minStack.GetMin() // return 1
	minStack.Push(-2147483648)

	//
	//	"top","getMin","pop","getMin"]
	//  [,[],[-2147483648],[],[],[],[]]

	// remove 2
	_ = r
	minStack.Top()        // remove 2
	r = minStack.GetMin() // return 1
	minStack.Pop()        // remove 2
	minStack.GetMin()
}

type MinStack struct {
	stack    []int
	stackLen int
	minIdx   int
	minVal   int
}

func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0),
		minIdx: -1,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.stackLen++
	if val < this.minVal || this.minIdx < 0 {
		this.minVal = val
		this.minIdx = this.stackLen - 1
	}
}

func (this *MinStack) Pop() {
	val := this.stack[this.stackLen-1]

	this.stack = this.stack[:this.stackLen-1]
	this.stackLen--
	if val == this.minVal {
		minVal, key := this.findMin()
		this.minVal = minVal
		this.minIdx = key
	}
}

func (this *MinStack) Top() int {
	return this.stack[this.stackLen-1]
}

func (this *MinStack) GetMin() int {
	return this.minVal
}

func (this *MinStack) findMin() (int, int) {
	if this.stackLen == 0 {
		return 0, -1
	}

	if this.stackLen == 1 {
		return this.stack[0], 0
	}

	minVal := this.stack[0]
	idx := 0
	for i := 1; i < this.stackLen; i++ {
		if this.stack[i] < minVal {
			minVal = this.stack[i]
			idx = i
		}
	}

	return minVal, idx
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
