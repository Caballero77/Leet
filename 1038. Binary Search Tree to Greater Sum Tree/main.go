package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var iter func(head *TreeNode)
	iter = func(head *TreeNode) {
		if head == nil {
			return
		}
		iter(head.Right)
		sum += head.Val
		head.Val = sum
		iter(head.Left)
	}
	iter(root)
	return root
}

func main() {
	return
}
