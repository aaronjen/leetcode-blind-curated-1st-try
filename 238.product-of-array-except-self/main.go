package main

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	prev := nums[n-1]
	for i := 1; i < n; i++ {
		res[n-i-1] *= prev
		prev *= nums[n-i-1]
	}

	return res
}

// @lc code=end
