package main

import (
	"fmt"
)

func orangesRotting(grid [][]int) int {
	fresh := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	minutes := 0
	newRotten := make([][]int, 0)
	rotten := 0
	for true {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 1 {
					up := i-1 >= 0 && grid[i-1][j] == 2
					down := i+1 < len(grid) && grid[i+1][j] == 2
					left := j-1 >= 0 && grid[i][j-1] == 2
					right := j+1 < len(grid[i]) && grid[i][j+1] == 2
					if up || down || left || right {
						newRotten = append(newRotten, []int{i, j})
					}
				}
			}
		}
		if len(newRotten) != 0 {
			for i := 0; i < len(newRotten); i++ {
				grid[newRotten[i][0]][newRotten[i][1]] = 2
			}
			rotten += len(newRotten)
			minutes++
			newRotten = make([][]int, 0)
			continue
		}
		break
	}

	if fresh == rotten {
		return minutes
	}
	return -1
}

func main() {
	grid := [][]int{
		{2, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1}}
	fmt.Println(orangesRotting(grid))
}
