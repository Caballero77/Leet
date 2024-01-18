package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	root := ListNode{Val: -10001, Next: nil}
	curr := root.Next

	empty := 0
	for true {
		empty = 0
		minI, minV := 0, 10001
		for i := range lists {
			if lists[i] == nil {
				empty++
				continue
			}
			if lists[i].Val < minV {
				minI = i
				minV = lists[i].Val
			}
		}

		if empty == len(lists) {
			break
		}

		if curr == nil {
			root.Next = &ListNode{Val: minV, Next: nil}
			curr = root.Next
		} else {
			curr.Next = &ListNode{Val: minV, Next: nil}
			curr = curr.Next
		}

		lists[minI] = lists[minI].Next
	}

	return root.Next
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
	list3 := arrayToList([]int{-1})
	res := mergeKLists([]*ListNode{list1, list2, list3})

	iterList(res, func(i int) { fmt.Printf("%d ", i) })
	fmt.Println()
}
