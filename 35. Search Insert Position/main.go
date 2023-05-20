package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		middle := (start + end) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return start
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6, 8}, 0))
}
