package main

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{}
	node := root
	var pre *TreeNode = nil
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		node = node.Right
	}
	return true
}

// @lc code=end
