package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	var inner func(*TreeNode)
	inner = func(root *TreeNode) {
		if root == nil {
			return
		}
		inner(root.Left)
		result = append(result, root.Val)
		inner(root.Right)
	}
	inner(root)
	return result
}

func main() {
	fmt.Println(1)
}
