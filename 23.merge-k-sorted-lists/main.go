package main

import "container/heap"

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	filteredlists := []*ListNode{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			filteredlists = append(filteredlists, lists[i])
		}
	}

	if len(filteredlists) == 0 {
		return nil
	}

	var mh minHeap = filteredlists
	heap.Init(&mh)
	res := heap.Pop(&mh).(*ListNode)
	if res.Next != nil {
		heap.Push(&mh, res.Next)
	}
	cur := res
	for len(mh) != 0 {
		tmp := heap.Pop(&mh).(*ListNode)
		cur.Next = tmp
		cur = cur.Next
		if tmp.Next != nil {
			heap.Push(&mh, tmp.Next)
		}
	}

	return res
}

// @lc code=end
