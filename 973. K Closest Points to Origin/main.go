package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Point struct {
	X         int
	Y         int
	Disctance float64
}

type MinHeap []Point

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Disctance > h[j].Disctance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	minHeap := &MinHeap{}
	for _, point := range points {
		x, y := point[0], point[1]

		heap.Push(minHeap, Point{X: x, Y: y, Disctance: math.Sqrt(float64(x*x + y*y))})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	result := make([][]int, minHeap.Len())
	for i := 0; minHeap.Len() > 0; i++ {
		point := heap.Pop(minHeap).(Point)
		result[i] = []int{point.X, point.Y}
	}
	return result
}

func main() {
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
