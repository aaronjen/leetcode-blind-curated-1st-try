package main

import "math"

/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start
func pacificAtlantic(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return [][]int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return [][]int{}
	}

	pa := make([][]bool, m)
	at := make([][]bool, m)
	for i := 0; i < m; i++ {
		pa[i] = make([]bool, n)
		at[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		flow(matrix, pa, i, 0, math.MinInt32)
		flow(matrix, at, i, n-1, math.MinInt32)
	}
	for i := 0; i < n; i++ {
		flow(matrix, pa, 0, i, math.MinInt32)
		flow(matrix, at, m-1, i, math.MinInt32)
	}

	res := [][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pa[i][j] && at[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func flow(matrix [][]int, grid [][]bool, row, col, height int) {
	m := len(matrix)
	n := len(matrix[0])
	if row < 0 || row >= m || col < 0 || col >= n {
		return
	}

	if matrix[row][col] < height {
		return
	}

	if grid[row][col] {
		return
	}

	grid[row][col] = true
	flow(matrix, grid, row+1, col, matrix[row][col])
	flow(matrix, grid, row-1, col, matrix[row][col])
	flow(matrix, grid, row, col+1, matrix[row][col])
	flow(matrix, grid, row, col-1, matrix[row][col])
}

// @lc code=end
