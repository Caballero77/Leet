package main

import (
	"fmt"
	"sort"
)

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	first_two := prices[0] + prices[1]
	if money >= first_two {
		return money - first_two
	}
	return money
}

func main() {
	fmt.Println(buyChoco([]int{99, 1}, 100))
}
