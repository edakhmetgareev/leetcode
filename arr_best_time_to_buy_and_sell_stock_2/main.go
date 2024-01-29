package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	p := maxProfit(prices)
	fmt.Printf("max profit: %d\n", p)
}

func maxProfit(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
