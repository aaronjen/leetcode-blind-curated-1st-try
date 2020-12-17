package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 1
	prevStart := 0
	for i := 1; i < len(s); i++ {
		b := s[i]
		prev := s[prevStart:i]
		if ind := strings.IndexByte(prev, b); ind == -1 {
			if res < i-prevStart+1 {
				res = i - prevStart + 1
			}
		} else {
			prevStart += ind + 1
		}

	}

	return res
}

// @lc code=end
