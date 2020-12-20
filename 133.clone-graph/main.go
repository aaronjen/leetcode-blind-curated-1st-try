package main

/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// Node Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// @lc code=start
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	table := map[int]*Node{}

	resVal := node.Val
	res := &Node{
		Val: resVal,
	}
	table[resVal] = res

	queue := []*Node{node}
	for len(queue) != 0 {
		org := queue[0]
		nbs := []*Node{}
		cur := table[org.Val]
		for i := 0; i < len(org.Neighbors); i++ {
			orgNb := org.Neighbors[i]
			if n, ok := table[orgNb.Val]; ok {
				nbs = append(nbs, n)
			} else {
				queue = append(queue, orgNb)
				newNode := &Node{
					Val: orgNb.Val,
				}
				table[orgNb.Val] = newNode
				nbs = append(nbs, newNode)
			}
		}
		cur.Neighbors = nbs
		queue = queue[1:]
	}

	return res
}

// @lc code=end
