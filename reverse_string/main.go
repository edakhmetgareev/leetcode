package main

import "fmt"

func main() {
	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)
	//for _, b := range s {
	//
	//}
	fmt.Printf("result: %s", s)
}

func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}
