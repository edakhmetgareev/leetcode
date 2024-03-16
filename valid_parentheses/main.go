package main

import "fmt"

func main() {
	fmt.Println(isValid("{[(()))]}"))
}

func isValid(str string) bool {
	if len(str) == 0 || len(str)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]rune, 0, len(str)/2+1)

	for _, s := range []rune(str) {
		if val, ok := pairs[s]; ok {
			stackLen := len(stack)
			if stackLen == 0 || val != stack[stackLen-1] {
				return false
			}

			stack = stack[:stackLen-1]
		} else {
			stack = append(stack, s)
		}
	}

	return len(stack) == 0
}
