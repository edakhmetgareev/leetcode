package main

import (
	"fmt"
)

func main() {
	fmt.Println(winnerOfGame("AAAAABBBBBBAAAAA"))
}

func winnerOfGame(colors string) bool {
	var aliceIndex, bobIndex int

	if len(colors) < 3 {
		return false
	}

	for i := 1; i < len(colors)-1; i++ {
		if colors[i-1] == colors[i] && colors[i] == colors[i+1] {
			if colors[i] == 'A' {
				aliceIndex++
			} else {
				bobIndex++
			}
		}
	}

	return aliceIndex > bobIndex
}
