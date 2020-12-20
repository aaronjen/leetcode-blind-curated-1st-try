package main

/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if sameTree(s, t) {
		return true
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func sameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return sameTree(s.Left, t.Left) && sameTree(s.Right, t.Right)
}

// @lc code=end
