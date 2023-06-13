package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < amount + 1; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}
	if dp[amount] != amount+1 {
		return dp[amount]
	}
	return -1
}
 
func main() {
	fmt.Println(coinChange([]int{186,419,83,408}, 6249))
}
