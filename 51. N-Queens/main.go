package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	columns, westDiag, eastDiag := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	currMap := make([]string, 0)
	var inner func(int)
	inner = func(row int) {
		if row == n {
			cp := make([]string, n)
			copy(cp, currMap)
			result = append(result, cp)
		}

		for column := 0; column < n; column++ {
			if columns[column] || westDiag[column+row] || eastDiag[n+column-row-1] {
				continue
			}
			columns[column], westDiag[column+row], eastDiag[n+column-row-1] = true, true, true

			rowMap := ""
			for i := 0; i < n; i++ {
				if i == column {
					rowMap += "Q"
				} else {
					rowMap += "."
				}
			}
			currMap = append(currMap, rowMap)
			inner(row + 1)
			currMap = currMap[:len(currMap)-1]
			columns[column], westDiag[column+row], eastDiag[n+column-row-1] = false, false, false
		}
	}
	inner(0)
	return result
}

func main() {
	fmt.Println(solveNQueens(4))
}
