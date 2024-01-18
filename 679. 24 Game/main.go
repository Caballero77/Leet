package main

import (
	"fmt"
	"math"
)

func judgePoint24(cards []int) bool {
	floatingCards := make([]float64, len(cards))
	for i, v := range cards {
		floatingCards[i] = float64(v)
	}
	return judgePoint(floatingCards, 24)
}

func judgePoint(cards []float64, number float64) bool {
	if len(cards) == 1 {
		return math.Abs(cards[0]-number) < 0.0111
	}
	for i, v := range cards {
		cp := make([]float64, len(cards))
		copy(cp, cards)
		newCards := append(cp[:i], cp[i+1:]...)
		if judgePoint(newCards, number-v) {
			return true
		}
		if judgePoint(newCards, number+v) {
			return true
		}
		if judgePoint(newCards, number*v) {
			return true
		}
		if judgePoint(newCards, number/v) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(judgePoint24([]int{3, 3, 8, 8}))
}
