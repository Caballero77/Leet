package main

import "fmt"

type Node struct {
	Val  int
	Key  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	Cache    map[int]*Node
	Capacity int
	Start    *Node
	End      *Node
}

func Constructor(capacity int) LRUCache {
	result := LRUCache{
		Cache:    make(map[int]*Node),
		Capacity: capacity,
		Start:    &Node{Val: -1, Key: -1},
		End:      &Node{Val: -1, Key: -1},
	}

	result.Start.Next = result.End
	result.End.Prev = result.Start

	return result
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.Cache[key]; ok {
		value.Prev.Next = value.Next
		value.Next.Prev = value.Prev

		value.Next = this.Start.Next
		value.Prev = this.Start
		this.Start.Next.Prev = value
		this.Start.Next = value
		return value.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	new := &Node{
		Val:  value,
		Key:  key,
		Next: this.Start.Next,
		Prev: this.Start,
	}
	this.Cache[key] = new
	this.Start.Next.Prev = new
	this.Start.Next = new

	if len(this.Cache) > this.Capacity {
		delete(this.Cache, this.End.Prev.Key)
		this.End.Prev.Prev.Next = this.End
		this.End.Prev = this.End.Prev.Prev
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))

	cache.Put(1, 1)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
}
