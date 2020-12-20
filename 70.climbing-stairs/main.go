package main

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	prev1 := 1
	prev2 := 1
	res := prev1 + prev2
	for i := 2; i < n; i++ {
		prev1 = prev2
		prev2 = res
		res = prev1 + prev2
	}
	return res
}

// @lc code=end
