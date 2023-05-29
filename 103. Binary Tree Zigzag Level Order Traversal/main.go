package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	var inner func(root *TreeNode, depth int)
	inner = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(result) == depth {
			result = append(result, []int{root.Val})
		} else {
			if depth%2 == 0 {
				result[depth] = append(result[depth], root.Val)
			} else {
				result[depth] = append([]int{root.Val}, result[depth]...)
			}
		}

		inner(root.Left, depth+1)
		inner(root.Right, depth+1)
	}

	inner(root, 0)
	return result
}

func main() {

}
