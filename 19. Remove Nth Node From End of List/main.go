package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEndRec(head *ListNode, n int) int {
	if head == nil {
		return 0
	}

	next := removeNthFromEndRec(head.Next, n)

	if next == n {
		head.Next = head.Next.Next
	}

	return next + 1
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	next := removeNthFromEndRec(head, n)
	if next == n {
		head = head.Next
	}
	return head
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
	list := arrayToList([]int{1})
	result := removeNthFromEnd(list, 1)

	iterList(result, func(i int) { fmt.Printf("%d ", i) })
}
