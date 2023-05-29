package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	result = append(result, []int{})
	x := 1
	for i := 0; i < len(nums); i++ {
		l := len(result)
		start := 0
		if i > 0 && nums[i-1] == nums[i] {
			start = l - x
		}
		for j := start; j < l; j++ {
			curr := make([]int, len(result[j])+1)
			copy(curr, result[j])
			curr[len(result[j])] = nums[i]
			result = append(result, curr)
		}
		x = l - start
	}
	return result
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2, 2}))
}
