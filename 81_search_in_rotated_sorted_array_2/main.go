package main

import "fmt"

// There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
// Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].
// Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.
// You must decrease the overall operation steps as much as possible.
//
// Example 1:
//
// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output: true

// Example 2:
//
// Input: nums = [2,5,6,0,0,1,2], target = 3
// Output: false

func main() {
	fmt.Println(search([]int{1, 1, 1, 1, 3, 1}, 3))
}

// Не мое решение, но и это тоже не бинарный поиск ни хе ра
func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1
	for left <= right {
		idx := left + (right-left)/2
		if nums[idx] == target {
			return true
		} else if nums[left] < nums[idx] {
			if nums[left] > target || target > nums[idx] {
				left = idx + 1
			} else {
				right = idx - 1
			}
		} else if nums[left] > nums[idx] {
			if target < nums[idx] || target > nums[right] {
				right = idx - 1
			} else {
				left = idx + 1
			}
		} else {
			left++
		}
	}
	return false
}

// Крайне странное решение, не бинарный поиск, но и хер знает как тут бинарно искать если честно
func search(nums []int, target int) bool {
	size := len(nums)
	if size == 0 {
		return false
	}
	left, right := 0, size-1

	if nums[left] == target || nums[right] == target {
		return true
	}
	if size < 3 {
		return false
	}
	idx := left + (right-left)/2
	if nums[idx] == target {
		return true
	}

	return search(nums[left+1:idx], target) || search(nums[idx:right], target)
}
