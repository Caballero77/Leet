package main

import (
	"fmt"
)

func numTrees(n int) int {
	results := make([]int, n+1)
	results[0] = 1
	results[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			results[i] += results[j-1] * results[i-j]
		}
	}
	return results[n]
}

func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
}
