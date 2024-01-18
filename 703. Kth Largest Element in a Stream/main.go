package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	heap *IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	kthLargest := KthLargest{heap: h, k: k}
	for _, val := range nums {
		kthLargest.Add(val)
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}
	return (*this.heap)[0]
}

func main() {
	test := Constructor(5, []int{0, -1, -2, -3, -4, -5, -6, -7})
	fmt.Println(test.Add(-7))
	fmt.Println(test.Add(-8))
	fmt.Println(test.Add(-9))
	fmt.Println(test.Add(-11))
	fmt.Println(test.Add(-12))
}
