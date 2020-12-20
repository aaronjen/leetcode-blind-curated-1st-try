package main

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) < 2 {
		return nums[0]
	}

	return findmax(robLine(nums[:len(nums)-1]), robLine(nums[1:]))
}

func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robLine(nums []int) int {
	prev1 := 0
	prev2 := nums[0]
	for i := 1; i < len(nums); i++ {
		prev1, prev2 = prev2, findmax(prev1+nums[i], prev2)
	}
	return prev2
}

// @lc code=end
