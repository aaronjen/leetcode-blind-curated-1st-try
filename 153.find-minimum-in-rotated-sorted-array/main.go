package main

/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}

		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

// @lc code=end
