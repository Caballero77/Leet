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

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0

	var inner func(root *TreeNode) int

	inner = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		left := inner(root.Left) + 1
		right := inner(root.Right) + 1

		result = max(result, left+right)

		return max(right, left)
	}

	inner(root)

	return result
}

func main() {
	fmt.Println(diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
}
