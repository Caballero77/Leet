package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	result := make(map[*Node]*Node)

	curr := head
	for curr != nil {
		clone := &Node{Val: curr.Val}
		result[curr] = clone
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		clone := result[curr]
		clone.Next = result[curr.Next]
		clone.Random = result[curr.Random]
		curr = curr.Next
	}

	return result[head]
}

func main() {

}
