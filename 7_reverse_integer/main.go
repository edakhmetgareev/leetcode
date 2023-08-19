package main

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
// Example 1:
//
// Input: x = 123
// Output: 321
// Example 2:
//
// Input: x = -123
// Output: -321
// Example 3:
//
// Input: x = 120
// Output: 21
//
// Constraints:
//
// -231 <= x <= 231 - 1
func main() {
	result := reverse(1234)

	println(result)
}

func reverse(x int) int {
	negative := x < 0
	if negative {
		x = -x
	}

	var result int
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	maxInt := 1<<31 - 1
	if result >= maxInt || result <= -maxInt {
		return 0
	}

	if negative {
		result = -result
	}

	return result
}
