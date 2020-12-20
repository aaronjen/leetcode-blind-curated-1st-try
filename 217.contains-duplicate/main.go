package main

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	table := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if table[nums[i]] {
			return true
		}
		table[nums[i]] = true
	}

	return false
}

// @lc code=end
