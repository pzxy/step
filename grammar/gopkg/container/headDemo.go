package main

import (
	"container/heap"
	"fmt"
)

// 本示例演示了使用堆接口构建的整数堆。

// IntHeap是一个整型的整数。
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push和Pop使用指针接收器，因为它们修改切片的长度，
	// 不仅仅是其内容。
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 本示例将几个int插入到IntHeap中，检查最小值，
// 并按优先顺序删除它们。

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	fmt.Printf("Init: %d\n", (*h)[0])
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
