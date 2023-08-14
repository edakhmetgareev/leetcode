package main

import "fmt"

func main() {
	result := removeDuplicates([]int{1, 1, 2, 3, 4, 5, 5, 5, 5, 5, 5})
	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
