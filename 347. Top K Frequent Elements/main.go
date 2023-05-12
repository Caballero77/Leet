package main

import (
	"fmt"
)

func min(arr [][]int) (int, int) {
	index := 0
	m := arr[index][1]
	for i := 1; i < len(arr); i++ {
		if len(arr[i]) == 0 {
			return i, 0
		}
		if arr[i][1] < m {
			m = arr[i][1]
			index = i
		}
	}
	return index, m
}

func mapResult(arr [][]int) []int {
	result := make([]int, len(arr))
	for i, x := range arr {
		result[i] = x[0]
	}
	return result
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, v := range nums {
		if freq, contains := freqMap[v]; contains {
			freqMap[v] = freq + 1
		} else {
			freqMap[v] = 1
		}
	}
	result := make([][]int, k)
	minI, minV := 0, 0
	for value, freq := range freqMap {
		if freq > minV {
			result[minI] = []int{value, freq}
		}
		minI, minV = min(result)
	}
	return mapResult(result)
}

func main() {
	fmt.Println(topKFrequent([]int{-1, -2, 0, 1, 2, 1, 2, -1}, 3))
}
