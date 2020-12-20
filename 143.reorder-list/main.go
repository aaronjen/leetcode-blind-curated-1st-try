package main

/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse
	var prev *ListNode
	for slow.Next != nil {
		tmp := slow.Next
		slow.Next = slow.Next.Next
		tmp.Next = prev
		prev = tmp
	}
	slow.Next = prev

	first := head
	second := slow.Next
	slow.Next = nil
	for first != nil && second != nil {
		tmp1 := first.Next
		tmp2 := second.Next
		second.Next = first.Next
		first.Next = second
		first = tmp1
		second = tmp2
	}
}

// @lc code=end
