package main

import (
	"container/heap"
	"fmt"
)

// Дан массив чисел nums и некоторое неизвестное науке число k. Нужно написать функцию, которая вынимает k самых больших чисел из массива nums.
//
// Пример
// // ввод
// nums := []int{100, 50, 0, 150, 100, 0, -30, 70}
// k := 3
// // ожидаемый вывод (в любом порядке)
// // 100 150 100

func main() {
	// ввод
	// nums := []int{100, 50, 0, 150, 100, 0, -30, 70}
	nums := []int{100, 50, 2, 1, 3, 49, 51, 99, 101, 89, -30, 70, 500, 101}
	k := 3
	// ожидаемый вывод (в любом порядке)
	// 100 150 100
	h := MinHeap(nums[:k])
	heap.Init(&h)

	for i := k; i < len(nums); i++ {
		n := nums[i]
		if n > h[0] {
			a := heap.Pop(&h)
			fmt.Printf("i:%d, val:%+v", i, a)
			heap.Push(&h, nums[i])
		}
	}

	fmt.Printf("result: %v", h)
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
