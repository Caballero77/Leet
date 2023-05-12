package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func containsDuplicate(nums []int) bool {
	dict := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, containt := dict[nums[i]]; containt {
			return true
		}
		dict[nums[i]] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
}
