package main

import "fmt"

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//
// You must implement a solution with a linear runtime complexity and use only constant extra space.
// Example 1:
//
// Input: nums = [2,2,1]
// Output: 1
// Example 2:
//
// Input: nums = [4,1,2,1,2]
// Output: 4
// Example 3:
//
// Input: nums = [1]
// Output: 1

func main() {
	arr := []int{1, 2, 2, 1}
	arr2 := []int{2}
	result := intersect(arr, arr2)
	fmt.Printf("%+v", result)
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		return intersect(nums2, nums1)
	}

	cache := make(map[int]int)
	var result []int

	for _, i := range nums1 {
		cache[i]++
	}

	for _, i := range nums2 {
		if cache[i] > 0 {
			result = append(result, i)
			cache[i]--
		}
	}

	return result
}
