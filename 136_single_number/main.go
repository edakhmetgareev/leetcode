package main

import "fmt"

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.
// Example 1:
// Input: nums = [2,2,1]
// Output: 1

// Example 2:
// Input: nums = [4,1,2,1,2]
// Output: 4

// Example 3:
// Input: nums = [1]
// Output: 1
//
// Constraints:
// 1 <= nums.length <= 3 * 104
// -3 * 104 <= nums[i] <= 3 * 104
// Each element in the array appears twice except for one element which appears only once.

func main() {
	fmt.Println("Result:", singleNumberV1([]int{2, 2, 1}))
}

func singleNumberV1(nums []int) int {
	hash := make(map[int]int, len(nums)/2+1)
	for idx, num := range nums {
		_, ok := hash[num]
		if !ok {
			hash[num] = idx
			continue
		}

		delete(hash, num)
	}

	for v := range hash {
		return v
	}

	return 0
}

func singleNumberV2(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}

	return result
}
