package main

import "fmt"

// We are playing the Guess Game. The game is as follows:
//
// I pick a number from 1 to n. You have to guess which number I picked.
//
// Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.
//
// You call a pre-defined API int guess(int num), which returns three possible results:
//
// -1: Your guess is higher than the number I picked (i.e. num > pick).
// 1: Your guess is lower than the number I picked (i.e. num < pick).
// 0: your guess is equal to the number I picked (i.e. num == pick).
// Return the number that I picked.
//
//
//
// Example 1:
//
// Input: n = 10, pick = 6
// Output: 6
// Example 2:
//
// Input: n = 1, pick = 1
// Output: 1
// Example 3:
//
// Input: n = 2, pick = 1
// Output: 1

var myNum = 6

func main() {
	fmt.Println(guessNumber(10))
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	min := 0
	max := n

	for min <= max {
		val := (min + max) / 2
		if guess(val) == 0 {
			return val
		}

		if guess(val) == -1 {
			max = val - 1
			continue
		}

		if guess(val) == 1 {
			min = val + 1
			continue
		}
	}

	return 0
}

func guess(n int) int {
	if n > myNum {
		return -1
	}

	if n < myNum {
		return 1
	}

	return 0
}
