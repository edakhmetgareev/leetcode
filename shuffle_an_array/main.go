package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	param_1 := obj.Reset()
	param_2 := obj.Shuffle()
	fmt.Printf("param_1 = %v, param_2 = %v\n", param_1, param_2)
}

type Solution struct {
	orig     []int
	shuffled []int
}

func Constructor(nums []int) Solution {
	shuf := make([]int, len(nums))
	copy(shuf, nums)
	return Solution{orig: nums, shuffled: shuf}
}

func (this *Solution) Reset() []int {
	return this.orig
}

func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.orig), func(i, j int) {
		this.shuffled[i], this.shuffled[j] = this.shuffled[j], this.shuffled[i]
	})

	return this.shuffled
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
