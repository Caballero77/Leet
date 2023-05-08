package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	len := len(prices)
	for i := 1; i < len; i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{1, 1, 0, 1, 1, 0, 1}))
}
