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

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		if matrix[i][0] <= target && matrix[i+1][0] > target {
			return search(matrix[i], target) != -1
		}
	}
	return search(matrix[len(matrix)-1], target) != -1
}

func main() {
	fmt.Println(searchMatrix([][]int{{1}, {3}}, 3))
}
