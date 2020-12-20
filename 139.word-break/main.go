package main

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i <= len(s); i++ {
		if dp[i] {
			for _, w := range wordDict {
				if i+len(w) <= len(s) && s[i:i+len(w)] == w {
					dp[i+len(w)] = true
				}
			}
		}
	}

	return dp[len(s)]
}

// @lc code=end
