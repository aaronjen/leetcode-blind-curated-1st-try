package main

import "strings"

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isChar(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1

	for left < right {
		if !isChar(s[left]) {
			left++
			continue
		}

		if !isChar(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true

}

// @lc code=end
