package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	p := maxProfit(prices)
	fmt.Printf("max profit: %d\n", p)
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var maxProfit int
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			maxProfit += profit
		}
	}

	return maxProfit
}
