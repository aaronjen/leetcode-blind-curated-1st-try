package main

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	res := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev < 0 {
			prev = nums[i]
		} else {
			prev += nums[i]
		}
		if prev > res {
			res = prev
		}
	}

	return res
}

// @lc code=end
