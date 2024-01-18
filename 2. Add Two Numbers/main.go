package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type tee func(int)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	x := 0
	var new *ListNode = &ListNode{Val: -1}
	newCurr := new
	for l1 != nil || l2 != nil {
		curr := 0
		if l1 != nil && l2 != nil {
			curr = l1.Val + l2.Val + x
		} else if l1 != nil {
			curr = l1.Val + x
		} else {
			curr = l2.Val + x
		}
		newCurr.Next = &ListNode{Val: curr % 10}
		newCurr = newCurr.Next
		x = curr / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if x != 0 {
		newCurr.Next = &ListNode{Val: x}
	}
	return new.Next
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
	list1 := arrayToList([]int{1, 2})
	list2 := arrayToList([]int{8, 8})
	res := addTwoNumbers(list1, list2)
	iterList(res, func(i int) { fmt.Printf("%d ", i) })
}
