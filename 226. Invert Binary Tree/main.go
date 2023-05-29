package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertTree(root.Left)
	invertTree(root.Right)

	buf := root.Left
	root.Left = root.Right
	root.Right = buf

	return root
}

func main() {
	fmt.Println(1)
}
