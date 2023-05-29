	package main

	import (
		"fmt"
	)

	func subsets(nums []int) [][]int {
		result := make([][]int, 0)
		result = append(result, []int{})
		for i := 0; i < len(nums); i++ {
			l := len(result)
			for j := 0; j < l; j++ {
				curr := make([]int, len(result[j])+1)
				copy(curr, result[j])
				curr[len(result[j])] = nums[i]
				result = append(result, curr)
			}
		}
		return result
	}

	func main() {
		fmt.Println(subsets([]int{1, 2, 3, 4, 5, 6}))
	}
