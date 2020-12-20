package main

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(board, word, 0, i, j) {
				return true
			}
		}
	}

	return false
}

func backtrack(board [][]byte, word string, ind int, row, col int) bool {
	m := len(board)
	n := len(board[0])
	if row < 0 || row >= m || col < 0 || col >= n {
		return false
	}

	if word[ind] != board[row][col] {
		return false
	}

	if ind+1 == len(word) {
		return true
	}
	b := board[row][col]
	board[row][col] = '-'
	if backtrack(board, word, ind+1, row+1, col) ||
		backtrack(board, word, ind+1, row-1, col) ||
		backtrack(board, word, ind+1, row, col+1) ||
		backtrack(board, word, ind+1, row, col-1) {
		return true
	}
	board[row][col] = b

	return false
}

// @lc code=end
