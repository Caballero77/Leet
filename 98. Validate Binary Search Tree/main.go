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

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode = nil
	var inner func(*TreeNode) bool
	inner = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		left := inner(root.Left)
		if !left {
			return left
		}
		if prev != nil && root.Val <= prev.Val {
			return false
		}
		prev = root
		right := inner(root.Right)
		if !right {
			return right
		}
		return true
	}

	return inner(root)
}

func main() {
	fmt.Println(1)
}
