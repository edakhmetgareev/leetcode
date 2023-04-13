package main

import (
	"errors"
	"fmt"
	"strings"
)

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
//
// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

func main() {
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	openBrackets := map[string]struct{}{
		"[": {},
		"{": {},
		"(": {},
	}
	pairs := map[string]string{
		"]": "[",
		"}": "{",
		")": "(",
	}
	strSlice := strings.Split(s, "")
	stack := Stack{}
	for _, ss := range strSlice {
		if _, ok := openBrackets[ss]; ok {
			stack.Push(ss)
		} else {
			theLast := stack.Last()
			if val, ok := pairs[ss]; ok && val == theLast {
				err := stack.Pop()
				if err != nil {
					return false
				}
			} else {
				return false
			}
		}
	}

	return stack.Size() == 0
}

type Stack struct {
	items []string
}

func (s *Stack) Size() int {
	return len(s.items)
}
func (s *Stack) Pop() error {
	if s.Size() < 1 {
		return errors.New("out of range")
	}
	if s.Size() == 1 {
		s.items = nil
		return nil
	}

	s.items = s.items[0 : len(s.items)-1]

	return nil
}

func (s *Stack) Last() string {
	if s.Size() == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}
