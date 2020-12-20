package main

import "container/heap"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
type element struct {
	key int
	val int
}

type maxHeap []element

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[j].val < h[i].val }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(element))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return x
}

func topKFrequent(nums []int, k int) []int {
	table := map[int]int{}

	for i := 0; i < len(nums); i++ {
		table[nums[i]]++
	}

	h := maxHeap{}
	for key, val := range table {
		heap.Push(&h, element{
			key: key,
			val: val,
		})
	}

	res := []int{}
	for i := 0; i < k; i++ {
		ele := heap.Pop(&h).(element)
		res = append(res, ele.key)
	}

	return res
}

// @lc code=end
