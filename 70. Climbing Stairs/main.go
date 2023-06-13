package main

import (
	"fmt"
)

var memory map[int]int = map[int]int{
	1: 1,
	2: 2,
}

func climbStairs(n int) int {
	if value, ok := memory[n]; ok {
		return value
	}
	prev := 0
	if value, ok := memory[n-1]; ok {
		prev = value
	} else {
		prev = climbStairs(n - 1)
	}
	prevPrev := 0
	if value, ok := memory[n-2]; ok {
		prevPrev = value
	} else {
		prevPrev = climbStairs(n - 2)
	}
	curr := prev + prevPrev
	memory[n] = curr
	return curr
}

func main() {
	fmt.Println(climbStairs(45))
}
