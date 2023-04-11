package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
}

func search(nums []int, target int) int {
	min, max := 0, len(nums)-1
	var idx int

	for min <= max {
		idx = min + (max-min)/2

		if nums[idx] == target {
			return idx
		}

		// Input: nums = [4,5,6,^7,0,1,2], target = 4
		if nums[min] > nums[idx] {
			if target > nums[idx] && target <= nums[max] {
				min = idx + 1
			} else {
				max = idx - 1
			}
		} else {
			if target >= nums[min] && target < nums[idx] {
				max = idx - 1
			} else {
				min = idx + 1
			}
		}
	}

	return -1
}
