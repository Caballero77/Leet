package main

import (
	"fmt"
)

func markVisited(grid *[][]byte, x, y int) {
	if (*grid)[x][y] != '1' {
		return
	}
	(*grid)[x][y] = '2'

	possibleConected := [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}}
	for _, point := range possibleConected {
		if point[0] < 0 || point[0] >= len((*grid)) ||
			point[1] < 0 || point[1] >= len(((*grid)[0])) {
			continue
		}
		markVisited(grid, point[0], point[1])
	}
}

func numIslands(grid [][]byte) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				markVisited(&grid, i, j)
				result++
			}
		}
	}
	return result
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '1', '1'},
		{'1', '1', '0', '0', '1'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
