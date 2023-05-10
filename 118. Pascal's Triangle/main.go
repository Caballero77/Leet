package main

import (
	"fmt"
)

func nextRow(row []int) []int {
	len := len(row)
	result := make([]int, len+1)
	result[0] = 1
	result[len] = 1
	for i := 1; i < len; i++ {
		result[i] = row[i-1] + row[i]
	}
	return result
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		result[i] = nextRow(result[i-1])
	}
	return result
}

func main() {
	res := generate(5)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println()
	}
}
