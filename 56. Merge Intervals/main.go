package main

import (
	"fmt"
	"sort"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals[:], func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	i, removed := 0, 0
	for i < len(intervals)-1-removed && len(intervals) > 1 {
		if intervals[i][1] >= intervals[i+1+removed][0] {
			intervals[i][1] = max(intervals[i][1], intervals[i+1+removed][1])
			intervals[i+1+removed] = nil
			removed++
		} else {
			i += removed + 1
			removed = 0
		}
	}
	result, j := make([][]int, len(intervals)), 0
	for i := 0; i < len(intervals); i++ {
		if intervals[i] != nil {
			result[j] = intervals[i]
			j++
		}
	}
	return result[:j]
}

func main() {
	result := merge([][]int{{-2, -1}, {-1, 3}, {1, 6}, {8, 10}, {15, 18}})
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}
