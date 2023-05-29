package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {

	var inner func(root *TreeNode, depth int) int

	inner = func(root *TreeNode, depth int) int {
		if root == nil {
			return depth - 1
		}

		return max(inner(root.Left, depth+1), inner(root.Right, depth+1))
	}

	return inner(root, 1)
}

func main() {
	fmt.Println(1)
}
