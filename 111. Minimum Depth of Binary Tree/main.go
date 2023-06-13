package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := int(^uint(0) >> 1)
	var inner func(*TreeNode, int)
	inner = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			result = min(result, depth)
			return
		}

		inner(root.Left, depth+1)
		inner(root.Right, depth+1)
	}
	inner(root, 1)
	return result
}

func main() {
	fmt.Println(1)
}
