package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	start, end := 0, len(height)-1
	startMax, endMax := height[start], height[end]
	result := 0

	for start < end {
		if startMax < endMax {
			start++
			startMax = max(startMax, height[start])
			result += startMax - height[start]
		} else {
			end--
			endMax = max(endMax, height[end])
			result += endMax - height[end]
		}
	}
	return result
}

func main() {
	fmt.Println(trap([]int{6, 7, 5, 4, 3, 2, 1, 2, 3, 4, 5, 6}))
}
