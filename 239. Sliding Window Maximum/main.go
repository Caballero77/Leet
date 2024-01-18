package main

import "fmt"

type Node struct {
	Val   int
	Index int
	Next  *Node
	Prev  *Node
}

type Deque struct {
	Start *Node
	End   *Node
}

func New() *Deque {
	return &Deque{}
}

func (deque *Deque) Push(val int, i int) {
	node := &Node{Val: val, Index: i}
	for deque.End != nil && deque.End.Val < val {
		deque.End = deque.End.Next
	}
	if deque.End == nil {
		deque.Start = node
		deque.End = node
	} else {
		node.Next = deque.End
		deque.End.Prev = node
		deque.End = node
	}
}

func (deque *Deque) Peek() (bool, int) {
	if deque.Start != nil {
		return true, deque.Start.Val
	}
	return false, 0
}

func (deque *Deque) Pop(i int) {
	if deque.Start != nil && deque.Start.Index == i {
		if deque.Start.Prev != nil {
			deque.Start.Prev.Next = nil
		}
		deque.Start = deque.Start.Prev
		if deque.Start == nil {
			deque.End = nil
		}
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)
	deque := New()

	for i := 0; i < len(nums); i++ {
		deque.Push(nums[i], i)

		if i+1 >= k {
			ok, v := deque.Peek()
			if ok {
				result[i+1-k] = v
			}

			deque.Pop(i + 1 - k)
		}
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}
