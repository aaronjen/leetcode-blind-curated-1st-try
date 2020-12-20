package main

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func reverseList(head *ListNode) *ListNode {
	res := &ListNode{}
	for head != nil {
		tmp := head.Next
		head.Next = res.Next
		res.Next = head
		head = tmp
	}

	return res.Next
}

// @lc code=end
