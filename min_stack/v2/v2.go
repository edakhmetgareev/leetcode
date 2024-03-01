package main

import (
	"container/heap"
	"fmt"
)

func main() {
	minStack := Constructor()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(-3)
	r := minStack.GetMin()
	minStack.Pop()
	r = minStack.Top()
	r = minStack.GetMin()
	fmt.Println(r)
}

type MinStack struct {
	stack   []int
	minHeap MinHeap
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func Constructor() MinStack {
	h := MinHeap(make([]int, 0))
	heap.Init(&h)
	return MinStack{
		stack:   make([]int, 0),
		minHeap: h,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	heap.Push(&this.minHeap, val)
}

func (this *MinStack) Pop() {
	val := this.stack[this.minHeap.Len()-1]
	this.stack = this.stack[:this.minHeap.Len()-1]
	for i, v := range this.minHeap {
		if v == val {
			heap.Remove(&this.minHeap, i)
			return
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[this.minHeap.Len()-1]
}

func (this *MinStack) GetMin() int {
	return this.minHeap[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
