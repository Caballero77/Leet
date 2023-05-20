package main

import (
	"fmt"
)

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	min := nums[0]
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] >= nums[0] {
			start = middle + 1
		} else {
			if nums[middle] < min {
				min = nums[middle]
			}
			end = middle - 1
		}
	}
	return min
}

func main() {
	fmt.Println(findMin([]int{5, 1, 2, 3, 4}))
}
