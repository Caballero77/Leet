package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type tee func(int)

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	next := head.Next

	curr.Next = curr.Next.Next
	head = next
	head.Next = curr

	if curr == nil || curr.Next == nil {
		return head
	}

	curr = curr.Next
	next = curr.Next
	prev := head.Next

	for curr != nil && next != nil {
		curr.Next = curr.Next.Next
		next.Next = curr
		prev.Next = next

		if curr == nil || curr.Next == nil {
			break
		}
		prev = curr
		curr = curr.Next
		next = curr.Next
	}
	return head
}

func iterList(head *ListNode, fn tee) {
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
	list := arrayToList([]int{1, 2})
	res := swapPairs(list)

	iterList(res, func(i int) { fmt.Printf("%d ", i) })
}
