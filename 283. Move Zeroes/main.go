package main

import "fmt"

func moveZeroes(nums []int) {
	x := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[x] = nums[i]
			if i != x {
				nums[i] = 0
			}
			x++
		} else if i >= len(nums)-x {
			nums[i] = 0
		}
	}
}

func main() {
	arr := []int{1}
	moveZeroes(arr)
	fmt.Println(arr)
}
