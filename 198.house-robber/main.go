package main

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	prev1 := 0
	prev2 := nums[0]
	for i := 1; i < len(nums); i++ {
		prev1, prev2 = prev2, findmax(prev1+nums[i], prev2)
	}
	return prev2
}

// @lc code=end
