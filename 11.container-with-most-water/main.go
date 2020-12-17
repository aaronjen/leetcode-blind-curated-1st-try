package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func findmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0

	for left < right {
		lh := height[left]
		rh := height[right]
		h := findmin(lh, rh)
		volume := h * (right - left)
		if res < volume {
			res = volume
		}

		if lh > rh {
			right--
		} else {
			left++
		}

	}

	return res
}

// @lc code=end
