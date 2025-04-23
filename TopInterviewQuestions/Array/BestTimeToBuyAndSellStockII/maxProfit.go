package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	minIndex := 0
	maxIndex := 0

	for {
		for i := minIndex + 1; i < len(prices) && prices[minIndex] > prices[i]; i++ {
			minIndex = i
		}

		maxIndex = minIndex
		for i := maxIndex + 1; i < len(prices) && prices[maxIndex] <= prices[i]; i++ {
			maxIndex = i
		}

		if minIndex >= maxIndex {
			return profit
		}

		profit += prices[maxIndex] - prices[minIndex]
		minIndex = maxIndex + 1
	}
}

func main() {
	arr := []int{2, 2, 5}

	profit := maxProfit(arr)
	fmt.Print("Profit: ", profit)
}
