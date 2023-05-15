package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			curr := numbers[i] + numbers[j]
			if curr == target {
				return []int{i + 1, j + 1}
			}
			if curr > target {
				break
			}
		}
	}
	return []int{-1, -1}
}

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("%d %d\n", result[0], result[1])
}
