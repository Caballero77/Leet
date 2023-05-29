package main

import (
	"fmt"
	"sort"
)

func inner(candidates []int, index int, current []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	result := [][]int{}

	for i := index; i < len(candidates); i++ {
		val := candidates[i]
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		curr := make([]int, len(current)+1)
		copy(curr, current)
		curr[len(current)] = val
		if target == val {
			result = append(result, curr)
		} else if target > val {
			x := inner(candidates, i+1, curr, target-val)
			result = append(result, x...)
		} else {
			return result
		}
	}
	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return inner(candidates, 0, []int{}, target)
}

func main() {
	fmt.Println(combinationSum2([]int{1, 1, 1, 1, 2, 2, 2, 2}, 4))
}
