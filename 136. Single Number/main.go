package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func main() {
	fmt.Println(singleNumber([]int{-1, 2, 2, 5, 3, 5, 3}))
}
