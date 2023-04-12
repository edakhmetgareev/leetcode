package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2}))
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		idx := (left + right) / 2

		// in the right side
		if nums[idx] > nums[right] {
			left = idx + 1
		} else {
			right = idx
		}
	}

	return nums[left]
}
