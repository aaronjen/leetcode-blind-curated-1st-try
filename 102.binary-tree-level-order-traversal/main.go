package main

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	res := [][]int{}
	for len(q) != 0 {
		tmpArr := []int{}
		tmp := []*TreeNode{}
		for i := 0; i < len(q); i++ {
			node := q[i]
			tmpArr = append(tmpArr, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		res = append(res, tmpArr)
		q = tmp
	}

	return res
}

// @lc code=end
