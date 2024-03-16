package main

import (
	"fmt"
)

func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-2, -1}
	res := maxSubArray(nums)
	fmt.Printf("result: %d", res)
}

// Алгоритм Кадана. Движемся по массиву слева направо и накапливаем в переменной s текущую частичную сумму.
// Если в какой-то момент s окажется отрицательной, то присвоим s = 0. Максимум из всех значений переменной s за время
// прохода по массиву и будет ответом на задачу.
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var sum int
	maxValue := nums[0]
	if maxValue > 0 {
		sum = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		sum += nums[i]

		if sum > maxValue {
			maxValue = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return maxValue
}
