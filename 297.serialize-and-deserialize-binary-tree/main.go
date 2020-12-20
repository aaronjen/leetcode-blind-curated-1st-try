package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

// Codec struct
type Codec struct {
}

// Constructor func
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (t *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var sb strings.Builder
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := []*TreeNode{}
		for _, nd := range queue {
			if nd == nil {
				sb.WriteString("x,")
			} else {
				sb.WriteString(fmt.Sprintf("%v,", nd.Val))
				tmp = append(tmp, nd.Left, nd.Right)
			}
		}
		queue = tmp
	}

	return sb.String()
}

// Deserializes your encoded data to tree.
func (t *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	dataList := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(dataList[0])
	root := &TreeNode{
		Val: rootVal,
	}
	ind := 1
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := []*TreeNode{}
		for _, nd := range queue {
			left := valueToNode(dataList[ind])
			ind++
			if left != nil {
				tmp = append(tmp, left)
			}
			right := valueToNode(dataList[ind])
			if right != nil {
				tmp = append(tmp, right)
			}
			ind++
			nd.Left = left
			nd.Right = right
		}
		queue = tmp
	}

	return root
}

func valueToNode(s string) *TreeNode {
	if s == "x" {
		return nil
	}
	val, _ := strconv.Atoi(s)
	return &TreeNode{
		Val: val,
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
