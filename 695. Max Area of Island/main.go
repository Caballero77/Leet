package main

import (
	"fmt"
)

func markVisited(grid *[][]int, x, y int) int {
	if (*grid)[x][y] != 1 {
		return 0
	}
	(*grid)[x][y] = 2
	area := 1

	possibleConected := [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}}
	for _, point := range possibleConected {
		if point[0] < 0 || point[0] >= len((*grid)) ||
			point[1] < 0 || point[1] >= len(((*grid)[0])) {
			continue
		}
		area += markVisited(grid, point[0], point[1])
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				result = max(result, markVisited(&grid, i, j))
			}
		}
	}
	return result
}

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(grid))
}
