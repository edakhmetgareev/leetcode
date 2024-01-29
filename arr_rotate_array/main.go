package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7} // 7, 6, 5, 4, 3, 2, 1
	rotate1(arr, 3)
	fmt.Printf("result 1: %+v", arr)

	arr2 := []int{1, 2, 3, 4, 5, 6, 7} // 7, 6, 5, 4, 3, 2, 1
	rotate2(arr2, 3)
	fmt.Printf("result 2: %+v", arr)
}

func rotate1(nums []int, k int) {
	l := len(nums)
	k %= l
	k = l - k

	if k < 1 {
		return
	}
	fmt.Println(k)
	v := append(nums[k:], nums[:k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = v[i]
	}
}

func rotate2(nums []int, k int) {
	l := len(nums)
	k %= l
	k = l - k

	if k < 1 {
		return
	}

	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
	reverse(nums, 0, l-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
