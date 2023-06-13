package main

import (
	"fmt"
)

func calculateOcean(up, down, left, right byte) byte {
	p := false
	a := false
	for _, ocean := range []byte{up, down, left, right} {
		if ocean == 'a' {
			a = true
		}
		if ocean == 'p' {
			p = true
		}
		if ocean == 'b' {
			a = true
			p = true
		}
	}

	if p && a {
		return 'b'
	} else if a {
		return 'a'
	} else if p {
		return 'p'
	}
	return 0
}

func pacificAtlantic(heights [][]int) [][]int {
	visited := make([][]byte, len(heights), len(heights[0]))
	for i := 0; i < len(heights); i++ {
		visited[i] = make([]byte, len(heights[i]))
		for j := 0; j < len(heights[i]); j++ {
			visited[i][j] = 0
		}
	}

	var inner func(int, int, int) byte

	inner = func(x, y, prev int) byte {
		if x < 0 || y < 0 {
			return 'p'
		}
		if x >= len(heights) || y >= len(heights[0]) {
			return 'a'
		}
		if heights[x][y] > prev || visited[x][y] == 'v' {
			return 0
		}
		if visited[x][y] != 0 {
			return visited[x][y]
		}

		visited[x][y] = 'v'

		right := inner(x, y+1, heights[x][y])
		down := inner(x+1, y, heights[x][y])
		up := inner(x-1, y, heights[x][y])
		left := inner(x, y-1, heights[x][y])

		visited[x][y] = calculateOcean(up, down, left, right)

		return visited[x][y]
	}

	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if visited[i][j] == 0 {
				inner(i, j, int(^uint(0)>>1))
			}
		}
	}

	result := make([][]int, 0)
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			if visited[i][j] == 'b' {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func main() {
	grid := [][]int{{2, 2, 2, 2, 2},
		{2, 1, 1, 1, 1},
		{2, 1, 1, 1, 1},
		{3, 4, 4, 1, 1},
		{2, 1, 1, 1, 1}}
	fmt.Println(pacificAtlantic(grid))
}
