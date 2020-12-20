package main

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	ind := 0
	return helper(preorder, inorder, &ind, 0, len(preorder))
}

// [start, end)
func helper(preorder []int, inorder []int, ind *int, start, end int) *TreeNode {
	val := preorder[*ind]
	inorederInd := start
	for i := start; i < end; i++ {
		if inorder[i] == val {
			inorederInd = i
		}
	}
	leftStart := start
	leftEnd := inorederInd
	var left *TreeNode
	if leftEnd > leftStart {
		*ind++
		left = helper(preorder, inorder, ind, leftStart, leftEnd)
	}

	rightStart := inorederInd + 1
	rightEnd := end
	var right *TreeNode
	if rightEnd > rightStart {
		*ind++
		right = helper(preorder, inorder, ind, rightStart, rightEnd)
	}

	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

// @lc code=end
