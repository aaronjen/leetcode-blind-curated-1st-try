package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		prevLeft := math.MinInt16
		prevRight := math.MinInt16
		for left < right {
			ln := nums[left]
			rn := nums[right]
			if ln == prevLeft {
				left++
				continue
			}
			if rn == prevRight {
				right--
				continue
			}
			sum := nums[i] + ln + rn
			if sum == 0 {
				res = append(res, []int{nums[i], ln, rn})
				prevLeft = ln
				prevRight = rn
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}

		}
		for i+1 < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}

// @lc code=end
