package main

import "container/heap"

/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 */

// @lc code=start

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return x
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return x
}

// MedianFinder struct
type MedianFinder struct {
	left   maxHeap
	right  minHeap
	isEven bool
}

// Constructor func
/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		left:   maxHeap{},
		right:  minHeap{},
		isEven: true,
	}
}

// AddNum func
func (t *MedianFinder) AddNum(num int) {
	if t.isEven {
		heap.Push(&t.right, num)
		heap.Push(&t.left, heap.Pop(&t.right))
	} else {
		heap.Push(&t.left, num)
		heap.Push(&t.right, heap.Pop(&t.left))
	}
	t.isEven = !t.isEven
}

// FindMedian func
func (t *MedianFinder) FindMedian() float64 {
	if t.isEven {
		return (float64(t.left[0]) + float64(t.right[0])) / 2
	}
	return float64(t.left[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
