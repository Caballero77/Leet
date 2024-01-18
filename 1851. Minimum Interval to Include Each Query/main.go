package main

import (
	"fmt"
	"sort"
)

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][1] - intervals[i][0] + 1) > (intervals[j][1] - intervals[j][0] + 1)
	})
	dict := make(map[int]int)
	for _, interval := range intervals {
		size := interval[1] - interval[0] + 1
		for i := interval[0]; i <= interval[1]; i++ {
			dict[i] = size
		}
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		v, ok := dict[query]
		if ok {
			result[i] = v
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	fmt.Println(minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}))
}
