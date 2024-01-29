package main

import (
	"fmt"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
//
//
//
// Example 1:
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
//
// Input: nums = [3,3], target = 6
// Output: [0,1]

func main() {
	arr := []int{3, 2, 4}
	result := twoSumV2(arr, 6)
	fmt.Printf("%+v", result)
}

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSumV2(nums []int, target int) []int {
	c := make(map[int]int)
	for i, num := range nums {
		if v, ok := c[target-num]; ok {
			return []int{v, i}
		}

		c[num] = i
	}

	return nil
}
