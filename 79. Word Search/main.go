package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func getPossiblePoint(board [][]byte, point Point) []Point {
	xLen := len(board)
	yLen := len(board[0])
	allMoves := []Point{{X: point.X, Y: point.Y - 1}, {X: point.X, Y: point.Y + 1}, {X: point.X - 1, Y: point.Y}, {X: point.X + 1, Y: point.Y}}
	for i := 0; i < len(allMoves); i++ {
		if allMoves[i].X < 0 || allMoves[i].X >= xLen ||
			allMoves[i].Y < 0 || allMoves[i].Y >= yLen || board[allMoves[i].X][allMoves[i].Y] == '@' {
			allMoves = append(allMoves[:i], allMoves[i+1:]...)
			i--
		}
	}
	return allMoves
}

func getNext(board [][]byte, point Point) (bool, Point) {
	if point.Y == len(board[0])-1 {
		if point.X == len(board)-1 {
			return false, point
		}
		point.Y = 0
		point.X++
	} else {
		point.Y++
	}
	return true, point
}

func inner(board [][]byte, currentPoint Point, currentWord int, target string) bool {
	if board[currentPoint.X][currentPoint.Y] != target[currentWord] {
		return false
	}
	if currentWord == len(target)-1 {
		return true
	}

	buf := board[currentPoint.X][currentPoint.Y]
	board[currentPoint.X][currentPoint.Y] = '@'

	moves := getPossiblePoint(board, currentPoint)
	currentWord++
	for _, move := range moves {
		if board[move.X][move.Y] == target[currentWord] {
			if inner(board, move, currentWord, target) {
				return true
			}
		}
	}
	board[currentPoint.X][currentPoint.Y] = buf

	return false
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if inner(board, Point{X: i, Y: j}, 0, word) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A'},
	}
	word := "B"
	fmt.Println(exist(board, word))
}
