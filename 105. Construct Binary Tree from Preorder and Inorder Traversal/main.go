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

func getIndex(array []int, value int) int {
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return i
		}
	}
	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	curr := &TreeNode{Val: preorder[0]}
	middle := getIndex(inorder, curr.Val)
	curr.Left = buildTree(preorder[1:middle+1], inorder[:middle])
	curr.Right = buildTree(preorder[middle+1:], inorder[middle+1:])
	return curr
}

func main() {
	fmt.Println(1)
}
