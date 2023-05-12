package main

import "fmt"

type MinStackNode struct {
	Val  int
	Prev *MinStackNode
	Min  int
}

type MinStack struct {
	Node *MinStackNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Constructor() MinStack {
	return MinStack{Node: nil}
}

func (this *MinStack) Push(val int) {
	if this.Node == nil {
		this.Node = &MinStackNode{Val: val, Min: val, Prev: nil}
		return
	}

	newNode := &MinStackNode{Val: val, Prev: this.Node, Min: min(this.Node.Min, val)}
	this.Node = newNode
}

func (this *MinStack) Pop() {
	this.Node = this.Node.Prev
}

func (this *MinStack) Top() int {
	return this.Node.Val
}

func (this *MinStack) GetMin() int {
	return this.Node.Min
}

func main() {
	minStack := Constructor()
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	minStack.Pop()
	minStack.Pop()
	minStack.Pop()
	minStack.Push(2147483647)
	minStack.Push(-2147483648)
	minStack.Pop()
	fmt.Println(minStack.GetMin()) // return -2
}
