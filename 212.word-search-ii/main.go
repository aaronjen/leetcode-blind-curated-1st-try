package main

/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */

// @lc code=start

type node struct {
	next []*node
	word string
}

func findWords(board [][]byte, words []string) []string {
	root := buildTrie(words)
	res := []string{}
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, root, i, j, &res)
		}
	}
	return res
}

func dfs(board [][]byte, root *node, row, col int, res *[]string) {
	m := len(board)
	n := len(board[0])
	if row < 0 || row >= m || col < 0 || col >= n {
		return
	}

	if board[row][col] == '#' {
		return
	}
	nd := root.next[int(board[row][col]-'a')]
	if nd == nil {
		return
	}

	if nd.word != "" {
		*res = append(*res, nd.word)
		nd.word = ""
	}

	tmp := board[row][col]
	board[row][col] = '#'
	dfs(board, nd, row-1, col, res)
	dfs(board, nd, row+1, col, res)
	dfs(board, nd, row, col-1, res)
	dfs(board, nd, row, col+1, res)
	board[row][col] = tmp
}

func buildTrie(words []string) *node {
	root := &node{
		next: make([]*node, 26),
	}
	for _, w := range words {
		nd := root
		for i := 0; i < len(w); i++ {
			ind := int(w[i] - 'a')
			if tmp := nd.next[ind]; tmp != nil {
				nd = tmp
			} else {
				nd.next[ind] = &node{
					next: make([]*node, 26),
				}
				nd = nd.next[ind]
			}
		}
		nd.word = w
	}

	return root
}

// @lc code=end
