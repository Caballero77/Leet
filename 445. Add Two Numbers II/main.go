package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type tee func(int)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	stack2 := make([]int, 0)
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	var result *ListNode = nil
	next := 0
	for len(stack1) != 0 || len(stack2) != 0 {
		x := next
		if len(stack1) != 0 {
			x += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			x += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		result = &ListNode{Val: x % 10, Next: result}
		next = x / 10
	}
	if next != 0 {
		result = &ListNode{Val: next, Next: result}
	}
	return result
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
	list1 := arrayToList([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	list2 := arrayToList([]int{1})
	res := addTwoNumbers(list1, list2)

	iterList(res, func(i int) { fmt.Printf("%d ", i) })
}
