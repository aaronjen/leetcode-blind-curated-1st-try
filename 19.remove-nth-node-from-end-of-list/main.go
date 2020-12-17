package main

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{}
	res.Next = head

	fast := res
	slow := res

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return res.Next
}

// @lc code=end
