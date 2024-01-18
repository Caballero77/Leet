package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	buf := head
	len := 1
	for buf.Next != nil {
		len++
		buf = buf.Next
	}
	n := len - k%len

	if n == 0 {
		return head
	}

	buf.Next = head

	buf = head

	for i := 0; i < n-1; i++ {
		buf = buf.Next
	}

	newStart := buf.Next

	buf.Next = nil

	return newStart
}

func print(root *ListNode) {
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
}

func main() {
	list := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}
	print(list)
	list = rotateRight(list, 10000001)
	print(list)
}
