package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	max := nums[0]

	if sum%k != 0 || max > sum/k {
		return false
	}

	mem := make(map[int]bool)

	var search func(int, int, int, int, int) bool
	search = func(visited int, target int, i int, sets int, set int) bool {
		if target == 0 {
			if sets == 1 {
				return true
			}
			v, ok := mem[set]
			if ok {
				return v
			}
			res := search(visited, sum/k, 0, sets-1, 1<<max)
			mem[set] = res
			return res
		}
		for j := i; j < len(nums); j++ {
			if 0 == (visited&(1<<j)) && nums[j] <= target {
				if search(visited|(1<<j), target-nums[j], i+1, sets, set|(1<<nums[j])) {
					return true
				}
			}
		}
		return false
	}

	return search(1<<len(nums), sum/k, 0, k, 1<<max)
}

func main() {
	fmt.Println(canPartitionKSubsets([]int{3, 3, 10, 2, 6, 5, 10, 6, 8, 3, 2, 1, 6, 10, 7, 2}, 6))
}
