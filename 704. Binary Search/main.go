package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
