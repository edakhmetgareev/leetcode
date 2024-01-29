package main

import (
	"fmt"
	"sort"
)

// Дан массив целых чисел nums и целое число k. Нужно написать функцию, которая вынимает из массива k наиболее часто встречающихся элементов.
//
// Пример
// # ввод
// nums = [1,1,1,2,2,3]
// k = 2
// # вывод (в любом порядке)
// [1, 2]

func main() {
	// ввод
	nums := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 4, 4}
	k := 2
	// вывод (в любом порядке)
	// [1, 2]
	res := topKFrequent(nums, k)

	fmt.Printf("result: %+v", res)
}

func topKFrequent(nums []int, k int) []int {
	c := make(map[int]int, len(nums))
	for _, v := range nums {
		c[v]++
	}

	uniqNums := make([]int, 0, len(c))
	for v := range c {
		uniqNums = append(uniqNums, v)
	}

	sort.Slice(uniqNums, func(x, y int) bool {
		return c[uniqNums[x]] > c[uniqNums[y]]
	})

	return uniqNums[:k]
}
