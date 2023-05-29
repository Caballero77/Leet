package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func goodNodes(root *TreeNode) int {
	result := 0
	var inner func(*TreeNode, int)
	inner = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		newMax := max
		if root.Val >= max {
			result++
			newMax = root.Val
		}
		inner(root.Left, newMax)
		inner(root.Right, newMax)
	}
	inner(root, -10001)
	return result
}

func main() {
	fmt.Println(1)
}
