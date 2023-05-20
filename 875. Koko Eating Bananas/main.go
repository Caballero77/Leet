package main

import (
	"fmt"
)

func canEat(piles []int, h int, k int) bool {
	time := 0
	for i := 0; i < len(piles); i++ {
		time += (piles[i] + k - 1) / k
		if time > h {
			return false
		}
	}
	return true
}

func minEatingSpeed(piles []int, h int) int {
	start, end := 1, 1000000000
	result := end / 2
	for start <= end {
		middle := (start + end) / 2
		if canEat(piles, h, middle) {
			result = middle
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return result
}

func main() {
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}
