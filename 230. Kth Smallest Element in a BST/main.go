package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var number = 0

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}

	left := kthSmallest(root.Left, k)

	if left != -1 {
		return left
	}

	number++
	if number == k {
		number = 0
		return root.Val
	}

	return kthSmallest(root.Right, k)
}

func main() {
	root := TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 8}}
	fmt.Println(kthSmallest(&root, 2))
}


