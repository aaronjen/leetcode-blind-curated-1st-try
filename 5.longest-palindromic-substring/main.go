package main

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		extendString(s, i, i, &res)
		extendString(s, i, i+1, &res)
	}

	return res
}

func extendString(s string, left, right int, res *string) {
	tmp := ""
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			tmp = s[left : right+1]
			left--
			right++
		} else {
			break
		}
	}
	if len(tmp) > len(*res) {
		*res = tmp
	}
}

// @lc code=end
