package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	new := &ListNode{Val: 0}
	curr := new
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				curr.Next = &ListNode{Val: list1.Val}
				list1 = list1.Next
			} else {
				curr.Next = &ListNode{Val: list2.Val}
				list2 = list2.Next
			}
		} else if list2 == nil {
			curr.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			curr.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		curr = curr.Next
	}

	return new.Next
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
	list1 := arrayToList([]int{1, 3, 5, 7})
	list2 := arrayToList([]int{8, 9, 10, 11})
	res := mergeTwoLists(list1, list2)

	iterList(res, func(i int) { fmt.Printf("%d ", i) })
}
