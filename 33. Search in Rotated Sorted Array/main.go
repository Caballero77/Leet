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
		}

		if nums[start] <= nums[middle] {
			if target > nums[middle] || target < nums[start] {
				start = middle + 1
			} else {
				end = middle - 1
			}
		} else {
			if target < nums[middle] || target > nums[end] {
				end = middle - 1
			} else {
				start = middle + 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
