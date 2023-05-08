package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	current := prices[0]
	max := 0
	len := len(prices)
	for i := 1; i < len; i++ {
		if prices[i] < current {
			current = prices[i]
		} else {
			candidate := prices[i] - current
			if candidate > max {
				max = candidate
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{1, 1, 1, 1, 1, 0, 1}))
}
