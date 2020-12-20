package main

/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	dfs(root, &k, &res)
	return res
}

func dfs(node *TreeNode, k *int, res *int) {
	if node == nil {
		return
	}
	dfs(node.Left, k, res)
	if *k == 0 {
		return
	}
	*k--
	if *k == 0 {
		*res = node.Val
		return
	}
	dfs(node.Right, k, res)
}

// @lc code=end
