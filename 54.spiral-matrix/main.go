package main

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	res := []int{}

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	row := 0
	col := 0
	dir := 0
	for len(res) != m*n {
		res = append(res, matrix[row][col])
		matrix[row][col] = '-'

		nextRow := row + directions[dir][0]
		nextCol := col + directions[dir][1]
		if nextRow < 0 ||
			nextRow >= m ||
			nextCol < 0 ||
			nextCol >= n ||
			matrix[nextRow][nextCol] == '-' {
			dir = (dir + 1) % 4
			nextRow = row + directions[dir][0]
			nextCol = col + directions[dir][1]
		}
		row = nextRow
		col = nextCol
	}

	return res
}

// @lc code=end
