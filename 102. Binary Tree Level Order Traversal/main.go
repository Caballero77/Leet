package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	var inner func(root *TreeNode, depth int)
	inner = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(result) == depth {
			result = append(result, []int{root.Val})
		} else {
			result[depth] = append(result[depth], root.Val)
		}

		inner(root.Left, depth+1)
		inner(root.Right, depth+1)
	}

	inner(root, 0)
	return result
}

func main() {
	fmt.Println(1)
}
