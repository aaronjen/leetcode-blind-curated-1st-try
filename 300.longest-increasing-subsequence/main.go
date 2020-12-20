package main

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tail := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		ind := 0
		for ind < len(tail) && n > tail[ind] {
			ind++
		}
		if ind == len(tail) {
			tail = append(tail, n)
		} else {
			tail[ind] = n
		}
	}

	return len(tail)
}

// @lc code=end
