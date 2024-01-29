package main

import (
	"fmt"
)

// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
//
//
// Example 1:
//
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
// Example 2:
//
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
//
//
// Constraints:
//
// 1 <= n <= 45

func main() {
	result := climbStairs(8)
	fmt.Printf("%+v \n", result)
}

// 6
// 2 + 2 + 2
// 2 + 2 + 1 + 1
// 2 + 1 + 1 + 1 + 1
// 1 + 1 + 1 + 1 + 1 + 1

// 7
// 2 + 2 + 2 + 1
// 2 + 2 + 1 + 1 + 1
// 2 + 1 + 1 + 1 + 1 + 1
// 1 + 1 + 1 + 1 + 1 + 1 + 1

// 8
// 2 + 2 + 2 + 2
// 2 + 2 + 2  + 1 + 1
// 2 + 2 + 1 + 1 + 1 + 1
// 2 + 1 + 1 + 1 + 1 + 1 + 1
// 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1

// 9
// 2 + 2 + 2 + 2 + 1
// 2 + 2 + 2 + 1 + 1 + 1
// 2 + 2 + 1 + 1 + 1 + 1 + 1
// 2 + 1 + 1 + 1 + 1 + 1 + 1 + 1
// 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1

// 10
// 2 + 2 + 2 + 2 + 2
// 2 + 2 + 2 + 2 + 1 + 1
// 2 + 2 + 2 + 1 + 1 + 1 + 1
// 2 + 2 + 1 + 1 + 1 + 1 + 1 + 1
// 2 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1
// 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1

// 2
// 2
// 1 + 1

// 3
// 2 + 1
// 1 + 2
// 1 + 1 + 1

// 4
// 2 + 2
// 2 + 1 + 1
// 1 + 2 + 1
// 1 + 1 + 2
// 1 + 1 + 1 + 1

// это блять числа фибоначчи
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
