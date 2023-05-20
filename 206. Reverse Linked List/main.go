package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var new *ListNode = nil
	for curr != nil {
		new = &ListNode{Val: curr.Val, Next: new}
		curr = curr.Next
	}
	return new
}

type tee func(int)

func iterList(head *ListNode, fn tee) {
	if head == nil {
		return
	}
	fn(head.Val)
	if head.Next != nil {
		iterList(head.Next, fn)
	}
}

func arrayToList(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}

	head := ListNode{Val: array[0], Next: nil}
	curr := &head
	for i := 1; i < len(array); i++ {
		new := ListNode{Val: array[i], Next: nil}
		curr.Next = &new
		curr = &new
	}
	return &head

}

func main() {
	list := arrayToList([]int{})
	res := reverseList(list)

	iterList(res, func(i int) { fmt.Printf("%d ", i) })
}
