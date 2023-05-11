package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	var maxDepth = 0
	var sum = 0

	var deepestLeavesSumInner func(root *TreeNode, depth int)

	deepestLeavesSumInner = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		deepestLeavesSumInner(root.Left, depth+1)

		if depth > maxDepth {
			maxDepth = depth
			sum = root.Val
		} else if depth == maxDepth {
			sum += root.Val
		}

		deepestLeavesSumInner(root.Right, depth+1)
	}
	deepestLeavesSumInner(root, 1)
	return sum
}

func main() {
	root := TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 8}}
	fmt.Println(deepestLeavesSum(&root))
}
