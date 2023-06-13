package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	expexted := 0
	for i := 1; i <= len(nums); i++ {
		expexted += i
	}

	actual := 0
	for _, v := range nums {
		actual += v
	}

	return expexted - actual
}

func main() {
	fmt.Println(missingNumber([]int{0, 1}))
}
