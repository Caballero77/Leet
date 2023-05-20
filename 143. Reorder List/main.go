package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	curr := head
	var new *ListNode = nil
	for curr != nil {
		new = &ListNode{Val: curr.Val, Next: new}
		curr = curr.Next
	}
	return new
}

func reorderList(head *ListNode) {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversed := reverse(slow.Next)
	slow.Next = nil

	curr := head

	for curr != nil && reversed != nil {
		currNext := curr.Next
		reversedNext := reversed.Next
		curr.Next = reversed
		reversed.Next = currNext
		curr = currNext
		reversed = reversedNext
	}
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
	list := arrayToList([]int{1, 3, 5, 7})
	reorderList(list)

	iterList(list, func(i int) { fmt.Printf("%d ", i) })
}
