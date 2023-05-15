package main

import (
	"fmt"
)

type RectangleItem struct {
	Height int
	Index  int
}

type StackNode struct {
	Val  RectangleItem
	Prev *StackNode
}

type Stack struct {
	Node *StackNode
}

func Constructor() Stack {
	return Stack{Node: nil}
}

func (this *Stack) Push(val RectangleItem) {
	if this.Node == nil {
		this.Node = &StackNode{Val: val, Prev: nil}
		return
	}

	newNode := &StackNode{Val: val, Prev: this.Node}
	this.Node = newNode
}

func (this *Stack) Top() RectangleItem {
	return this.Node.Val
}

func (this *Stack) Pop() RectangleItem {
	result := this.Node.Val
	this.Node = this.Node.Prev
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxStack(stack Stack, heightsLength int) int {
	result := 0
	for stack.Node != nil {
		val := stack.Pop()
		result = max(result, val.Height*(heightsLength-val.Index))
	}
	return result
}

func largestRectangleArea(heights []int) int {
	stack := Constructor()
	result := 0
	stack.Push(RectangleItem{Height: heights[0], Index: 0})
	for i := 1; i < len(heights); i++ {
		curr := i
		for stack.Node != nil && stack.Top().Height > heights[i] {
			val := stack.Pop()
			result = max(result, val.Height*(i-val.Index))
			curr = val.Index
		}
		stack.Push(RectangleItem{Height: heights[i], Index: curr})
	}
	return max(result, maxStack(stack, len(heights)))
}

func main() {
	fmt.Println(largestRectangleArea([]int{1, 2, 3, 4, 5, 6}))
}
