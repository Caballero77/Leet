package main

import (
	"fmt"
	"strconv"
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

func (this *Stack) Push(val int) {
	if this.Node == nil {
		this.Node = &StackNode{Val: val, Prev: nil}
		return
	}

	newNode := &StackNode{Val: val, Prev: this.Node}
	this.Node = newNode
}

func (this *Stack) Pop() int {
	val := this.Node.Val
	this.Node = this.Node.Prev
	return val
}

func isOperator(a string) bool {
	return a == "+" || a == "-" || a == "*" || a == "/"
}

func eval(operator string, a int, b int) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	}
	return a / b // "/"
}

func evalRPN(tokens []string) int {
	stack := Constructor()
	for _, val := range tokens {
		if !isOperator(val) {
			number, _ := strconv.Atoi(val)
			stack.Push(number)
			continue
		}
		b := stack.Pop()
		a := stack.Pop()
		stack.Push(eval(val, a, b))
	}
	return stack.Pop()
}

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
