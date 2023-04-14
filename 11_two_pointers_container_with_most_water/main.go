package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0

	for left < right {
		area := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}
