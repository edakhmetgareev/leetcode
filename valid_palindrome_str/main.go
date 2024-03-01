package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
// Example 1:
//
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:
//
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:
//
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.
func main() {
	s := "A man, a plan, a canal: Panama"

	res := isPalindrome(s)
	fmt.Printf("res: %v", res)
}

func isPalindrome(s string) bool {
	// how to filter version 1
	f := func(char rune) rune {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return -1
		}

		return unicode.ToLower(char)
	}

	s = strings.Map(f, s)

	s = strings.ToLower(s)
	p := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = p.ReplaceAllString(s, "")

	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
