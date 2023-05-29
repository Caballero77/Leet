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

func rightSideView(root *TreeNode) []int {
	result := []int{}
	var inner func(*TreeNode, int)
	inner = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(result) == depth {
			result = append(result, root.Val)
		}
		inner(root.Right, depth+1)
		inner(root.Left, depth+1)
	}

	inner(root, 0)
	return result
}

func main() {
	fmt.Println(1)
}
