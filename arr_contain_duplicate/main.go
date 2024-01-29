package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result := containsDuplicate(arr)
	fmt.Printf("result: %+v", result)
}

func containsDuplicate(nums []int) bool {
	c := make(map[int]struct{})

	for _, i := range nums {
		if _, ok := c[i]; ok {
			return true
		}

		c[i] = struct{}{}
	}

	return false
}
