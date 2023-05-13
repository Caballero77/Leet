package main

import (
	"fmt"
)

type StackNode struct {
	Val  int
	Prev *StackNode
}

type Stack struct {
	Node *StackNode
}

func Constructor() Stack {
	return Stack{Node: nil}
}

func (this *Stack) IsEmpty() bool {
	return this.Node == nil
}

func (this *Stack) Push(val int) {
	if this.Node == nil {
		this.Node = &StackNode{Val: val, Prev: nil}
		return
	}

	newNode := &StackNode{Val: val, Prev: this.Node}
	this.Node = newNode
}

func (this *Stack) Pop() {
	this.Node = this.Node.Prev
}

func (this *Stack) Top() int {
	return this.Node.Val
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := Constructor()
	for i := 0; i < len(temperatures); i++ {
		for !stack.IsEmpty() && temperatures[stack.Top()] < temperatures[i] {
			result[stack.Top()] = i - stack.Top()
			stack.Pop()
		}
		if i == len(temperatures)-1 {
			break
		}

		if temperatures[i] < temperatures[i+1] {
			result[i] = 1
			continue
		}

		if temperatures[i] >= temperatures[i+1] {
			stack.Push(i)
		}
	}
	return result
}

func main() {
	for _, val := range dailyTemperatures([]int{100, 99, 98, 97, 96, 95, 100}) {
		fmt.Println(val)
	}
}
