package main

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		prev := 1
		for j := 1; j < n; j++ {
			dp[j] += prev
			prev = dp[j]
		}
	}

	return dp[n-1]
}

// @lc code=end
