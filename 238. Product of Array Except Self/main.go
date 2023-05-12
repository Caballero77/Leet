package main

import (
	"fmt"
)

func productExceptSelfRec(nums *[]int, i int, n int) int {
	if i == len(*nums)-1 {
		result := (*nums)[i]
		(*nums)[i] = n
		return result
	}
	x := productExceptSelfRec(nums, i+1, n*(*nums)[i])
	result := (*nums)[i] * x
	(*nums)[i] = n * x
	return result
}

func productExceptSelf(nums []int) []int {
	productExceptSelfRec(&nums, 0, 1)
	return nums
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 0, 3, 4}))
}
