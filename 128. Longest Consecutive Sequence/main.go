package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if _, contains := set[nums[i]-1]; contains {
			continue
		}

		curr := 1
		_, contains := set[nums[i]+curr]
		for contains {
			curr++
			_, contains = set[nums[i]+curr]
		}

		if curr > max {
			max = curr
		}
	}

	return max
}

func main() {
	fmt.Println(longestConsecutive([]int{}))
}
