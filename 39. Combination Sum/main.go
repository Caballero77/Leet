package main

import "fmt"

func combinationSumInner(candidates []int, target int, resultCandidate []int) [][]int {
	results := [][]int{}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] == target {
			result := make([]int, len(resultCandidate))
			copy(result, resultCandidate)
			results = append(results, append(result, candidates[i]))
		} else if candidates[i] < target {
			results = append(results, combinationSumInner(candidates[i:], target-candidates[i], append(resultCandidate, candidates[i]))...)
		}
	}

	return results
}

func combinationSum(candidates []int, target int) [][]int {
	return combinationSumInner(candidates, target, []int{})
}

func main() {
	result := combinationSum([]int{2, 3, 5, 7}, 13)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}
