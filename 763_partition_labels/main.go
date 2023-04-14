package main

import "fmt"

func main() {
	// fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eaaaabaaec"))
}

func partitionLabels(s string) []int {
	last := map[rune]int{}
	for i, ss := range s {
		last[ss] = i
	}

	start, end := 0, 0
	var res []int
	for i, ss := range s {
		if last[ss] > end {
			end = last[ss]
		}

		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}

	return res
}
