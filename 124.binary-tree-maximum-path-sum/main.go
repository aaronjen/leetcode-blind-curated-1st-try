package main

/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func maxPathSum(root *TreeNode) int {
	max := root.Val
	solve(root, &max)
	return max
}

func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}

	left := findmax(solve(node.Left, max), 0)
	right := findmax(solve(node.Right, max), 0)

	sum := node.Val + left + right
	*max = findmax(*max, sum)

	return findmax(left, right) + node.Val
}

// @lc code=end
