package main

import (
	"fmt"
)

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

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func isBalanced(root *TreeNode) bool {
	balanced, _ := isBalancedInner(root, 1)

	return balanced
}

func isBalancedInner(root *TreeNode, depth int) (bool, int) {
	if root == nil {
		return true, depth - 1
	}
	balanced, left := isBalancedInner(root.Left, depth+1)

	if !balanced {
		return balanced, -1
	}

	balanced, right := isBalancedInner(root.Right, depth+1)

	if !balanced {
		return balanced, -1
	}

	return abs(left-right) <= 1, max(left, right)
}

func main() {
	fmt.Println(isBalanced(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				}},
			Right: &TreeNode{
				Val: 2,
			}}))
}
