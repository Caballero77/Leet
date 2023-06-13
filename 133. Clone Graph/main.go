package main

import (
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	var visited = make(map[*Node]*Node)
	var inner func(node *Node) *Node

	inner = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if visitedValue, ok := visited[node]; ok {
			return visitedValue
		}
		newNode := &Node{Val: node.Val, Neighbors: []*Node{}}

		visited[node] = newNode

		for _, neighbor := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, inner(neighbor))
		}

		return newNode
	}

	return inner(node)
}

func main() {
	fmt.Println(1)
}
