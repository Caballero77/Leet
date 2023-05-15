package main

import (
	"fmt"
	"sort"
)

type StackNode struct {
	Val  float32
	Prev *StackNode
	Size int
}

type Stack struct {
	Node *StackNode
}

func Constructor() Stack {
	return Stack{Node: nil}
}

func (this *Stack) Push(val float32) {
	if this.Node == nil {
		this.Node = &StackNode{Val: val, Prev: nil, Size: 1}
		return
	}

	newNode := &StackNode{Val: val, Prev: this.Node, Size: this.Node.Size + 1}
	this.Node = newNode
}

func (this *Stack) Top() float32 {
	return this.Node.Val
}

func (this *Stack) Size() int {
	return this.Node.Size
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]struct {
		Position int
		Speed    int
	}, len(position))
	for i := 0; i < len(position); i++ {
		cars[i] = struct {
			Position int
			Speed    int
		}{Position: position[i], Speed: speed[i]}
	}
	sort.Slice(cars, func(i int, j int) bool {
		return cars[i].Position > cars[j].Position
	})

	stack := Constructor()
	stack.Push(float32(target-cars[0].Position) / float32(cars[0].Speed))
	for i := 1; i < len(cars); i++ {
		curr := float32(target-cars[i].Position) / float32(cars[i].Speed)
		if stack.Top() < curr {
			stack.Push(curr)
		}
	}

	return stack.Size()
}

func main() {
	fmt.Println(carFleet(10, []int{9, 8, 7, 6, 5, 4}, []int{1, 2, 3, 4, 5, 6}))
}
