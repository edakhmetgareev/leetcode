package main

import (
	"errors"
	"fmt"
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
	fmt.Println(isValid2("{[(]}"))
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	openBrackets := map[rune]struct{}{
		'[': {},
		'{': {},
		'(': {},
	}
	pairs := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := Stack{}
	for _, ss := range s {
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
	items []rune
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

func (s *Stack) Last() rune {
	if s.Size() == 0 {
		return rune(0)
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func isValid2(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	openBrackets := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}

	var st []rune
	for _, ss := range s {
		val, ok := openBrackets[ss]
		if ok {
			st = append(st, val)
		} else {
			if len(st) == 0 {
				return false
			}
			theLast := st[len(st)-1]
			if theLast != ss {
				return false
			}

			st = st[:len(st)-1]
		}
	}
	return len(st) == 0

}
