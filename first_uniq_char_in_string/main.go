package main

import (
	"fmt"
)

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
//
// Example 1:
//
// Input: s = "leetcode"
// Output: 0
// Example 2:
//
// Input: s = "loveleetcode"
// Output: 2
// Example 3:
//
// Input: s = "aabb"
// Output: -1
//
// Constraints:
//
// 1 <= s.length <= 105
// s consists of only lowercase English letters.
func main() {
	s := "leetcode"

	res := firstUniqChar(s)
	fmt.Printf("res: %d", res)
}

func firstUniqChar(s string) int {
	acc := make([]int, 26)
	for _, ss := range s {
		acc[ss-'a']++
	}

	for i, ss := range s {
		if acc[ss-'a'] == 1 {
			return i
		}
	}

	return -1
}
